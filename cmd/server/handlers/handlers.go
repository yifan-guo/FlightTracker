package handlers

import (
	"context"
	"log"
	"net/http"
)

const traceLevel = "trace"

type WebResponder interface {
	Respond(ctx context.Context, w http.ResponseWriter, data interface{}, httpCode int) error
	Err(ctx context.Context, w http.ResponseWriter, err error)
}

type RootHandler struct {
	WebResponder WebResponder
	Handler      func(http.ResponseWriter, *http.Request) error
}

func (rh *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Context(), "FlightTracker: Request received")
	err := rh.Handler(w, r)
	if err == nil {
		log.Print("FlightTracker: Call Successful", traceLevel)
		return
	}
	log.Printf("Error from service handler:%+v", err)
}

func New(w WebResponder) *RootHandler {
	return &RootHandler{
		WebResponder: w,
	}
}
