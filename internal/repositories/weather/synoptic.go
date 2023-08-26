package weather

import (
	"context"
	"log"
	"time"

	"github.com/colevoss/temperature-blanket-backend/internal/integrations/synoptic"
	"github.com/colevoss/temperature-blanket-backend/internal/util"
)

type SynopticWeatherRepo struct {
	api *synoptic.SynopticApi
}

func (swr *SynopticWeatherRepo) GetSummary(ctx context.Context, time time.Time) (*DailySummary, error) {
	timeSeries, err := swr.api.GetTimeSeriesTemperatureData(ctx, time)

	if err != nil {
		log.Printf("Error %v", err)
		return nil, err
	}

	low := 0.0
	high := 0.0

	station := timeSeries.Station[0]
	temps := station.Observations.AirTemp

	for i, temp := range temps {
		fTemp := util.CelciusToFarenheit(temp)

		if i == 0 {
			high = fTemp
			low = fTemp
		}

		if fTemp > high {
			high = fTemp
		}

		if fTemp < low {
			low = fTemp
		}
	}

	average := (high + low) / 2

	summaryTime, _ := util.GetStartOfDay(time)

	return &DailySummary{
		Low:     low,
		High:    high,
		Average: average,
		Date:    summaryTime,
	}, nil
}

func NewSynopticWeatherRepo() *SynopticWeatherRepo {
	return &SynopticWeatherRepo{
		api: synoptic.NewSynopticApi(),
	}
}
