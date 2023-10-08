package handlers

import (
	"github.com/colevoss/temperature-blanket-backend/internal/api"
	"github.com/colevoss/temperature-blanket-backend/internal/repositories"
)

type Handlers struct {
	Ping       *PingHandlers
	TimeSeries *TimeSeriesHandlers
	Networks   *NetworkHandlers
}

func NewHandlers(api *api.API, repos *repositories.Repositories) *Handlers {
	return &Handlers{
		Ping:       NewPingHandlers(),
		TimeSeries: NewTimeSeriesHandlers(repos),
		Networks:   NewNetworkHandler(repos),
	}
}
