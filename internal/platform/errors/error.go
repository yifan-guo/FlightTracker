package errors

import "fmt"

const (
	ReqBody             = "ReqBody"
	reqBodyMessage      = "unable to read request body"
	RespBody            = "RespBody"
	respBodyMessage     = "unable to write response body"
	unknownErrorMessage = "something went wrong"
)

type ErrorResponse struct {
	e    error
	Type string
}

func (r *ErrorResponse) Error() string {
	return r.e.Error()
}

func New(failType string, message string) *ErrorResponse {
	var err error
	switch failType {
	case ReqBody:
		err = fmt.Errorf("%s: %s", reqBodyMessage, message)
	case RespBody:
		err = fmt.Errorf("%s: %s", respBodyMessage, message)
	default:
		err = fmt.Errorf("%s: %s", unknownErrorMessage, message)
	}

	return &ErrorResponse{
		e:    err,
		Type: failType,
	}
}
