package reader

import (
	"encoding/json"
	"fmt"
	"net/http"

	"events-api/internal/events"
	"events-api/internal/events/adapter/repository"
)

type Handler interface {
	GetEventsInTimeRange(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	providerXRepository events.Repository
}

func NewHTTPHandler(providerXRepository events.Repository) Handler {
	return &handler{
		providerXRepository: providerXRepository,
	}
}

func (h *handler) GetEventsInTimeRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()

	startsAt := queryParams.Get("starts_at")
	endsAt := queryParams.Get("ends_at")

	if startsAt == "" || endsAt == "" {
		buildResponse(w, nil, http.StatusBadRequest,
			"query parameters starts_at and ends_at are mandatory")

		return
	}

	currentEvents, err := h.providerXRepository.GetCurrentEvents()
	if err != nil {
		println(fmt.Sprintf("failed to get current events from provider x: %s", err.Error()))
		buildResponse(w, nil, http.StatusInternalServerError,
			fmt.Sprintf("failed to get current events from provider x: %s", err.Error()))

		return
	}

	eventsResponse := repository.AppEventsToEventsResponseModel(currentEvents)

	buildResponse(w, eventsResponse, http.StatusOK, "")

	return
}

func buildResponse(w http.ResponseWriter, events []*repository.EventModel, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)

	response := NewEventsResponseModel(events,
		fmt.Sprintf("%d", statusCode),
		fmt.Sprintf(errorMessage))
	json.NewEncoder(w).Encode(response)
}
