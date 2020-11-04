package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/lgranade/minesweeperapi/service"
)

type playReq struct {
	Row    *int               `json:"row,omitempty"`
	Column *int               `json:"column,omitempty"`
	Action service.PlayAction `json:"action,omitempty"`
}

// Play creates a new move on the game board
func Play(w http.ResponseWriter, r *http.Request) {
	gameID, err := uuid.Parse(chi.URLParam(r, "gameID"))
	if err != nil {
		apiError(w, r, http.StatusBadRequest, "Error parsing parameter, read documentation", IErrorIllFormedRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	reqB := playReq{}
	err = decoder.Decode(&reqB)
	if err != nil {
		apiError(w, r, http.StatusBadRequest, "Couldn't parse body", IErrorIllFormedRequest)
		return
	}

	if reqB.Row == nil || reqB.Column == nil || reqB.Action == "" {
		apiError(w, r, http.StatusBadRequest, "Missing required parameters. See documentation", IErrorLackingFields)
		return
	}

	game, err := service.Play(r.Context(), hardcodedUserID, gameID, *reqB.Row, *reqB.Column, reqB.Action)
	if err != nil {
		if errors.Is(err, service.ErrForbidden) {
			apiError(w, r, http.StatusForbidden, "Operation nor allowed", IErrorForbidden)
			return
		} else if errors.Is(err, service.ErrNonexistentGame) {
			apiError(w, r, http.StatusNotFound, "Nonexistent Game", IErrorNonexistentGame)
			return
		}
		apiError(w, r, http.StatusInternalServerError, "Couldn't make the play, report error", 0)
		return
	}

	apiResponse(w, r, http.StatusOK, game)
}
