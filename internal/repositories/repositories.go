package repositories

import (
	"github.com/colevoss/temperature-blanket-backend/internal/config"
	"github.com/colevoss/temperature-blanket-backend/internal/repositories/weather"
)

type Repositories struct {
	Weather weather.WeatherRepository
}

func NewRepositories(cfg *config.Config) *Repositories {
	return &Repositories{
		// Weather: weather.NewSynopticWeatherRepo(),
		Weather: weather.NewIEMWeatherRepo(),
	}
}
