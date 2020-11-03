package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/lgranade/minesweeperapi/service"
)

type createUserReq struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	reqB := createUserReq{}
	err := decoder.Decode(&reqB)
	if err != nil {
		apiError(w, r, http.StatusBadRequest, "Couldn't parse body", IErrorIllFormedRequest)
		return
	}

	if reqB.Name == "" || reqB.Password == "" {
		apiError(w, r, http.StatusBadRequest, "Missing required parameters. See documentation", IErrorLackingFields)
		return
	}

	user, err := service.CreateUser(r.Context(), reqB.Name, reqB.Password)
	if err != nil {
		apiError(w, r, http.StatusInternalServerError, "Couldn't create the user, report error", 0)
		return
	}

	apiResponse(w, r, http.StatusCreated, user)
}

// ReadUser reads an existing user
func ReadUser(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(chi.URLParam(r, "userID"))
	if err != nil {
		apiError(w, r, http.StatusBadRequest, "Error parsing parameter, read documentation", IErrorIllFormedRequest)
		return
	}

	user, err := service.ReadUser(r.Context(), userID)
	if err != nil {
		// TODO: check for more errors
		apiError(w, r, http.StatusInternalServerError, "Couldn't read user, report error", 0)
		return
	}

	apiResponse(w, r, http.StatusOK, user)
}
