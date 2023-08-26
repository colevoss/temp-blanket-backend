package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/colevoss/temperature-blanket-backend/internal/config"
	"github.com/colevoss/temperature-blanket-backend/internal/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

type API struct {
	App *chi.Mux
	Cfg *config.Config
}

func NewApi(cfg *config.Config) *API {
	return &API{
		Cfg: cfg,
		App: chi.NewRouter(),
	}
}

func (a *API) Run() {
	log.Raw().Infow("Starting app", "port", a.Cfg.Port)
	port := fmt.Sprintf(":%s", a.Cfg.Port)
	http.ListenAndServe(port, a.App)
}

func (a *API) Init() {
	a.App.Use(middleware.Logger)
	a.App.Use(requestIdMiddleware)

	a.App.Use(cors.Handler(cors.Options{}))
}

func requestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := uuid.NewString()
		ctx := context.WithValue(r.Context(), "requestId", requestId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
