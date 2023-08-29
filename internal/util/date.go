package util

import (
	"log"
	"time"
)

func GetStartEndOfDay(date *time.Time) (time.Time, time.Time) {
	tz, err := time.LoadLocation("America/Chicago")

	if err != nil {
		log.Fatalf("Cannot load timezone %s", err)
	}

	start := time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		0, 0, 0, 0, tz,
	)

	end := start.Add(time.Hour * 23).Add(time.Minute * 59)

	return start, end
}

func GetStartOfDay(date time.Time) (time.Time, error) {
	tz, err := time.LoadLocation("America/Chicago")

	if err != nil {
		log.Fatalf("Cannot load timezone %s", err)
		return date, err
	}

	date = date.In(tz)

	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		0, 0, 0, 0,
		tz,
	), nil
}

func GetEndOfDay(date time.Time) (time.Time, error) {
	tz, err := time.LoadLocation("America/Chicago")

	if err != nil {
		log.Fatalf("Cannot load timezone %s", err)
		return date, err
	}

	date = date.In(tz)

	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		23, 59, 59, 0,
		tz,
	), nil
}

func FormatDate(date time.Time) string {
	return date.Format("200601021504")
}
