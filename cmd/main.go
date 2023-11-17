package main

import (
	_ "FlightTracker/cmd/docs"
	"FlightTracker/cmd/server"
	"FlightTracker/cmd/server/handlers"
	"FlightTracker/cmd/server/handlers/routes"
	"FlightTracker/internal/marshaler"
	"FlightTracker/internal/responder"
	"github.com/labstack/echo/v4"
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

// @title Flight Path API
// @version 1.0
// @description This is REST API server to determine the flight.go path of a person.
// @termsOfService http://swagger.io/terms/

// @contact.name Yifan Guo
// @contact.url https://github.com/yifan-guo/FlightTracker
// @contact.email yifanguo247@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
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

	e := echo.New()
	routes.SwaggerRoutes(e)
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- e.Start(":1323")
	}()

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
