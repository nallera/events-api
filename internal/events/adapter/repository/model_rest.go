package repository

import (
	"encoding/xml"
	"fmt"
	"time"

	"events-api/internal/events"
)

const SellModeOnline = "online"

type RestEventListModel struct {
	Output RestOutputModel `xml:"output"`
}

type RestOutputModel struct {
	BaseEvents []*RestBaseEventModel `xml:"base_event"`
}

type RestBaseEventModel struct {
	BaseEventID        uint64            `xml:"base_event_id,string,attr"`
	SellMode           string            `xml:"sell_mode,attr"`
	Title              string            `xml:"title,attr"`
	OrganizerCompanyID uint64            `xml:"organizer_company_id,string,attr"`
	Events             []*RestEventModel `xml:"event"`
}

func (f *RestBaseEventModel) String() string {
	return fmt.Sprintf("{BaseEventID:%d SellMode:%s Title:%s OrganizerCompanyID:%d Events:%+v}",
		f.BaseEventID, f.SellMode, f.Title, f.OrganizerCompanyID, f.Events)
}

type RestEventModel struct {
	EventID        uint64           `xml:"event_id,string,attr"`
	EventStartDate time.Time        `xml:"event_start_date,attr"`
	EventEndDate   time.Time        `xml:"event_end_date,attr"`
	SellFrom       time.Time        `xml:"sell_from,attr"`
	SellTo         time.Time        `xml:"sell_to,attr"`
	SoldOut        bool             `xml:"sold_out,string,attr"`
	Zone           []*RestZoneModel `xml:"zone"`
}

func (f *RestEventModel) String() string {
	return fmt.Sprintf("{EventID:%d EventStartDate:%+v EventEndDate:%+v SellFrom:%+v SellTo:%+v SoldOut:%v Zone:%+v}",
		f.EventID, f.EventStartDate, f.EventEndDate, f.SellFrom, f.SellTo, f.SoldOut, f.Zone)
}

func (e *RestEventModel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	aux := &struct {
		EventID        uint64           `xml:"event_id,string,attr"`
		EventStartDate string           `xml:"event_start_date,attr"`
		EventEndDate   string           `xml:"event_end_date,attr"`
		SellFrom       string           `xml:"sell_from,attr"`
		SellTo         string           `xml:"sell_to,attr"`
		SoldOut        bool             `xml:"sold_out,string,attr"`
		Zone           []*RestZoneModel `xml:"zone"`
	}{}

	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}

	var err error
	e.EventStartDate, err = time.Parse("2006-01-02T15:04:05", aux.EventStartDate)
	if err != nil {
		e.EventStartDate = time.Time{}
	}
	e.EventEndDate, err = time.Parse("2006-01-02T15:04:05", aux.EventEndDate)
	if err != nil {
		e.EventStartDate = time.Time{}
	}
	e.SellFrom, err = time.Parse("2006-01-02T15:04:05", aux.SellFrom)
	if err != nil {
		e.EventStartDate = time.Time{}
	}
	e.SellTo, err = time.Parse("2006-01-02T15:04:05", aux.SellTo)
	if err != nil {
		e.EventStartDate = time.Time{}
	}

	e.EventID = aux.EventID
	e.SoldOut = aux.SoldOut
	e.Zone = aux.Zone

	return nil
}

type RestZoneModel struct {
	ZoneID   uint64  `xml:"zone_id,string,attr"`
	Capacity uint64  `xml:"capacity,string,attr"`
	Price    float64 `xml:"price,string,attr"`
	Name     string  `xml:"name,attr"`
	Numbered bool    `xml:"numbered,string,attr"`
}

func (f *RestZoneModel) String() string {
	return fmt.Sprintf("{ZoneID:%d Capacity:%d Price:%.2f Name:%s Numbered: %v}",
		f.ZoneID, f.Capacity, f.Price, f.Name, f.Numbered)
}

func RestBaseEventModelToAppEvents(restModel []*RestBaseEventModel) []*events.Event {
	var result []*events.Event

	for _, restBaseEvent := range restModel {
		if restBaseEvent.SellMode != SellModeOnline {
			println(fmt.Sprintf("discarding base event %d because sell mode is not online", restBaseEvent.BaseEventID))
			continue
		}
		for _, restEvent := range restBaseEvent.Events {
			var appZones []*events.Zone

			minPrice, maxPrice := -1.0, 0.0

			for _, restZone := range restEvent.Zone {
				if restZone.Price < minPrice || minPrice == -1.0 {
					minPrice = restZone.Price
				}
				if restZone.Price > maxPrice {
					maxPrice = restZone.Price
				}

				appZones = append(appZones, &events.Zone{
					ZoneID:   restZone.ZoneID,
					Capacity: restZone.Capacity,
					Price:    restZone.Price,
					Name:     restZone.Name,
					Numbered: restZone.Numbered,
				})
			}

			if datesAreValid(restEvent) {
				result = append(result, &events.Event{
					EventID:        restEvent.EventID,
					BaseEventID:    restBaseEvent.BaseEventID,
					Title:          restBaseEvent.Title,
					EventStartDate: restEvent.EventStartDate,
					EventEndDate:   restEvent.EventEndDate,
					SellFrom:       restEvent.SellFrom,
					SellTo:         restEvent.SellTo,
					SoldOut:        restEvent.SoldOut,
					Zone:           appZones,
					MinPrice:       minPrice,
					MaxPrice:       maxPrice,
				})
			} else {
				println(fmt.Sprintf("the event %d has an invalid date", restEvent.EventID))
			}
		}
	}

	return result
}

func datesAreValid(restEvent *RestEventModel) bool {
	return !restEvent.EventStartDate.IsZero() && !restEvent.EventEndDate.IsZero() &&
		!restEvent.SellFrom.IsZero() && !restEvent.SellTo.IsZero()
}

type EventModel struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	StartDate string  `json:"start_date"`
	StartTime string  `json:"start_time"`
	EndDate   string  `json:"end_date"`
	EndTime   string  `json:"end_time"`
	MinPrice  float64 `json:"min_price"`
	MaxPrice  float64 `json:"max_price"`
}

func AppEventsToEventsResponseModel(events []*events.Event) []*EventModel {
	var result []*EventModel

	for _, event := range events {
		result = append(result, &EventModel{
			ID:        fmt.Sprintf("B%d-%d", event.BaseEventID, event.EventID),
			Title:     event.Title,
			StartDate: event.EventStartDate.Format("2006-01-02"),
			StartTime: event.EventStartDate.Format("15:04:05"),
			EndDate:   event.EventEndDate.Format("2006-01-02"),
			EndTime:   event.EventEndDate.Format("15:04:05"),
			MinPrice:  event.MinPrice,
			MaxPrice:  event.MaxPrice,
		})
	}

	return result
}
