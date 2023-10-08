package handlers

import (
	"net/http"

	"github.com/colevoss/temperature-blanket-backend/internal/log"
	"github.com/colevoss/temperature-blanket-backend/internal/repositories"
	"github.com/colevoss/temperature-blanket-backend/internal/response"
	"github.com/go-chi/chi/v5"
)

type NetworkHandlers struct {
	Repos *repositories.Repositories
}

func (h *NetworkHandlers) GetNetworks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.C(ctx).Infow("Fetching networks")

	networks, err := h.Repos.Weather.GetNetworks(ctx)

	if err != nil {
		response.Error(w, r, err)
		return
	}

	log.C(ctx).Infow("Networks fetched successfully")

	response.Ok(w, r, networks)
}

func (h *NetworkHandlers) GetStations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	networkId := chi.URLParam(r, "networkId")

	log.C(ctx).Infow("Fetching network stations", "networkId", networkId)

	stations, err := h.Repos.Weather.GetStations(ctx, networkId)

	if err != nil {
		response.Error(w, r, err)
		return
	}

	log.C(ctx).Infow("Networks fetched successfully", "networkId", networkId)

	response.Ok(w, r, stations)
}

func NewNetworkHandler(repos *repositories.Repositories) *NetworkHandlers {
	return &NetworkHandlers{repos}
}
