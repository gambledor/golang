// Package timeutil provides utility functions for my own date time management
package timeutil

import (
	"errors"
	"regexp"
	"strconv"
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
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return myDate{}, err
	}

	now := time.Now().In(location)

	return myDate{location, now, now.Unix()}, nil
}

// TimestampToDateTime converts a timestamp to a date time in the timezone
func TimestampToDateTime(timestampstr, dateForm, timezone string) (myDate, error) {
	if !validateInputTimestamp(timestampstr) {
		return myDate{}, errors.New("Wrong timestamp value provided")
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return myDate{}, err
	}
	timestamp, err := strconv.ParseInt(timestampstr, 10, 64)
	if err != nil {
		return myDate{}, err
	}

	time := time.Unix(timestamp, 0).In(location)

	return myDate{location, time, timestamp}, nil
}

func validateInputDateTimeFormat(dateTimeToConvert string) bool {
	// format: dd\mm\YYYY HH:MM:ii
	dateTimeRegex := "\\d\\d/\\d\\d/\\d\\d\\d\\d \\d\\d:\\d\\d:\\d\\d"
	r1 := regexp.MustCompile(dateTimeRegex)

	return r1.MatchString(dateTimeToConvert)
}

func validateInputTimestamp(timestamp string) bool {
	timestampRegex := `^1\d{9}`
	r2 := regexp.MustCompile(timestampRegex)

	return r2.MatchString(timestamp)
}
