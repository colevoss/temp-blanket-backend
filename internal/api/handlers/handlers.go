package handlers

import (
	"github.com/colevoss/temperature-blanket-backend/internal/api"
	"github.com/colevoss/temperature-blanket-backend/internal/repositories/weather"
)

type Handlers struct {
	Ping       *PingHandlers
	TimeSeries *TimeSeriesHandlers
}

func NewHandlers(api *api.API, weatherRepo weather.WeatherRepository) *Handlers {
	return &Handlers{
		Ping:       NewPingHandlers(),
		TimeSeries: NewTimeSeriesHandlers(weatherRepo),
	}
}
