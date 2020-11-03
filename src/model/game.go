package model

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

// Game represents the public model for a game session
type Game struct {
	ID                 uuid.UUID  `json:"id"`
	UserID             uuid.UUID  `json:"userId"`
	Status             GameStatus `json:"status"`
	Rows               int        `json:"rows"`
	Columns            int        `json:"columns"`
	Mines              int        `json:"mines"`
	AccumulatedSeconds int        `json:"accumulatedSeconds"`
	MinesLeft          int        `json:"minesLeft"`
	CreatedAt          int64      `json:"createdAt"`
	Board              [][]Cell   `json:"board,omitempty"`
	CellsStepped       int        `json:"-"`
	CellAmount         int        `json:"-"`
	ResumedAt          int64      `json:"-"`
}

// GameStatus represents the status of the game
type GameStatus string

const (
	// GameCreated indicates that the game has just been created and no play has been executed
	GameCreated GameStatus = "created"

	// GamePlaying indicates that the user is currently playing and the clock is ticking
	GamePlaying GameStatus = "playing"

	// GamePaused indicates that the game is paused, clock is not ticking
	GamePaused GameStatus = "paused"

	// GameLost indicates that the user has lost this game (step on a mine)
	GameLost GameStatus = "lost"

	// GameWon indicates that the user has won this game
	GameWon GameStatus = "won"
)

// BoardString serializes board to be stored in db
func (g *Game) BoardString() string {
	boardRaw, err := json.Marshal(&g.Board)
	if err != nil {
		log.Println("Bug: An error occurred trying to store in db, err: ", err)
	}
	return string(boardRaw)
}
