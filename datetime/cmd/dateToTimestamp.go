//
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "golang/datetime/timeutil"
    "time"
)


func main() {

    var dateToConvert string
    var timezone string

    shortForm := "02/01/2006 15:04:05" // date composed by standard format constants

    flag.StringVar(&dateToConvert, "date", time.Now().Format(shortForm), "The date to convert to timestamp")
    flag.StringVar(&timezone, "tz", "Europe/Rome", "The timezone")
    flag.Parse()

    if len(dateToConvert) == 0 || len(timezone) == 0 {
        flag.PrintDefaults()
        os.Exit(1)
    }

    date, err := timeutil.ToTimestamp(dateToConvert, shortForm, timezone)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Location:", date.Location, "Time:", date.DateInLocation, "Timestamp:", date.Timestamp)
}
