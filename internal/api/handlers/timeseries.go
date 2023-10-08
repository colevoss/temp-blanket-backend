package handlers

import (
	"net/http"
	"time"

	"github.com/colevoss/temperature-blanket-backend/internal/log"
	"github.com/colevoss/temperature-blanket-backend/internal/repositories"
	"github.com/colevoss/temperature-blanket-backend/internal/response"
)

var DATE_FORMAT = "20060102"

type TimeSeriesHandlers struct {
	Repos *repositories.Repositories
}

func (tsh *TimeSeriesHandlers) GetSummary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dateQuery := r.URL.Query().Get("date")
	stationId := r.URL.Query().Get("station")

	if stationId == "" {
		response.Error(w, r, response.BadRequest("station required", nil, nil))
		return
	}

	station, err := tsh.Repos.Weather.GetStation(ctx, stationId)

	if err != nil {
		response.Error(w, r, response.NotFound("Station not found", nil, err))
		return
	}

	parsedDate := tsh.parseIsoDateQuery(dateQuery)

	log.C(ctx).Infow(
		"Summary date request",
		"dateQuery", dateQuery,
		"date", parsedDate,
	)

	if err := tsh.validateDate(parsedDate); err != nil {
		response.Error(w, r, err)
		return
	}

	data, err := tsh.Repos.Weather.GetWeatherData(ctx, parsedDate, stationId)

	if err != nil {
		response.Error(w, r, err)
		return
	}

	// summary, err := tsh.Repos.Weather.GetSummary(ctx, parsedDate, stationId)
	summary, err := tsh.Repos.Weather.Summary(data, parsedDate)

	if err != nil {
		response.Error(w, r, err)
		return
	}

	// response.Ok(w, r, summary)
	response.Ok(w, r, map[string]interface{}{
		"summary": summary,
		"station": station,
		"data":    data,
	})
}

func (tsh *TimeSeriesHandlers) parseIsoDateQuery(dateQuery string) time.Time {
	parsedDate, err := time.Parse(time.RFC3339, dateQuery)

	if err != nil {
		parsedDate = time.Now()
	}

	return parsedDate
}

func (tsh *TimeSeriesHandlers) validateDate(date time.Time) error {
	now := time.Now()
	// yearAgo := time.Now().AddDate(-1, 0, 0)

	// if date.After(now) || date.Before(yearAgo) {
	if date.After(now) {
		return response.BadRequest("Invalid date", response.Map{
			"date": date,
		}, nil)
	}

	return nil
}

func NewTimeSeriesHandlers(repos *repositories.Repositories) *TimeSeriesHandlers {
	return &TimeSeriesHandlers{
		Repos: repos,
	}
}
