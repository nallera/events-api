package reader

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"events-api/internal/events"
	"events-api/internal/events/adapter/repository"
)

type Handler interface {
	GetEventsInTimeRange(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	eventService events.EventService
}

func NewHTTPHandler(eventService events.EventService) Handler {
	return &handler{eventService: eventService}
}

func (h *handler) GetEventsInTimeRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	startsAt, endsAt, err := parseRequestParameters(r)
	if err != nil {
		buildResponse(w, nil, http.StatusBadRequest,
			fmt.Sprintf("error parsing request parameters: %s", err.Error()))
		return
	}

	filteredEvents, err := h.eventService.GetEventsInTimeRange(*startsAt, *endsAt)
	if err != nil {
		buildResponse(w, nil, http.StatusInternalServerError,
			fmt.Sprintf("error getting events in time range: %s", err.Error()))
		return
	}

	eventsResponse := repository.AppEventsToEventsResponseModel(filteredEvents)

	buildResponse(w, eventsResponse, http.StatusOK, "")

	return
}

func parseRequestParameters(r *http.Request) (*time.Time, *time.Time, error) {
	queryParams := r.URL.Query()

	startsAtString := queryParams.Get("starts_at")
	endsAtString := queryParams.Get("ends_at")

	if startsAtString == "" || endsAtString == "" {
		return nil, nil, fmt.Errorf("query parameters starts_at and ends_at are mandatory")
	}
	startsAt, err := time.Parse(time.RFC3339, startsAtString)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing starts_at query param: %w", err)
	}
	endsAt, err := time.Parse(time.RFC3339, endsAtString)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing ends_at query param: %w", err)
	}

	return &startsAt, &endsAt, nil
}

func buildResponse(w http.ResponseWriter, events []*repository.EventModel, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)

	response := NewEventsResponseModel(events,
		fmt.Sprintf("%d", statusCode),
		fmt.Sprintf(errorMessage))
	json.NewEncoder(w).Encode(response)
}
