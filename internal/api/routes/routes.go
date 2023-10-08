package routes

import (
	"github.com/colevoss/temperature-blanket-backend/internal/api"
	"github.com/colevoss/temperature-blanket-backend/internal/api/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(api *api.API, handlers *handlers.Handlers) {
	api.App.Route("/ping", func(r chi.Router) {
		r.Get("/", handlers.Ping.Ping)
	})

	api.App.Route("/weather", func(r chi.Router) {
		r.Get("/summary", handlers.TimeSeries.GetSummary)
	})

	api.App.Route("/networks", func(r chi.Router) {
		r.Get("/", handlers.Networks.GetNetworks)
		r.Get("/{networkId}/stations", handlers.Networks.GetStations)
	})
}
