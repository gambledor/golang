// Package timeutil provides utility functions for my own date time management
package timeutil

import (
	"errors"
	"regexp"
	"time"
)

type myDate struct {
	Location       *time.Location
	DateInLocation time.Time
	Timestamp      int64
}

// ToTimestamp converts a dateTime string format dd/mm/YYYY HH:MM:ii in a formatted dateForm with timezone
func ToTimestamp(dateTimeToConvert, dateForm, timezone string) (myDate, error) {
	if !validateInputDateTimeFormat(dateTimeToConvert) {
		return myDate{}, errors.New("Wrong date time input form")
	}

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
	if !validateInputDateTimeFormat(dateTimeToConvert) {
		return myDate{}, errors.New("Wrong date time input form")
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return myDate{}, err
	}

	now := time.Now().In(location)

	return myDate{location, now, now.Unix()}, nil
}

func validateInputDateTimeFormat(dateTimeToConvert string) bool {
	// format: dd\mm\YYYY HH:MM:ii
	myRegexp := "\\d\\d/\\d\\d/\\d\\d\\d\\d \\d\\d:\\d\\d:\\d\\d"
	r := regexp.MustCompile(myRegexp)

	return r.MatchString(dateTimeToConvert)
}

