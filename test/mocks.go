package test

import (
	"events-api/internal/events"
	"github.com/stretchr/testify/mock"
	"time"
)

type RestClientMock struct {
	mock.Mock
}

func (m *RestClientMock) Get(uriKey string, uriParams map[string]string, body interface{}, result interface{}, headers map[string]string) error {
	args := m.Called(uriKey)
	e, _ := args.Get(0).(error)

	return e
}

type ProviderXRepositoryMock struct {
	mock.Mock
}

func (m *ProviderXRepositoryMock) GetCurrentEvents() ([]*events.Event, error) {
	args := m.Called()
	ev, _ := args.Get(0).([]*events.Event)
	e, _ := args.Get(1).(error)

	return ev, e
}

type ProviderXDatabaseMock struct {
	mock.Mock
}

func (m *ProviderXDatabaseMock) GetByID(id uint64) (*events.Event, error) {
	args := m.Called(id)
	ev, _ := args.Get(0).(*events.Event)
	e, _ := args.Get(1).(error)

	return ev, e
}

func (m *ProviderXDatabaseMock) GetByDateRange(dateStart, dateEnd time.Time) ([]*events.Event, error) {
	args := m.Called(dateStart, dateEnd)
	ev, _ := args.Get(0).([]*events.Event)
	e, _ := args.Get(1).(error)

	return ev, e
}

func (m *ProviderXDatabaseMock) Insert(event *events.Event) error {
	args := m.Called()
	e, _ := args.Get(0).(error)

	return e
}

func (m *ProviderXDatabaseMock) MultiInsert(events []*events.Event) error {
	args := m.Called()
	e, _ := args.Get(0).(error)

	return e
}

type EventServiceMock struct {
	mock.Mock
}

func (m *EventServiceMock) GetEventsInTimeRange(startsAt, endsAt time.Time) ([]*events.Event, error) {
	args := m.Called()
	ev, _ := args.Get(0).([]*events.Event)
	e, _ := args.Get(1).(error)

	return ev, e
}
