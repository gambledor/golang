// Package util provides utility functions for my own date time management
package timeutil

import (
	"time"
)

type MyDate struct {
	Location       *time.Location
	DateInLocation time.Time
	Timestamp      int64
}

// ToTimestamp converts a dateTime in a formatted dateForm in the timezone
func ToTimestamp(dateTimeToConvert, dateForm, timezone string) (MyDate, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return MyDate{}, err
	}

	var convertedDate time.Time
	convertedDate, err = time.ParseInLocation(dateForm, dateTimeToConvert, location)
	if err != nil {
		return MyDate{}, err
	}

	return MyDate{location, convertedDate, convertedDate.Unix()}, nil
}
