package main

import (
	"github.com/colevoss/temperature-blanket-backend/internal/api"
	"github.com/colevoss/temperature-blanket-backend/internal/api/handlers"
	"github.com/colevoss/temperature-blanket-backend/internal/api/routes"
	"github.com/colevoss/temperature-blanket-backend/internal/config"
	"github.com/colevoss/temperature-blanket-backend/internal/log"
	"github.com/colevoss/temperature-blanket-backend/internal/repositories"
)

func main() {
	cfg := config.NewConfig()
	cfg.ParseFlags()

	log.InitLogger(cfg)
	defer log.CloseLogger()

	api := api.NewApi(cfg)
	api.Init()

	repos := repositories.NewRepositories(cfg)

	// weatherRepo := weather.NewMockWeatherRepo()

	handlers := handlers.NewHandlers(api, repos)

	routes.RegisterRoutes(api, handlers)

	api.Run()
}
