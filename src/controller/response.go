package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AlekSi/pointer"
)

func apiResponse(w http.ResponseWriter, r *http.Request, code int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if resp != nil {
		b, err := json.Marshal(resp)
		if err == nil {
			w.Write(b)
		}
	}
	return
}

func apiError(w http.ResponseWriter, r *http.Request, code int, msg string, internalCode int) {
	resp := errorResp{
		Error: &errorRespError{
			Message: msg,
		},
	}
	if internalCode != 0 {
		resp.Error.Code = pointer.ToInt(internalCode)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	b, err := json.Marshal(resp)
	if err == nil {
		w.Write(b)
	}
	return
}

type errorResp struct {
	Error *errorRespError `json:"error,omitempty"`
}

type errorRespError struct {
	Code    *int   `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// HTTP 400
const (
	IErrorIllFormedRequest = iota + 20100
	IErrorLackingFields
)

// HTTP 401
const (
	IErrorNotAuthorized = iota + 20200
)

// HTTP 404
const (
	IErrorNonexistentGame = iota + 20500
)
