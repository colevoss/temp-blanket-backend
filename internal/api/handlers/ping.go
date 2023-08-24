package handlers

import (
	"net/http"
	"time"

	"github.com/colevoss/temperature-blanket-backend/internal/integrations/synoptic"
	"github.com/gin-gonic/gin"
)

type PingHandlers struct{}

func (ph *PingHandlers) Ping(c *gin.Context) {
	s := synoptic.NewSynopticApi()

	ts, _ := s.GetTimeSeriesTemperatureData(time.Now())

	c.JSON(http.StatusOK, ts)
}

func NewPingHandlers() *PingHandlers {
	handlers := &PingHandlers{}

	return handlers
}
