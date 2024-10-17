package repository

import (
	"database/sql"
	"fmt"
	"time"

	"events-api/internal/events"
	eventserrors "events-api/pkg/errors"
)

const QueryCreate string = `CREATE TABLE IF NOT EXISTS events (
    event_id INTEGER NOT NULL,  
    base_event_id INTEGER NOT NULL,  
    title TEXT NOT NULL,
    event_start_date DATETIME NOT NULL,
    event_end_date DATETIME NOT NULL,
    sell_from DATETIME NOT NULL,
    sell_to DATETIME NOT NULL,
    sold_out BOOL NOT NULL,
    min_price FLOAT NOT NULL,  
    max_price FLOAT NOT NULL,
	PRIMARY KEY (event_id, base_event_id));
	CREATE INDEX idx_event_start_date ON events(event_start_date);
	CREATE INDEX idx_event_end_date ON events(event_end_date);`
const QueryGetByID = `SELECT * FROM events WHERE event_id = ?`
const QueryGetByTimeRange = `SELECT * FROM events WHERE event_start_date >= ? AND event_end_date <= ?`
const QueryInsertOrUpdate = `INSERT INTO events VALUES(?,?,?,?,?,?,?,?,?,?)  
	ON CONFLICT (event_id, base_event_id) 
	DO UPDATE SET title = ?, event_start_date = ?, event_end_date = ?, sell_from = ?, sell_to = ?, sold_out = ?, min_price = ?, max_price = ? 
   WHERE event_id = ? AND base_event_id = ?;`

type sqlite struct {
	client *sql.DB
}

func NewSQLiteRepository(client *sql.DB) events.Database {
	return &sqlite{
		client: client,
	}
}

func (s *sqlite) GetByID(id uint64) (*events.Event, error) {
	row := s.client.QueryRow(QueryGetByID, id)

	sqliteEventModel := new(SQLiteEventModel)
	var err error
	if err = row.Scan(&sqliteEventModel.EventID, &sqliteEventModel.BaseEventID, &sqliteEventModel.Title,
		&sqliteEventModel.EventStartDate, &sqliteEventModel.EventEndDate, &sqliteEventModel.SellFrom,
		&sqliteEventModel.SellTo, &sqliteEventModel.SoldOut, &sqliteEventModel.MinPrice, &sqliteEventModel.MaxPrice); err == sql.ErrNoRows {
		println(fmt.Sprintf("error getting event %d from sqlite db: id not found", id))
		return &events.Event{}, eventserrors.NewNotFoundError(fmt.Sprintf("error getting event %d from sqlite db: id not found", id))
	}

	println(fmt.Sprintf("got event %d from sqlite db: %+v", id, sqliteEventModel))
	appEventModel := SQLiteEventModelToApp(sqliteEventModel)

	return appEventModel, err
}

func (s *sqlite) GetByDateRange(dateStart, dateEnd time.Time) ([]*events.Event, error) {
	rows, err := s.client.Query(QueryGetByTimeRange, dateStart, dateEnd)
	if err != nil {
		return nil, fmt.Errorf("error getting events by date range in sqlite db: %w", err)
	}
	defer rows.Close()

	var sqliteEventsModel []*SQLiteEventModel

	for rows.Next() {
		sqliteEventModel := new(SQLiteEventModel)
		err = rows.Scan(&sqliteEventModel.EventID, &sqliteEventModel.BaseEventID, &sqliteEventModel.Title,
			&sqliteEventModel.EventStartDate, &sqliteEventModel.EventEndDate, &sqliteEventModel.SellFrom,
			&sqliteEventModel.SellTo, &sqliteEventModel.SoldOut, &sqliteEventModel.MinPrice, &sqliteEventModel.MaxPrice)
		if err != nil {
			return nil, fmt.Errorf("error getting events by date range in sqlite db: %w", err)
		}
		sqliteEventsModel = append(sqliteEventsModel, sqliteEventModel)
	}

	appEventsModel := SQLiteEventsModelToApp(sqliteEventsModel)

	return appEventsModel, err
}

func (s *sqlite) Insert(event *events.Event) error {
	_, err := s.client.Exec(QueryInsertOrUpdate, event.EventID, event.BaseEventID, event.Title, event.EventStartDate,
		event.EventEndDate, event.SellFrom, event.SellTo, event.SoldOut, event.MinPrice, event.MaxPrice,
		event.Title, event.EventStartDate, event.EventEndDate, event.SellFrom, event.SellTo, event.SoldOut, event.MinPrice, event.MaxPrice,
		event.EventID, event.BaseEventID)
	if err != nil {
		return fmt.Errorf("error inserting event with id %d to sqlite db: %w", event.EventID, err)
	}

	return nil
}

func (s *sqlite) MultiInsert(events []*events.Event) error {
	for _, event := range events {
		if err := s.Insert(event); err != nil {
			println("error while multi-inserting events to sqlite db, event with id %d: %w", event.EventID, err)
		}
	}

	return nil
}
