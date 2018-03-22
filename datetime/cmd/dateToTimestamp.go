//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gambledor/golang/datetime/timeutil"
)

func main() {

	var (
		command,
		dateToConvert,
		timezone string
	)

	shortForm := "02/01/2006 15:04:05" // date composed by standard format constants

	flag.StringVar(&command, "command", "timestamp", "[timestamp|timezone]")
	flag.StringVar(&dateToConvert, "date", time.Now().Format(shortForm), "the date to convert to timestamp")
	flag.StringVar(&timezone, "tz", "Europe/Rome", "The timezone")
	flag.Parse()

	if len(dateToConvert) == 0 || len(timezone) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

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
	}

}
