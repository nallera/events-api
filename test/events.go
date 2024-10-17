package test

import (
	"events-api/internal/events"
	"events-api/internal/events/adapter/repository"
)

func MakeAppEvents(noZones bool) []*events.Event {
	if noZones {
		return []*events.Event{
			MakeAppEvent(11, 22, nil),
			MakeAppEvent(12, 22, nil),
			MakeAppEvent(13, 22, nil),
		}
	}

	return []*events.Event{
		MakeAppEvent(11, 22, MakeAppZones()),
		MakeAppEvent(12, 22, MakeAppZones()),
		MakeAppEvent(13, 22, MakeAppZones()),
	}
}

func MakeAppEvent(eventID, baseEventID uint64, zones []*events.Zone) *events.Event {
	return &events.Event{
		EventID:        eventID,
		BaseEventID:    baseEventID,
		Title:          "test-title",
		EventStartDate: CreateTime(0, 0),
		EventEndDate:   CreateTime(1, 0),
		SellFrom:       CreateTime(-10, 0),
		SellTo:         CreateTime(-1, 0),
		SoldOut:        false,
		Zones:          zones,
		MinPrice:       10.0,
		MaxPrice:       30.0,
	}
}

func MakeAppZones() []*events.Zone {
	return []*events.Zone{
		MakeAppZoneModel(66, 10.0),
		MakeAppZoneModel(67, 30.0),
		MakeAppZoneModel(68, 18.0),
	}
}

func MakeAppZoneModel(zoneID uint64, price float64) *events.Zone {
	return &events.Zone{
		ZoneID:   zoneID,
		Capacity: 20,
		Price:    price,
		Name:     "test-name",
		Numbered: false,
	}
}

func MakeRestBaseEventsModel(baseEventID uint64, sellMode string) []*repository.RestBaseEventModel {
	return []*repository.RestBaseEventModel{
		MakeRestBaseEventModel(baseEventID, sellMode),
		MakeRestBaseEventModel(baseEventID, sellMode),
		MakeRestBaseEventModel(baseEventID, sellMode),
	}
}

func MakeRestBaseEventModel(baseEventID uint64, sellMode string) *repository.RestBaseEventModel {
	return &repository.RestBaseEventModel{
		BaseEventID:        baseEventID,
		SellMode:           sellMode,
		Title:              "test-title",
		OrganizerCompanyID: 987654,
		Events:             MakeRestEvents(),
	}
}

func MakeRestEvents() []*repository.RestEventModel {
	return []*repository.RestEventModel{
		MakeRestEventModel(11, MakeRestZones()),
		MakeRestEventModel(12, MakeRestZones()),
		MakeRestEventModel(13, MakeRestZones()),
	}
}

func MakeRestEventModel(eventID uint64, zones []*repository.RestZoneModel) *repository.RestEventModel {
	return &repository.RestEventModel{
		EventID:        eventID,
		EventStartDate: CreateTime(0, 0),
		EventEndDate:   CreateTime(1, 0),
		SellFrom:       CreateTime(-10, 0),
		SellTo:         CreateTime(-1, 0),
		SoldOut:        false,
		Zones:          zones,
	}
}

func MakeRestZones() []*repository.RestZoneModel {
	return []*repository.RestZoneModel{
		MakeRestZoneModel(66, 10.0),
		MakeRestZoneModel(67, 30.0),
		MakeRestZoneModel(68, 18.0),
	}
}

func MakeRestZoneModel(zoneID uint64, price float64) *repository.RestZoneModel {
	return &repository.RestZoneModel{
		ZoneID:   zoneID,
		Capacity: 20,
		Price:    price,
		Name:     "test-name",
		Numbered: false,
	}
}

func MakeSQLiteEvents() []*repository.SQLiteEventModel {
	return []*repository.SQLiteEventModel{
		MakeSQLiteEvent(11, 22),
		MakeSQLiteEvent(12, 22),
		MakeSQLiteEvent(13, 22),
	}
}

func MakeSQLiteEvent(eventID, baseEventID uint64) *repository.SQLiteEventModel {
	return &repository.SQLiteEventModel{
		EventID:        eventID,
		BaseEventID:    baseEventID,
		Title:          "test-title",
		EventStartDate: CreateTime(0, 0),
		EventEndDate:   CreateTime(1, 0),
		SellFrom:       CreateTime(-10, 0),
		SellTo:         CreateTime(-1, 0),
		SoldOut:        false,
		MinPrice:       10.0,
		MaxPrice:       30.0,
	}
}

func MakeResponseEvents() []*repository.EventModel {
	return []*repository.EventModel{
		MakeResponseEvent("B22-11"),
		MakeResponseEvent("B22-12"),
		MakeResponseEvent("B22-13"),
	}
}

func MakeResponseEvent(eventID string) *repository.EventModel {
	startDate := CreateTime(0, 0)
	endDate := CreateTime(1, 0)

	return &repository.EventModel{
		ID:        eventID,
		Title:     "test-title",
		StartDate: startDate.Format("2006-01-02"),
		StartTime: startDate.Format("15:04:05"),
		EndDate:   endDate.Format("2006-01-02"),
		EndTime:   endDate.Format("15:04:05"),
		MinPrice:  10.0,
		MaxPrice:  30.0,
	}
}
