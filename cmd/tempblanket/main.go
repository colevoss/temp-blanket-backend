package main

import (
	"github.com/colevoss/temperature-blanket-backend/internal/api"
	"github.com/colevoss/temperature-blanket-backend/internal/api/handlers"
	"github.com/colevoss/temperature-blanket-backend/internal/api/routes"
	"github.com/colevoss/temperature-blanket-backend/internal/repositories/weather"
	"github.com/gin-gonic/gin"
)

func main() {
	api := api.NewApi()
	api.Init()

	weatherRepo := weather.NewSynopticWeatherRepo()
	// weatherRepo := weather.NewMockWeatherRepo()

	handlers := handlers.NewHandlers(api, weatherRepo)

	routes.RegisterRoutes(api, handlers)

	api.Run()
}

func init() {
	gin.SetMode(gin.DebugMode)
}
