package events

import (
	"context"
	"time"
)

type Repository interface {
	GetCurrentEvents() ([]*Event, error)
}

type Reader interface {
	Get(ctx context.Context, dateStart, dateEnd time.Time) (*Event, error)
	Save(ctx context.Context, event *Event) error
}
