// Code generated by sqlc. DO NOT EDIT.

package minesweeper

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID            uuid.UUID `json:"id"`
	LoginName     string    `json:"login_name"`
	LoginPassword string    `json:"login_password"`
	CreatedAt     time.Time `json:"created_at"`
}

type Game struct {
	ID                 uuid.UUID `json:"id"`
	AccountID          uuid.UUID `json:"account_id"`
	RowAmount          int32     `json:"row_amount"`
	ColumnAmount       int32     `json:"column_amount"`
	AccumulatedSeconds int32     `json:"accumulated_seconds"`
	Board              string    `json:"board"`
	Mines              int32     `json:"mines"`
	MinesLeft          int32     `json:"mines_left"`
	CellsStepped       int32     `json:"cells_stepped"`
	GameStatus         string    `json:"game_status"`
	CreatedAt          time.Time `json:"created_at"`
	ResumedAt          time.Time `json:"resumed_at"`
}
