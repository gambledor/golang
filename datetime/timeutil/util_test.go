// Package timeutl provides ...
package timeutil

import (
	"testing"
)

func TestValidateInputDateTimeFormatWithSuccess(t *testing.T) {
	t.Log("Test validate input date-time format with success")
	got := validateInputDateTimeFormat("17/05/2018 19:00:00")
	want := true
	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
	t.Log("Test validate input date-time format success")
}

func TestValidateInputDateTimeFormatWithError(t *testing.T) {
	t.Log("Test validate input date-time format with error")
	got := validateInputDateTimeFormat("2018/05/17 19:00:00")
	want := false
	if got != want {
		t.Errorf("got '%t', want '%t'", got, want)
	}
	t.Log("Test validate input date-time format with error")
}

func TestValidateInputTimestampWithSuccess(t *testing.T) {
	got := validateInputTimestamp("1527876016")
	want := true
	if got != want {
		t.Errorf("got '%t', want '%t'", got, want)
	}
}

func TestValidateInputTimestampWithError(t *testing.T) {
	got := validateInputTimestamp("152787601")
	want := false
	if got != want {
		t.Errorf("got '%t', want '%t'", got, want)
	}
}
