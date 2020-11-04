package controller

import (
	"net/http"
)

// Authenticate returns an access token
func Authenticate(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	apiError(w, r, http.StatusNotImplemented, "Not implemented yet", 0)
}
