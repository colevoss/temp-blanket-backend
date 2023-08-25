package main

import (
	"github.com/colevoss/temperature-blanket-backend/internal/api"
	"github.com/colevoss/temperature-blanket-backend/internal/api/handlers"
	"github.com/colevoss/temperature-blanket-backend/internal/api/routes"
	"github.com/colevoss/temperature-blanket-backend/internal/config"
	"github.com/colevoss/temperature-blanket-backend/internal/logger"
	"github.com/colevoss/temperature-blanket-backend/internal/repositories/weather"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()
	cfg.ParseFlags()

	if cfg.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	logger.InitLogger(cfg)
	defer logger.CloseLogger()

	api := api.NewApi(cfg)
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
