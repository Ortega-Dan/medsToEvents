package main

import (
	"os"
	"strconv"
	"time"

	ics "github.com/arran4/golang-ical"
)

func main() {

	if len(os.Args) < 4 {
		println("Usage: medtocal <medication-name> <dose-per-take> <hours-interval> <total-days>")
		os.Exit(1)
	}

	// input variables
	medicationName := os.Args[1]
	dosePerTake := os.Args[2]
	hoursInterval, _ := strconv.Atoi(os.Args[3])
	totalDays, _ := strconv.Atoi(os.Args[4])

	// generate calendar
	calendar := ics.NewCalendar()

	// start and end dates
	start := time.Now()

	for i := 0; i < 24/hoursInterval; i++ {
		dateTime := start.Add(time.Hour * time.Duration(hoursInterval*i))

		// generate event
		event := calendar.AddEvent(medicationName+dateTime.Format("20060102T150405"))
		event.SetSummary(medicationName+"-"+dosePerTake+"-"+strconv.Itoa(hoursInterval)+"hrs")
		event.SetStartAt(dateTime)
		event.SetEndAt(dateTime.Add(time.Minute * 30))
		event.AddRrule("FREQ=DAILY;COUNT=" + strconv.Itoa(totalDays))
	}

	// event.SetLocation()

	// write calendar to stdout with calendar.SerializeTo
	writer, err := os.Create(medicationName + ".ics")
	if err != nil {
		panic(err)
	}

	calendar.SerializeTo(writer)

}
