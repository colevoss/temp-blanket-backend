package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/colevoss/temperature-blanket-backend/internal/integrations/synoptic"
	"github.com/colevoss/temperature-blanket-backend/internal/repositories/weather"
	"github.com/gin-gonic/gin"
)

var DATE_FORMAT = "20060102"

type TimeSeriesHandlers struct {
	WeatherRepo weather.WeatherRepository
}

type SummaryInvalidDateError struct {
	Date time.Time
}

func (side *SummaryInvalidDateError) Error() string {
	return "Invalid date"
}

func (tsh *TimeSeriesHandlers) GetTimeSeries(c *gin.Context) {
	dateQuery := c.Query("date")
	// parsedDate := tsh.parseDateQuery(dateQuery)
	parsedDate := tsh.parseIsoDateQuery(dateQuery)

	log.Printf("Requesting Timeseries Data for %v", parsedDate)

	synopticApi := synoptic.NewSynopticApi()
	ts, _ := synopticApi.GetTimeSeriesTemperatureData(parsedDate)

	c.JSON(http.StatusOK, ts)
}

func (tsh *TimeSeriesHandlers) GetSummary(c *gin.Context) {
	dateQuery := c.Query("date")
	// parsedDate := tsh.parseDateQuery(dateQuery)
	parsedDate := tsh.parseIsoDateQuery(dateQuery)

	if err := tsh.validateDate(parsedDate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	summary, err := tsh.WeatherRepo.GetSummary(parsedDate)

	if err != nil {
		var synErr *synoptic.SynopticApiError

		switch {
		case errors.As(err, &synErr):
			c.JSON(synErr.Summary.ResponseCode, gin.H{
				"error": synErr.Summary.ResponseMessage,
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	c.JSON(http.StatusOK, summary)
}

func (tsh *TimeSeriesHandlers) parseIsoDateQuery(dateQuery string) time.Time {
	log.Printf("DATE Query: %s", dateQuery)
	parsedDate, err := time.Parse(time.RFC3339, dateQuery)

	if err != nil {
		parsedDate = time.Now()
	}

	return parsedDate
}

func (tsh *TimeSeriesHandlers) parseDateQuery(dateQuery string) time.Time {
	parsedDate, err := time.Parse(DATE_FORMAT, dateQuery)

	if err != nil {
		parsedDate = time.Now()
	}

	return parsedDate
}

func (tsh *TimeSeriesHandlers) validateDate(date time.Time) error {
	now := time.Now()
	yearAgo := time.Now().AddDate(-1, 0, 0)

	if date.After(now) || date.Before(yearAgo) {
		return &SummaryInvalidDateError{Date: date}
	}

	return nil
}

func NewTimeSeriesHandlers(weatherRepo weather.WeatherRepository) *TimeSeriesHandlers {
	return &TimeSeriesHandlers{
		WeatherRepo: weatherRepo,
	}
}
