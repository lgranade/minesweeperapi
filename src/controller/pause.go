package controller

import "net/http"

// Pause pauses the game so the time counter stops running
func Pause(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	apiError(w, r, http.StatusNotImplemented, "Not implemented yet", 0)
}
