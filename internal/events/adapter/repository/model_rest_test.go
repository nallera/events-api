package repository_test

import (
	"events-api/internal/events"
	"events-api/internal/events/adapter/repository"
	"events-api/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRest_RestBaseEventModelToAppEvents(t *testing.T) {
	testRestBaseEvent := test.MakeRestBaseEventsModel(22, "online")
	var testAppEvents []*events.Event
	testAppEvents = append(testAppEvents, test.MakeAppEvents(false)...)
	testAppEvents = append(testAppEvents, test.MakeAppEvents(false)...)
	testAppEvents = append(testAppEvents, test.MakeAppEvents(false)...)

	result := repository.RestBaseEventModelToAppEvents(testRestBaseEvent)

	assert.Equal(t, testAppEvents, result)
}

func TestRest_RestBaseEventModelToAppEvents_DiscardNotOnlineSellMode(t *testing.T) {
	testRestBaseEvent := test.MakeRestBaseEventsModel(22, "online")
	testRestBaseEvent[2].SellMode = "offline"
	var testAppEvents []*events.Event
	testAppEvents = append(testAppEvents, test.MakeAppEvents(false)...)
	testAppEvents = append(testAppEvents, test.MakeAppEvents(false)...)

	result := repository.RestBaseEventModelToAppEvents(testRestBaseEvent)

	assert.Equal(t, testAppEvents, result)
}

func TestRest_RestBaseEventModelToAppEvents_DiscardInvalidDate(t *testing.T) {
	testRestBaseEvent := test.MakeRestBaseEventsModel(22, "online")
	testRestBaseEvent[2].Events[2].EventStartDate = time.Time{}
	var testAppEvents []*events.Event
	testAppEvents = append(testAppEvents, test.MakeAppEvents(false)...)
	testAppEvents = append(testAppEvents, test.MakeAppEvents(false)...)
	modifiedEvents := test.MakeAppEvents(false)
	modifiedEvents = modifiedEvents[:2]
	testAppEvents = append(testAppEvents, modifiedEvents...)

	result := repository.RestBaseEventModelToAppEvents(testRestBaseEvent)

	assert.Equal(t, testAppEvents, result)
}

func TestRest_AppEventsToEventsResponseModel(t *testing.T) {
	testAppEvents := test.MakeAppEvents(false)
	testEventsResponseModel := test.MakeResponseEvents()

	result := repository.AppEventsToEventsResponseModel(testAppEvents)

	assert.Equal(t, testEventsResponseModel, result)
}
