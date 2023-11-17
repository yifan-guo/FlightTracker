package main

import (
	"FlightTracker/cmd/server"
	"FlightTracker/cmd/server/handlers"
	"FlightTracker/internal/marshaler"
	"FlightTracker/internal/responder"
	"log"
	"net/http"
	"time"
)

type flight struct {
	Start string
	End   string
}

var oneFlight = []flight{
	{
		Start: "ATL",
		End:   "EWR",
	},
}

// start: BGY; end: AKL
var flights = []flight{
	{
		Start: "BCN",
		End:   "PSC",
	},
	{
		Start: "JFK",
		End:   "AAL",
	},
	{
		Start: "FCO",
		End:   "BCN",
	},
	{
		Start: "GSO",
		End:   "IND",
	},
	{
		Start: "SFO",
		End:   "ATL",
	},
	{
		Start: "AAL",
		End:   "HEL",
	},
	{
		Start: "PSC",
		End:   "BLQ",
	},
	{
		Start: "IND",
		End:   "EWR",
	},
	{
		Start: "BGY",
		End:   "RAR",
	},
	{
		Start: "BJZ",
		End:   "AKL",
	},
	{
		Start: "AUH",
		End:   "FCO",
	},
	{
		Start: "HEL",
		End:   "CAK",
	},
	{
		Start: "RAR",
		End:   "AUH",
	},
	{
		Start: "CAK",
		End:   "BJZ",
	},
	{
		Start: "ATL",
		End:   "GSO",
	},
	{
		Start: "CHI",
		End:   "JFK",
	},
	{
		Start: "BLQ",
		End:   "MAD",
	},
	{
		Start: "EWR",
		End:   "CHI",
	},
	{
		Start: "MAD",
		End:   "SFO",
	},
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("FATAL - could not start server: %v", err)
	}
}

func run() error {

	port := "8080"

	respHandler := responder.New(marshaler.New())
	rootHandler :=
		handlers.New(
			respHandler,
		)
	log.Printf("Main: starting service on port :%s", port)
	service := &http.Server{
		Addr:              ":" + port,
		Handler:           server.Routes(rootHandler),
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := service.ListenAndServe(); err != nil {
		return err
	}
	return nil

}
