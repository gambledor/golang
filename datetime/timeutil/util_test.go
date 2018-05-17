// Package timeutl provides ...
package timeutil

import (
	"testing"
)

func TestValidateInputDateTimeFormatWithSuccess(t *testing.T) {
	got := validateInputDateTimeFormat("17/05/2018 19:00:00")
	want := true
	if got != want {
		t.Errorf("got '%t' want '%t'", got, want)
	}
}

func TestValidateInputDateTimeFormatWithError(t *testing.T) {
	got := validateInputDateTimeFormat("2018/05/17 19:00:00")
	want := false
	if got != want {
		t.Errorf("got '%t', want '%t'", got, want)
	}
}
