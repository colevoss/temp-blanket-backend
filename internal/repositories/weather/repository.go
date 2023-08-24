package weather

import "time"

type DailySummary struct {
	Date    time.Time `json:"date"`
	High    float64   `json:"high"`
	Low     float64   `json:"low"`
	Average float64   `json:"average"`
}

type WeatherRepository interface {
	GetSummary(date time.Time) (*DailySummary, error)
}
