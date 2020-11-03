package model

import "github.com/google/uuid"

// Game represents the public model for a game session
type Game struct {
	ID                 uuid.UUID  `json:"id,omitempty"`
	UserID             uuid.UUID  `json:"userId,omitempty"`
	Status             GameStatus `json:"status,omitempty"`
	Rows               int        `json:"rows,omitempty"`
	Columns            int        `json:"columns,omitempty"`
	Mines              int        `json:"mines,omitempty"`
	AccumulatedSeconds int        `json:"accumulatedSeconds,omitempty"`
	MinesLeft          int        `json:"minesLeft,omitempty"`
	Board              [][]Cell   `json:"board,omitempty"`
	CellAmount         int        `json:"-"`
	CreatedAt          int64      `json:"createdAt,omitempty"`
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
