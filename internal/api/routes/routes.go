package routes

import (
	"github.com/colevoss/temperature-blanket-backend/internal/api"
	"github.com/colevoss/temperature-blanket-backend/internal/api/handlers"
)

func RegisterRoutes(api *api.API, handlers *handlers.Handlers) {
	pingRoute := api.App.Group("/ping")
	{
		pingRoute.GET("", handlers.Ping.Ping)
	}

	timeSeriesRoute := api.App.Group("/timeseries")
	{
		timeSeriesRoute.GET("", handlers.TimeSeries.GetTimeSeries)
		timeSeriesRoute.GET("/summary", handlers.TimeSeries.GetSummary)
	}
}
