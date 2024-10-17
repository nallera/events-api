package main

import (
	"database/sql"
	"events-api/internal/events"
	"fmt"
	"log"
	"net/http"
	"os"

	"events-api/config"
	"events-api/internal/events/adapter/repository"
	"events-api/internal/http/reader"
	eventsconfig "events-api/pkg/config"
	"events-api/pkg/server"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.GetConfig()

	providerXRestClient := server.NewRestClient(cfg.ProviderXRepository)
	providerXRepository := repository.NewRepository(providerXRestClient)

	sqliteDB := initializeSQLiteDB(cfg.ProviderXDatabase)
	providerXDatabase := repository.NewSQLiteRepository(sqliteDB)

	eventService := events.NewEventService(providerXRepository, providerXDatabase)

	readerHandler := reader.NewHTTPHandler(eventService)

	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/search", readerHandler.GetEventsInTimeRange).Methods("GET")

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), muxRouter)
}

func initializeSQLiteDB(sqliteConfig eventsconfig.SQLiteConfig) *sql.DB {
	createDBFile(sqliteConfig)

	db, err := sql.Open("sqlite3", sqliteConfig.Name)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec(repository.QueryCreate); err != nil {
		log.Fatal(err)
	}
	return db
}

func createDBFile(sqliteConfig eventsconfig.SQLiteConfig) {
	dbFile := sqliteConfig.Name
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		file, err := os.Create(dbFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}
