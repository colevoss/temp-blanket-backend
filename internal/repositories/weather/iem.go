package weather

import (
	"context"
	"errors"
	"time"

	"github.com/colevoss/go-iem-sdk"
	"github.com/colevoss/temperature-blanket-backend/internal/cache"
	"github.com/colevoss/temperature-blanket-backend/internal/log"
	"github.com/colevoss/temperature-blanket-backend/internal/response"
)

type IEMWeatherRepo struct {
	client *iem.Client
	cache  *cache.Cache
}

func (r *IEMWeatherRepo) checkError(err error) error {
	var notFoundError iem.IEMNotFoundError

	if errors.As(err, &notFoundError) {
		return response.NewError(notFoundError.Detail, notFoundError.Code, nil, notFoundError)
	} else {
		return err
	}
}

func (r *IEMWeatherRepo) GetNetworks(ctx context.Context) ([]*iem.Network, error) {
	log.C(ctx).Infow("Requesting networks from IEM")

	networks, err := r.client.Networks().GetNetworks(ctx)

	return networks, err
}

func (r *IEMWeatherRepo) GetStations(ctx context.Context, networkId string) ([]*iem.Station, error) {
	log.C(ctx).Infow(
		"Requesting stations from IEM for network",
		"networkId", networkId,
	)

	stations, err := r.client.Stations().GetStations(ctx, networkId)

	return stations, err
}

func (r *IEMWeatherRepo) Summary(data []*iem.IEMWeatherData, date time.Time) (*DailySummary, error) {
	low := 0.0
	high := 0.0

	for i, d := range data {
		if i == 0 {
			high = d.TemperatureF
			low = d.TemperatureF
		}

		if d.TemperatureF > high {
			high = d.TemperatureF
		}

		if d.TemperatureF < low {
			low = d.TemperatureF
		}
	}

	average := (high + low) / 2

	return &DailySummary{
		Low:     low,
		High:    high,
		Average: average,
		Date:    date,
	}, nil
}

func (r *IEMWeatherRepo) GetWeatherData(ctx context.Context, date time.Time, stationId string) ([]*iem.IEMWeatherData, error) {
	station, err := r.GetStation(ctx, stationId)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	tz, err := time.LoadLocation(station.Timezone)

	if err != nil {
		return nil, err
	}

	q := iem.NewWeatherDataQuery()
	localDate := date.In(tz)

	q.Data(iem.TempF)
	q.Stations(stationId)
	q.Start(localDate)
	q.End(localDate)
	q.ReportType(3)
	q.Timezone(station.Timezone)

	log.C(ctx).Infow(
		"Requesting IEM Weather data",
		"station", "LNK",
		"date", date,
		"station", stationId,
		"timezone", tz,
		"localDate", localDate,
	)

	data, err := r.client.Weather().Get(ctx, q)

	if err != nil {
		log.C(ctx).Errorw(
			"Error requesting weather from IEM",
			"error", err,
		)
		return nil, err
	}

	log.C(ctx).Infow(
		"IEM weather request successful",
		"count", len(data),
	)

	return data, nil
}

func (r *IEMWeatherRepo) GetSummary(ctx context.Context, date time.Time, stationId string) (*DailySummary, error) {
	station, err := r.GetStation(ctx, stationId)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	tz, err := time.LoadLocation(station.Timezone)

	if err != nil {
		return nil, err
	}

	q := iem.NewWeatherDataQuery()
	localDate := date.In(tz)

	q.Data(iem.TempF)
	q.Stations(stationId)
	q.Start(localDate)
	q.End(localDate)
	q.ReportType(3)
	q.Timezone(station.Timezone)

	log.C(ctx).Infow(
		"Requesting IEM Weather data",
		"station", "LNK",
		"date", date,
		"station", stationId,
		"timezone", tz,
		"localDate", localDate,
	)

	data, err := r.client.Weather().Get(ctx, q)

	if err != nil {
		log.C(ctx).Errorw(
			"Error requesting weather from IEM",
			"error", err,
		)
		return nil, err
	}

	log.C(ctx).Infow(
		"IEM weather request successful",
		"count", len(data),
	)

	low := 0.0
	high := 0.0

	for i, d := range data {
		if i == 0 {
			high = d.TemperatureF
			low = d.TemperatureF
		}

		if d.TemperatureF > high {
			high = d.TemperatureF
		}

		if d.TemperatureF < low {
			low = d.TemperatureF
		}
	}

	average := (high + low) / 2

	return &DailySummary{
		Low:     low,
		High:    high,
		Average: average,
		Date:    localDate,
	}, nil
}

func (r *IEMWeatherRepo) GetStation(ctx context.Context, stationId string) (*iem.Station, error) {
	cachedStation, ok := r.cache.GetAndRefresh(stationId)

	if ok {
		log.C(ctx).Infow(
			"Station retrieved from cache",
			"stationId", stationId,
		)

		return cachedStation.(*iem.Station), nil
	}

	log.C(ctx).Infow(
		"Requesting station from IEM",
		"stationId", stationId,
	)

	station, err := r.client.Stations().GetStation(ctx, stationId)

	if err != nil {
		return nil, r.checkError(err)
	}

	log.C(ctx).Infow(
		"Station requested from IEM successfully",
		"stationId", station.Id,
		"name", station.Name,
	)

	didSet := r.cache.Set(stationId, station)

	if !didSet {
		log.C(ctx).Infow(
			"Could not set station in cache. Too many items",
			"stationId", station.Id,
		)
	}

	return station, nil
}

func NewIEMWeatherRepo() *IEMWeatherRepo {
	cache := cache.NewCache(time.Duration(time.Minute*5), time.Duration(time.Minute*1), 50)
	cache.OnEvict(onEvict)
	cache.Run()

	return &IEMWeatherRepo{
		client: iem.NewClient(),
		cache:  cache,
	}
}

func onEvict(items map[string]interface{}) {
	keys := make([]string, len(items))

	i := 0
	for k := range items {
		keys[i] = k
		i++
	}

	log.Raw().Infow(
		"Stations purged from cache",
		"count", len(items),
		"keys", keys,
	)
}
