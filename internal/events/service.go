package events

import (
	"fmt"
	"time"
)

type EventService interface {
	GetEventsInTimeRange(startsAt, endsAt time.Time) ([]*Event, error)
}

type eventService struct {
	providerXRepository Repository
	providerXDatabase   Database
}

func NewEventService(providerXRepository Repository, providerXDatabase Database) EventService {
	return &eventService{
		providerXRepository: providerXRepository,
		providerXDatabase:   providerXDatabase,
	}
}

func (s *eventService) GetEventsInTimeRange(startsAt, endsAt time.Time) ([]*Event, error) {
	currentEvents, err := s.providerXRepository.GetCurrentEvents()
	if err != nil {
		println(fmt.Sprintf("failed to get current events from provider x: %s", err.Error()))
	}

	if currentEvents != nil {
		err = s.providerXDatabase.MultiInsert(currentEvents)
		if err != nil {
			println(fmt.Sprintf("failed to store current events from provider x into database: %s", err.Error()))
			return nil, fmt.Errorf("failed to store current events from provider x into database: %w", err)
		}
	}

	filteredEvents, err := s.providerXDatabase.GetByDateRange(startsAt, endsAt)
	if err != nil {
		println(fmt.Sprintf("failed to get events from the database: %s", err.Error()))
		return nil, fmt.Errorf("failed to get events from the database: %w", err)
	}

	return filteredEvents, nil
}
