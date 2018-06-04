//
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gambledor/golang/datetime/timeutil"
)

var (
	command,
	dateToConvert,
	timezone string
)

func init() {
	flag.StringVar(&command, "command", "timestamp", "[timestamp|timezone|datetime]")
	flag.StringVar(&dateToConvert, "date", time.Now().Format(shortForm), "the date to convert to timestamp")
	flag.StringVar(&timezone, "tz", "Europe/Rome", "The timezone")
}

func main() {

	shortForm := "02/01/2006 15:04:05" // date composed by standard format constants
	flag.Parse()
	if command == "timestamp" {
		date, err := timeutil.ToTimestamp(dateToConvert, shortForm, timezone)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("location:", date.Location, "time:", date.DateInLocation, "timestamp:", date.Timestamp)
	} else if command == "timezone" {
		date, err := timeutil.NowInTimezone(shortForm, timezone)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("location:", date.Location, "time:", date.DateInLocation, "timestamp:", date.Timestamp)
	} else if command == "datetime" {
		date, err := timeutil.TimestampToDateTime(dateToConvert, shortForm, timezone)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("location:", date.Location, "date-time:", date.DateInLocation, "timestamp:", date.Timestamp)
	}
}
