package controller

import "net/http"

type playReq struct {
	Row    int    `json:"row,omitempty"`
	Column int    `json:"column,omitempty"`
	Action string `json:"action,omitempty"`
}

// Play creates a new move on the game board
func Play(w http.ResponseWriter, r *http.Request) {
	// TODO
}
