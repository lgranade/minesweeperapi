package controller

import "net/http"

type createGameReq struct {
	Rows    int `json:"rows,omitempty"`
	Columns int `json:"columns,omitempty"`
	Mines   int `json:"mines,omitempty"`
}

// CreateGame creates a new game session for the user
func CreateGame(w http.ResponseWriter, r *http.Request) {
	// TODO
}
