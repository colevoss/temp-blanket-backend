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

	summary, err := tsh.Repos.Weather.GetSummary(ctx, parsedDate)

	if err != nil {
		response.Error(w, r, err)
		return
	}

	response.Ok(w, r, summary)
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
	yearAgo := time.Now().AddDate(-1, 0, 0)

	if date.After(now) || date.Before(yearAgo) {
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
