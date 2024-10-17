package repository_test

import (
	"errors"
	"events-api/internal/events"
	"events-api/internal/events/adapter/repository"
	"events-api/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRest_GetCurrentEvents(t *testing.T) {
	type depFields struct {
		restClient *test.RestClientMock
	}
	type output struct {
		response []*events.Event
		err      error
	}

	tests := []struct {
		name   string
		on     func(*depFields)
		assert func(*testing.T, *output)
	}{
		{
			name: "get current events successfully",
			on: func(df *depFields) {
				df.restClient.On("Get", "get_current_events").Return(nil).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.NoError(t, out.err)
			},
		},
		{
			name: "error if get addresses error",
			on: func(df *depFields) {
				df.restClient.On("Get", "get_current_events").Return(errors.New("test-error")).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.ErrorContains(t, out.err, "failed to get current events from provider x: test-error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Having
			client := new(test.RestClientMock)
			restRepository := repository.NewRepository(client)

			f := &depFields{restClient: client}

			tt.on(f)

			//When
			result, err := restRepository.GetCurrentEvents()

			//Then
			tt.assert(t, &output{response: result, err: err})
			client.AssertExpectations(t)
		})
	}
}
