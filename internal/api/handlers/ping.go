package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandlers struct{}

func (ph *PingHandlers) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pong": true,
	})
}

func NewPingHandlers() *PingHandlers {
	handlers := &PingHandlers{}

	return handlers
}
