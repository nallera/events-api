package reader_test

import (
	"encoding/json"
	"errors"
	"events-api/internal/http/reader"
	"events-api/test"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_GetEventsInTimeRange(t *testing.T) {
	type depFields struct {
		eventServiceMock *test.EventServiceMock
	}
	type input struct {
		startsAt string
		endsAt   string
	}
	type output struct {
		response reader.EventsResponseModel
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
				startsAt: "starts_at=" + test.CreateTime(0, 0).Format(time.RFC3339),
				endsAt:   "ends_at=" + test.CreateTime(1, 0).Format(time.RFC3339),
			},
			on: func(df *depFields) {
				df.eventServiceMock.On("GetEventsInTimeRange").Return(test.MakeAppEvents(false), nil).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.Nil(t, out.response.Error)
				assert.Equal(t, test.MakeResponseEvents(), out.response.Data.Events)
			},
		},
		{
			name: "error if missing query param",
			input: input{
				startsAt: "",
				endsAt:   "ends_at=" + test.CreateTime(1, 0).Format(time.RFC3339),
			},
			on: func(df *depFields) {},
			assert: func(t *testing.T, out *output) {
				assert.Contains(t, out.response.Error.Message, "query parameters starts_at and ends_at are mandatory")
			},
		},
		{
			name: "error if error in query param starts_at",
			input: input{
				startsAt: "starts_at=invalid-date",
				endsAt:   "ends_at=" + test.CreateTime(1, 0).Format(time.RFC3339),
			},
			on: func(df *depFields) {},
			assert: func(t *testing.T, out *output) {
				assert.Contains(t, out.response.Error.Message, "error parsing starts_at query param:")
			},
		},
		{
			name: "error if error in query param ends_at",
			input: input{
				startsAt: "starts_at=" + test.CreateTime(0, 0).Format(time.RFC3339),
				endsAt:   "ends_at=invalid-date",
			},
			on: func(df *depFields) {},
			assert: func(t *testing.T, out *output) {
				assert.Contains(t, out.response.Error.Message, "error parsing ends_at query param:")
			},
		},
		{
			name: "error if error in GetEventsInTimeRange",
			input: input{
				startsAt: "starts_at=" + test.CreateTime(0, 0).Format(time.RFC3339),
				endsAt:   "ends_at=" + test.CreateTime(1, 0).Format(time.RFC3339),
			},
			on: func(df *depFields) {
				df.eventServiceMock.On("GetEventsInTimeRange").Return(nil, errors.New("test-error")).Once()
			},
			assert: func(t *testing.T, out *output) {
				assert.Contains(t, out.response.Error.Message, "error getting events in time range: test-error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Having
			eventServiceMock := new(test.EventServiceMock)
			readerHandler := reader.NewHTTPHandler(eventServiceMock)

			localTestURL := fmt.Sprintf("http://localhost:8080/test?%s&%s",
				tt.input.startsAt, tt.input.endsAt)
			req := httptest.NewRequest(http.MethodGet, localTestURL, nil)

			w := httptest.NewRecorder()

			f := &depFields{eventServiceMock: eventServiceMock}

			tt.on(f)

			//When
			readerHandler.GetEventsInTimeRange(w, req)

			//Then
			body, _ := io.ReadAll(w.Body)
			var result reader.EventsResponseModel
			_ = json.Unmarshal(body, &result)
			tt.assert(t, &output{result})
			eventServiceMock.AssertExpectations(t)
		})
	}
}
