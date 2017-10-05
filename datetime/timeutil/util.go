// Package util provides utility functions for my own date time management
package timeutil

import (
    "time"
)
type MyDate struct {
    Location        *time.Location
    DateInLocation  time.Time
    Timestamp       int64
}

// ToTimestamp 
func ToTimestamp(dateToConvert, dateForm, timezone string) (MyDate, error) {
    location, err := time.LoadLocation(timezone)
    if err != nil {
        return MyDate{}, err
    }

    var convertedDate time.Time
    convertedDate, err = time.ParseInLocation(dateForm, dateToConvert, location)
    if err != nil {
        return MyDate{}, err
    }

    return MyDate{location, convertedDate, convertedDate.Unix()}, nil
}

