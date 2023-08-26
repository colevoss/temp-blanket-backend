package handlers

import (
	"net/http"

	"github.com/colevoss/temperature-blanket-backend/internal/log"
	"github.com/colevoss/temperature-blanket-backend/internal/response"
)

type PingHandlers struct{}

func (ph *PingHandlers) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.C(ctx).Infow("hello", "ping", "pong")

	response.Ok(w, r, response.Map{
		"ping": "pong",
	})
}

func NewPingHandlers() *PingHandlers {
	handlers := &PingHandlers{}

	return handlers
}
