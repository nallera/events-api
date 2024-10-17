package events

import (
	"fmt"
	"time"
)

type Event struct {
	EventID        uint64
	BaseEventID    uint64
	Title          string
	EventStartDate time.Time
	EventEndDate   time.Time
	SellFrom       time.Time
	SellTo         time.Time
	SoldOut        bool
	Zones          []*Zone
	MinPrice       float64
	MaxPrice       float64
}

func (f *Event) String() string {
	return fmt.Sprintf("{EventID:%d BaseEventID:%d Title:%s EventStartDate:%+v EventEndDate:%+v SellFrom:%+v "+
		"SellTo:%+v SoldOut:%v Zones:%+v MinPrice:%.2f MaxPrice:%.2f}",
		f.EventID, f.BaseEventID, f.Title, f.EventStartDate, f.EventEndDate, f.SellFrom, f.SellTo, f.SoldOut, f.Zones,
		f.MinPrice, f.MaxPrice)
}

type Zone struct {
	ZoneID   uint64
	Capacity uint64
	Price    float64
	Name     string
	Numbered bool
}

func (f *Zone) String() string {
	return fmt.Sprintf("{ZoneID:%d Capacity:%d Price:%.2f Name:%s Numbered: %v}",
		f.ZoneID, f.Capacity, f.Price, f.Name, f.Numbered)
}
