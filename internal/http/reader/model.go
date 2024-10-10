package reader

import "events-api/internal/events/adapter/repository"

type EventsResponseModel struct {
	Data  *DataModel  `json:"data"`
	Error *ErrorModel `json:"error"`
}

func NewEventsResponseModel(events []*repository.EventModel, errorCode, errorMessage string) *EventsResponseModel {
	var (
		dataModel  *DataModel
		errorModel *ErrorModel
	)

	if events != nil {
		dataModel = &DataModel{Events: events}
	}

	if errorCode != "" && errorCode != "200" {
		errorModel = &ErrorModel{Code: errorCode, Message: errorMessage}
	}

	return &EventsResponseModel{
		Data:  dataModel,
		Error: errorModel,
	}
}

type DataModel struct {
	Events []*repository.EventModel `json:"events"`
}

type ErrorModel struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
