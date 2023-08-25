package handlers

import (
	"net/http"

	"github.com/colevoss/temperature-blanket-backend/internal/logger"
	"github.com/gin-gonic/gin"
)

type PingHandlers struct {
}

func (ph *PingHandlers) Ping(c *gin.Context) {
	logger.Req(c).Infow("PING")
	c.JSON(http.StatusOK, gin.H{
		"pong": true,
	})
}

func NewPingHandlers() *PingHandlers {
	handlers := &PingHandlers{}

	return handlers
}
