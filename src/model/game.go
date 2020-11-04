package model

import (
	"encoding/json"
	"log"
	"time"

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
	AccumulatedSeconds int        `json:"-"`
	PlayingSeconds     int        `json:"playingSeconds"`
	MinesLeft          int        `json:"minesLeft"`
	CreatedAt          time.Time  `json:"createdAt"`
	Board              [][]Cell   `json:"board,omitempty"`
	CellsStepped       int        `json:"-"`
	CellAmount         int        `json:"-"`
	ResumedAt          time.Time  `json:"-"`
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

// Coord represents a coordinate in the board, this is used by some algos to simplify code
type Coord struct {
	Row int
	Col int
}

// BoardString serializes board to be stored in db
func (g *Game) BoardString() string {
	boardRaw, err := json.Marshal(&g.Board)
	if err != nil {
		log.Println("Bug: An error occurred trying to store in db, err: ", err)
	}
	return string(boardRaw)
}

// Get returns reference to cell by coordinates
func (g *Game) Get(c Coord) *Cell {
	return &g.Board[c.Row][c.Col]
}

// Surrounding returns all surrounding cells
func (g *Game) Surrounding(c Coord) []Coord {
	surrounding := []Coord{}

	if sc, ok := g.Upper(c); ok {
		surrounding = append(surrounding, sc)
	}
	if sc, ok := g.UpperRight(c); ok {
		surrounding = append(surrounding, sc)
	}
	if sc, ok := g.Right(c); ok {
		surrounding = append(surrounding, sc)
	}
	if sc, ok := g.LowerRight(c); ok {
		surrounding = append(surrounding, sc)
	}
	if sc, ok := g.Lower(c); ok {
		surrounding = append(surrounding, sc)
	}
	if sc, ok := g.LowerLeft(c); ok {
		surrounding = append(surrounding, sc)
	}
	if sc, ok := g.Left(c); ok {
		surrounding = append(surrounding, sc)
	}
	if sc, ok := g.UpperLeft(c); ok {
		surrounding = append(surrounding, sc)
	}

	return surrounding
}

// Upper gives you the upper cell relative to this one
func (g *Game) Upper(c Coord) (Coord, bool) {
	if c.Row == 0 {
		return Coord{}, false
	}
	return Coord{Row: c.Row - 1, Col: c.Col}, true
}

// UpperRight gives you the upper right cell relative to this one
func (g *Game) UpperRight(c Coord) (Coord, bool) {
	if c.Row == 0 {
		return Coord{}, false
	}
	if c.Col == g.Columns-1 {
		return Coord{}, false
	}
	return Coord{Row: c.Row - 1, Col: c.Col + 1}, true
}

// Right gives you the right cell relative to this one
func (g *Game) Right(c Coord) (Coord, bool) {
	if c.Col == g.Columns-1 {
		return Coord{}, false
	}
	return Coord{Row: c.Row, Col: c.Col + 1}, true
}

// LowerRight gives you the lower right cell relative to this one
func (g *Game) LowerRight(c Coord) (Coord, bool) {
	if c.Row == g.Rows-1 {
		return Coord{}, false
	}
	if c.Col == g.Columns-1 {
		return Coord{}, false
	}
	return Coord{Row: c.Row + 1, Col: c.Col + 1}, true
}

// Lower gives you the lower cell relative to this one
func (g *Game) Lower(c Coord) (Coord, bool) {
	if c.Row == g.Rows-1 {
		return Coord{}, false
	}
	return Coord{Row: c.Row + 1, Col: c.Col}, true
}

// LowerLeft gives you the lower left cell relative to this one
func (g *Game) LowerLeft(c Coord) (Coord, bool) {
	if c.Row == g.Rows-1 {
		return Coord{}, false
	}
	if c.Col == 0 {
		return Coord{}, false
	}
	return Coord{Row: c.Row + 1, Col: c.Col - 1}, true
}

// Left gives you the left cell relative to this one
func (g *Game) Left(c Coord) (Coord, bool) {
	if c.Col == 0 {
		return Coord{}, false
	}
	return Coord{Row: c.Row, Col: c.Col - 1}, true
}

// UpperLeft gives you the upper left cell relative to this one
func (g *Game) UpperLeft(c Coord) (Coord, bool) {
	if c.Row == 0 {
		return Coord{}, false
	}
	if c.Col == 0 {
		return Coord{}, false
	}
	return Coord{Row: c.Row - 1, Col: c.Col - 1}, true
}
