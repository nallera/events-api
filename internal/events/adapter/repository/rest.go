package repository

import (
	"events-api/internal/events"
	"events-api/pkg/server"
	"fmt"
)

type rest struct {
	client server.RestClient
}

func NewRestRepository(client server.RestClient) events.Repository {
	return &rest{
		client: client,
	}
}

func (r *rest) GetCurrentEvents() ([]*events.Event, error) {
	response := new(RestEventListModel)

	err := r.client.Get("get_current_events", nil, nil, response, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get current events from provider x: %s", err.Error())
	}

	appEvents := RestBaseEventModelToAppEvents(response.Output.BaseEvents)

	return appEvents, nil
}
