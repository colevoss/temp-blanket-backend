package synoptic

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/colevoss/temperature-blanket-backend/internal/util"
)

var SYNOPTIC_API_TOKEN string
var SYNOPTIC_API_URL = "https://api.synopticdata.com/v2/stations/timeseries"

type SynopticApi struct {
}

func NewSynopticApi() *SynopticApi {
	return &SynopticApi{}
}

func (s *SynopticApi) GetTimeSeriesTemperatureData(date time.Time) (*SynopticTimeSeriesResponse, error) {
	url, err := s.buildAPIUrl(date)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)

	if err != nil {
		log.Printf("Could not create request %s", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("Error making request: %s", err)
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Printf("Error reading body %s", err)
		return nil, err
	}

	var timeseriesResponse SynopticTimeSeriesResponse
	err = json.Unmarshal(resBody, &timeseriesResponse)

	if err != nil {
		log.Printf("Could not parse response body %s", err)
		return nil, err
	}

	log.Printf("TSR %v", timeseriesResponse.Summary.ResponseCode)

	if timeseriesResponse.Summary.ResponseCode != 1 {
		return nil, SynopticApiError{
			Summary: *timeseriesResponse.Summary,
		}
	}

	return &timeseriesResponse, nil
}

func (s *SynopticApi) buildAPIUrl(date time.Time) (*url.URL, error) {
	url, err := url.Parse(SYNOPTIC_API_URL)

	if err != nil {
		log.Printf("Could not parse url %s", err)
		return nil, err
	}

	query := url.Query()

	query.Set("token", SYNOPTIC_API_TOKEN)
	query.Set("stid", "klnk")
	query.Set("vars", "air_temp")

	start, err := util.GetStartOfDay(date)

	if err != nil {
		log.Printf("Fuck")
		return nil, err
	}

	end, err := util.GetEndOfDay(date)

	if err != nil {
		log.Printf("Fuck")
		return nil, err
	}

	startStr := util.FormatDate(start.UTC())
	endStr := util.FormatDate(end.UTC())

	log.Printf("Making request for %v - %v", startStr, endStr)

	query.Set("start", startStr)
	query.Set("end", endStr)

	url.RawQuery = query.Encode()

	log.Printf("Making request to url: %s", url.String())

	return url, nil
}

func init() {
	SYNOPTIC_API_TOKEN = os.Getenv("SYNOPTIC_API_TOKEN")
}
