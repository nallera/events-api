package test

import "time"

func CreateTime(dayDisplacement, hourDisplacement int) time.Time {
	return time.Date(2018, time.December, 9+dayDisplacement, hourDisplacement, 0, 0, 0, time.UTC)
}
