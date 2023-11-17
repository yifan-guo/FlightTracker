package responder

import (
	"FlightTracker/internal/platform/errors"
	"context"
	"log"
	"net/http"
)

type ResponseHandler struct {
	m marshaler
}

func New(m marshaler) *ResponseHandler {
	return &ResponseHandler{
		m: m,
	}
}

type msg struct {
	Details string `json:"details"`
}

type marshaler interface {
	MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
}

func (h *ResponseHandler) Respond(ctx context.Context, w http.ResponseWriter, data interface{}, httpCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	marshaledRes, err := h.m.MarshalIndent(data, "", " ")
	if _, err := w.Write(marshaledRes); err != nil {
		log.Printf("Erorr writing response bod %v", err)
	}

	return err
}

func (h *ResponseHandler) Err(ctx context.Context, w http.ResponseWriter, err error) {
	e, ok := err.(*errors.ErrorResponse)
	if ok {
		// The nolint directives are specified because "h.Respond" logs the error when it occurs
		switch e.Type {
		case errors.ReqBody:
			h.Respond(ctx, w, msg{Details: err.Error()}, http.StatusBadRequest) //nolint:errcheck
		case errors.RespBody:
			h.Respond(ctx, w, msg{Details: err.Error()}, http.StatusInternalServerError) //nolint:errcheck
		default:
			h.Respond(ctx, w, msg{Details: err.Error()}, http.StatusInternalServerError) //nolint:errcheck
		}
	}

}
