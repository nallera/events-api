package events

import (
	"time"
)

type Repository interface {
	GetCurrentEvents() ([]*Event, error)
}

type Database interface {
	Reader
	Writer
}

type Reader interface {
	GetByID(id uint64) (*Event, error)
	GetByDateRange(dateStart, dateEnd time.Time) ([]*Event, error)
}

type Writer interface {
	Insert(event *Event) error
	MultiInsert(events []*Event) error
}
