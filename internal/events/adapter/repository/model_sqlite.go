package repository

import (
	"events-api/internal/events"
	"time"
)

type SQLiteEventModel struct {
	EventID        uint64
	BaseEventID    uint64
	Title          string
	EventStartDate time.Time
	EventEndDate   time.Time
	SellFrom       time.Time
	SellTo         time.Time
	SoldOut        bool
	MinPrice       float64
	MaxPrice       float64
}

func SQLiteEventsModelToApp(sqliteModels []*SQLiteEventModel) []*events.Event {
	var result []*events.Event

	for _, sqliteModel := range sqliteModels {
		result = append(result, SQLiteEventModelToApp(sqliteModel))
	}

	return result
}

func SQLiteEventModelToApp(sqliteModel *SQLiteEventModel) *events.Event {
	return &events.Event{
		EventID:        sqliteModel.EventID,
		BaseEventID:    sqliteModel.BaseEventID,
		Title:          sqliteModel.Title,
		EventStartDate: sqliteModel.EventStartDate,
		EventEndDate:   sqliteModel.EventEndDate,
		SellFrom:       sqliteModel.SellFrom,
		SellTo:         sqliteModel.SellTo,
		SoldOut:        sqliteModel.SoldOut,
		Zones:          nil,
		MinPrice:       sqliteModel.MinPrice,
		MaxPrice:       sqliteModel.MaxPrice,
	}
}
