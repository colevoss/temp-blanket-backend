package weather

import (
	"context"
	"time"

	"github.com/colevoss/go-iem-sdk"
)

type DailySummary struct {
	Date    time.Time `json:"date"`
	High    float64   `json:"high"`
	Low     float64   `json:"low"`
	Average float64   `json:"average"`
}

type WeatherRepository interface {
	GetSummary(context context.Context, date time.Time, stationId string) (*DailySummary, error)
	GetNetworks(context context.Context) ([]*iem.Network, error)
	GetStations(ctx context.Context, networkId string) ([]*iem.Station, error)
	GetStation(ctx context.Context, stationId string) (*iem.Station, error)
	GetWeatherData(ctx context.Context, date time.Time, stationId string) ([]*iem.IEMWeatherData, error)
	Summary(dat []*iem.IEMWeatherData, date time.Time) (*DailySummary, error)
}
