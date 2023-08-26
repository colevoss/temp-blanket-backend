package synoptic

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/colevoss/temperature-blanket-backend/internal/log"
	"github.com/colevoss/temperature-blanket-backend/internal/util"
)

var SYNOPTIC_API_TOKEN string
var SYNOPTIC_API_URL = "https://api.synopticdata.com/v2/stations/timeseries"

type SynopticApi struct {
}

func NewSynopticApi() *SynopticApi {
	return &SynopticApi{}
}

func (s *SynopticApi) GetTimeSeriesTemperatureData(ctx context.Context, date time.Time) (*SynopticTimeSeriesResponse, error) {
	url, err := s.buildAPIUrl(ctx, date)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)

	if err != nil {
		log.C(ctx).Warnw(
			"Could not create request",
			"error", err,
		)

		return nil, err
	}

	log.C(ctx).Debugw(
		"Making request to Synoptic",
		"date", date,
	)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.C(ctx).Warnw(
			"Error making request to Synoptic",
			"error", err,
		)

		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Logger.Errorf("Error reading body %s", err)
		return nil, err
	}

	var timeseriesResponse SynopticTimeSeriesResponse
	err = json.Unmarshal(resBody, &timeseriesResponse)

	if err != nil {
		log.C(ctx).Error(err.Error())
		return nil, err
	}

	if timeseriesResponse.Summary.ResponseCode != 1 {
		return nil, SynopticApiError{
			Summary: *timeseriesResponse.Summary,
		}
	}

	log.C(ctx).Infow(
		"Synoptic request successfull",
		"date", date,
	)

	return &timeseriesResponse, nil
}

func (s *SynopticApi) buildAPIUrl(ctx context.Context, date time.Time) (*url.URL, error) {
	url, err := url.Parse(SYNOPTIC_API_URL)

	if err != nil {
		log.Raw().Errorf("Could not parse url %s", err)
		log.Logger.Error()
		return nil, err
	}

	query := url.Query()

	query.Set("token", SYNOPTIC_API_TOKEN)
	query.Set("stid", "klnk")
	query.Set("vars", "air_temp")

	start, err := util.GetStartOfDay(date)

	if err != nil {
		return nil, err
	}

	end, err := util.GetEndOfDay(date)

	if err != nil {
		return nil, err
	}

	startStr := util.FormatDate(start.UTC())
	endStr := util.FormatDate(end.UTC())

	log.C(ctx).Infow(
		"Building Synoptic request for dates",
		"start", startStr,
		"end", endStr,
	)

	query.Set("start", startStr)
	query.Set("end", endStr)

	url.RawQuery = query.Encode()

	return url, nil
}

func init() {
	SYNOPTIC_API_TOKEN = os.Getenv("SYNOPTIC_API_TOKEN")
}
