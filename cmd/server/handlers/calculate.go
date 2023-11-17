package handlers

import (
	"FlightTracker/cmd/server/api"
	"FlightTracker/cmd/server/handlers/responses"
	"FlightTracker/internal/models"
	"fmt"
	"log"
	"net/http"
)

type CalculateHander struct {
	WebResponder WebResponder
}

func (h RootHandler) FlightCalculate(w http.ResponseWriter, r *http.Request) {

	var flights []models.Flight

	if err := responses.RequestToModel(r, &flights); err != nil {
		h.WebResponder.Err(r.Context(), w, err)
	}

	log.Printf("log level '%s' - PostTrack: request body decoded successfully", traceLevel)

	var FlightPayload [][]string = make([][]string, len(flights))
	for i, flight := range flights {
		FlightPayload[i] = []string{flight.Start, flight.End}
	}

	var itinerary []string
	var start, finish string
	max := -1

	for key := range api.CreateGraph(FlightPayload) {
		itinerary = api.FindItinerary(FlightPayload, key)
		fmt.Println(itinerary)
		if len(itinerary) > max {
			max = len(itinerary)
			start = itinerary[0]
			finish = itinerary[len(itinerary)-1]
		}
	}

	tracked := []string{start, finish}
	h.WebResponder.Respond(r.Context(), w, tracked, http.StatusOK)

	log.Printf("log level '%s' - PostTrack: response body '%v' written with status code '%d'",
		traceLevel, tracked, http.StatusOK)
}
