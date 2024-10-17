package events_test

import (
	"errors"
	"events-api/internal/events"
	"events-api/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEventService_GetEventsInTimeRange(t *testing.T) {
	type depFields struct {
		providerXRepositoryMock *test.ProviderXRepositoryMock
		providerXDatabaseMock   *test.ProviderXDatabaseMock
	}
	type input struct {
		startsAt time.Time
		endsAt   time.Time
	}
	type output struct {
		response []*events.Event
		err      error
	}

	tests := []struct {
		name   string
		input  input
		on     func(*depFields)
		assert func(*testing.T, *output)
	}{
		{
			name: "get events in time range successfully",
			input: input{
				startsAt: test.CreateTime(0, 0),
				endsAt:   test.CreateTime(1, 0),
			},
			on: func(df *depFields) {
				df.providerXRepositoryMock.On("GetCurrentEvents").Return(test.MakeAppEvents(false), nil).Once()
				df.providerXDatabaseMock.On("MultiInsert").Return(nil).Once()
				df.providerXDatabaseMock.On("GetByDateRange", test.CreateTime(0, 0),
					test.CreateTime(1, 0)).Return(test.MakeAppEvents(false), nil).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.Nil(t, out.err)
				assert.Equal(t, test.MakeAppEvents(false), out.response)
			},
		},
		{
			name: "continue if error in get current events but no multiinsert is done",
			input: input{
				startsAt: test.CreateTime(0, 0),
				endsAt:   test.CreateTime(1, 0),
			},
			on: func(df *depFields) {
				df.providerXRepositoryMock.On("GetCurrentEvents").Return(nil, errors.New("test-error")).Once()
				df.providerXDatabaseMock.On("GetByDateRange", test.CreateTime(0, 0),
					test.CreateTime(1, 0)).Return(test.MakeAppEvents(false), nil).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.Nil(t, out.err)
				assert.Equal(t, test.MakeAppEvents(false), out.response)
			},
		},
		{
			name: "error if error in saving current events",
			input: input{
				startsAt: test.CreateTime(0, 0),
				endsAt:   test.CreateTime(1, 0),
			},
			on: func(df *depFields) {
				df.providerXRepositoryMock.On("GetCurrentEvents").Return(test.MakeAppEvents(false), nil).Once()
				df.providerXDatabaseMock.On("MultiInsert").Return(errors.New("test-error")).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.ErrorContains(t, out.err, "failed to store current events from provider x into database: test-error")
			},
		},
		{
			name: "error if error in getbydaterange",
			input: input{
				startsAt: test.CreateTime(0, 0),
				endsAt:   test.CreateTime(1, 0),
			},
			on: func(df *depFields) {
				df.providerXRepositoryMock.On("GetCurrentEvents").Return(test.MakeAppEvents(false), nil).Once()
				df.providerXDatabaseMock.On("MultiInsert").Return(nil).Once()
				df.providerXDatabaseMock.On("GetByDateRange", test.CreateTime(0, 0),
					test.CreateTime(1, 0)).Return(nil, errors.New("test-error")).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.ErrorContains(t, out.err, "failed to get events from the database: test-error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Having
			providerXRepositoryMock := new(test.ProviderXRepositoryMock)
			providerXDatabaseMock := new(test.ProviderXDatabaseMock)
			eventService := events.NewEventService(providerXRepositoryMock, providerXDatabaseMock)

			f := &depFields{providerXRepositoryMock: providerXRepositoryMock, providerXDatabaseMock: providerXDatabaseMock}

			tt.on(f)

			//When
			result, err := eventService.GetEventsInTimeRange(tt.input.startsAt, tt.input.endsAt)

			//Then
			tt.assert(t, &output{result, err})
			providerXRepositoryMock.AssertExpectations(t)
			providerXDatabaseMock.AssertExpectations(t)
		})
	}
}
