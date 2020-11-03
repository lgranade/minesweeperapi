package service

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/lgranade/minesweeperapi/dao"
	"github.com/lgranade/minesweeperapi/dao/minesweeper"
	"github.com/lgranade/minesweeperapi/model"
)

// PlayAction is the move the user does on a cell
type PlayAction string

const (
	// StepPlay indicates that the user wants to step on the cell
	StepPlay PlayAction = "step"
	// FlagPlay indicates that the user wants to flag the cell
	FlagPlay PlayAction = "flag"
)

// Play executes the action the user decided to make on the board and persists it
func Play(ctx context.Context, userID uuid.UUID, gameID uuid.UUID, row int, column int, action PlayAction) (*model.Game, error) {
	q, tx, err := dao.GetQuerier().WithTx()
	if err != nil {
		log.Println("An error occurred trying to establish db connection")
		return nil, ErrInternal
	}

	defer tx.Rollback()

	dbGame, err := q.GetAndLockGameByID(ctx, gameID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error occurred querying game: ", err)
			return nil, ErrInternal
		}
		return nil, ErrNonexistentGame
	}

	game := &model.Game{}
	fillGameFromDB(game, &dbGame)

	if game.UserID != userID {
		log.Println("Game belongs to differen user")
		return nil, ErrForbidden
	}

	calculatePlay(game, row, column, action)

	_, err = q.UpdateGame(ctx, minesweeper.UpdateGameParams{
		AccumulatedSeconds: int32(game.AccumulatedSeconds),
		GameStatus:         string(game.Status),
		Board:              game.BoardString(),
		MinesLeft:          int32(game.MinesLeft),
		CellsStepped:       int32(game.CellsStepped),
		ID:                 game.ID,
	})

	if err = tx.Commit(); err != nil {
		log.Println("Error occurred trying to commit tx: ", err)
		return nil, ErrInternal
	}

	return game, nil
}

func calculatePlay(game *model.Game, row int, column int, action PlayAction) error {
	if row >= game.Rows || column >= game.Columns {
		return ErrOutsideBoardBoundaries
	}

	cell := &game.Board[row][column]

	if action == FlagPlay {
		// flag acts as a toggle
		if cell.Action == model.NoAction {
			cell.Action = model.FlagAction
		} else if cell.Action == model.FlagAction {
			cell.Action = model.NoAction
		}
	} else if action == StepPlay {
		stepPlay(cell)
	}
	return nil
}

func stepPlay(cell *model.Cell) {
	if cell.Action == model.NoAction || cell.Action == model.FlagAction {
		cell.Action = model.StepAction
		cell.Game.CellsStepped++
		if cell.Type == model.MineCell {
			cell.Game.Status = model.GameLost
		} else if cell.Value == 0 {
			for _, sc := range cell.Surrounding() {
				if sc.Action == model.NoAction {
					stepPlay(sc)
				}
			}
		}
	}
}
