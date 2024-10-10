package main

import (
	"events-api/config"
	"events-api/internal/events/adapter/repository"
	"events-api/internal/http/reader"
	"events-api/pkg/server"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.GetConfig()

	providerXRestClient := server.NewRestClient(cfg.ProviderXRepository)
	providerXRepository := repository.NewRepository(providerXRestClient)
	readerHandler := reader.NewHTTPHandler(providerXRepository)

	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/search", readerHandler.GetEventsInTimeRange).Methods("GET")

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", port), muxRouter)
}
