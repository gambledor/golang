// Package util provides utility functions for my own date time management
package timeutil

import (
	"time"
)

type myDate struct {
	Location       *time.Location
	DateInLocation time.Time
	Timestamp      int64
}

// ToTimestamp converts a dateTime in a formatted dateForm in the timezone
func ToTimestamp(dateTimeToConvert, dateForm, timezone string) (myDate, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return myDate{}, err
	}

	var convertedDate time.Time
	convertedDate, err = time.ParseInLocation(dateForm, dateTimeToConvert, location)
	if err != nil {
		return myDate{}, err
	}

	return myDate{location, convertedDate, convertedDate.Unix()}, nil
}

// NowInTimezone return the datetime in the timezone
func NowInTimezone(dateForm, timezone string) (myDate, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return myDate{}, err
	}

	now := time.Now().In(location)

	return myDate{location, now, now.Unix()}, nil
}
