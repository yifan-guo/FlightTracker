package responses

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ReqVars struct {
	Context context.Context
}

func GetRequestVars(r *http.Request) ReqVars {
	return ReqVars{
		Context: r.Context(),
	}
}

func RequestToModel(r *http.Request, outputStruct interface{}) error {
	if r.Body == nil {
		log.Printf("nil body, coud not decode to struct")
		return http.ErrBodyNotAllowed
	}

	requestBody, _ := io.ReadAll(r.Body)
	decoder := json.NewDecoder(bytes.NewReader(requestBody))

	if err := decoder.Decode(&outputStruct); err != nil {
		log.Printf("could not decode %v to struct, error: %v", requestBody, err)
		return http.ErrBodyNotAllowed
	}

	return nil
}
