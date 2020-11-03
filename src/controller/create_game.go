package controller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/lgranade/minesweeperapi/service"
)

// TODO: take this from access token
var hardcodedUserID uuid.UUID

func init() {
	hardcodedUserID, _ = uuid.Parse("e341410d-752a-404f-9acc-904764fd38f3")
}

type createGameReq struct {
	Rows    int `json:"rows,omitempty"`
	Columns int `json:"columns,omitempty"`
	Mines   int `json:"mines,omitempty"`
}

// CreateGame creates a new game session for the user
func CreateGame(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	reqB := createGameReq{}
	err := decoder.Decode(&reqB)
	if err != nil {
		apiError(w, r, http.StatusBadRequest, "Couldn't parse body", IErrorIllFormedRequest)
		return
	}

	if reqB.Rows == 0 || reqB.Columns == 0 || reqB.Mines == 0 {
		apiError(w, r, http.StatusBadRequest, "Missing required parameters. See documentation", IErrorLackingFields)
		return
	}

	game, err := service.CreateGame(r.Context(), hardcodedUserID, reqB.Rows, reqB.Columns, reqB.Mines)
	if err != nil {
		apiError(w, r, http.StatusInternalServerError, "Couldn't create the game session, report error", 0)
		return
	}

	apiResponse(w, r, http.StatusCreated, game)
}
