package server

import (
	"FlightTracker/cmd/server/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Routes(h *handlers.RootHandler) chi.Router {
	r := chi.NewRouter()

	r.MethodFunc(http.MethodPost, "/calculate", h.FlightCalculate)

	return r
}
