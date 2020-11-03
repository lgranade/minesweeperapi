// Code generated by sqlc. DO NOT EDIT.
// source: game_query.sql

package minesweeper

import (
	"context"

	"github.com/google/uuid"
)

const createGame = `-- name: CreateGame :one
insert into game (
	id, account_id, row_amount, column_amount, accumulated_seconds, board, mines, mines_left, cells_stepped, game_status
) values (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
returning id, account_id, row_amount, column_amount, accumulated_seconds, board, mines, mines_left, cells_stepped, game_status, created_at
`

type CreateGameParams struct {
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
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, createGame,
		arg.ID,
		arg.AccountID,
		arg.RowAmount,
		arg.ColumnAmount,
		arg.AccumulatedSeconds,
		arg.Board,
		arg.Mines,
		arg.MinesLeft,
		arg.CellsStepped,
		arg.GameStatus,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.RowAmount,
		&i.ColumnAmount,
		&i.AccumulatedSeconds,
		&i.Board,
		&i.Mines,
		&i.MinesLeft,
		&i.CellsStepped,
		&i.GameStatus,
		&i.CreatedAt,
	)
	return i, err
}

const getGameByID = `-- name: GetGameByID :one
select id, account_id, row_amount, column_amount, accumulated_seconds, board, mines, mines_left, cells_stepped, game_status, created_at from game
where id = $1
`

func (q *Queries) GetGameByID(ctx context.Context, id uuid.UUID) (Game, error) {
	row := q.db.QueryRowContext(ctx, getGameByID, id)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.RowAmount,
		&i.ColumnAmount,
		&i.AccumulatedSeconds,
		&i.Board,
		&i.Mines,
		&i.MinesLeft,
		&i.CellsStepped,
		&i.GameStatus,
		&i.CreatedAt,
	)
	return i, err
}

const updateGame = `-- name: UpdateGame :one
update game
set accumulated_seconds = $1, game_status = $2, board = $3, mines_left = $4, cells_stepped = $5
where id = $6
returning id, account_id, row_amount, column_amount, accumulated_seconds, board, mines, mines_left, cells_stepped, game_status, created_at
`

type UpdateGameParams struct {
	AccumulatedSeconds int32     `json:"accumulated_seconds"`
	GameStatus         string    `json:"game_status"`
	Board              string    `json:"board"`
	MinesLeft          int32     `json:"mines_left"`
	CellsStepped       int32     `json:"cells_stepped"`
	ID                 uuid.UUID `json:"id"`
}

func (q *Queries) UpdateGame(ctx context.Context, arg UpdateGameParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, updateGame,
		arg.AccumulatedSeconds,
		arg.GameStatus,
		arg.Board,
		arg.MinesLeft,
		arg.CellsStepped,
		arg.ID,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.RowAmount,
		&i.ColumnAmount,
		&i.AccumulatedSeconds,
		&i.Board,
		&i.Mines,
		&i.MinesLeft,
		&i.CellsStepped,
		&i.GameStatus,
		&i.CreatedAt,
	)
	return i, err
}
