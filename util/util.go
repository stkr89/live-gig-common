package util

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/stkr89/live-gig-common/types"
	"net/http"
)

func err2code(err error) int {
	switch err {
	case types.InvalidRequestBody:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	_ = json.NewEncoder(w).Encode(types.ErrorWrapper{Error: err.Error()})
}

func EncodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
