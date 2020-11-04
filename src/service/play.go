package service

import (
	"context"
	"database/sql"
	"log"
	"time"

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

	if game.Status != model.GameCreated &&
		game.Status != model.GamePlaying &&
		game.Status != model.GamePaused {
		return nil, ErrGameNonCompatibleStatus
	}
	if game.Status == model.GameCreated ||
		game.Status == model.GamePaused {
		game.ResumedAt = time.Now().UTC()
		game.Status = model.GamePlaying
	}

	if game.UserID != userID {
		log.Println("Game belongs to different user")
		return nil, ErrForbidden
	}

	err = calculatePlay(game, model.Coord{Row: row, Col: column}, action)
	if err != nil {
		return nil, err
	}

	_, err = q.UpdateGame(ctx, minesweeper.UpdateGameParams{
		AccumulatedSeconds: int32(game.AccumulatedSeconds),
		GameStatus:         string(game.Status),
		Board:              game.BoardString(),
		MinesLeft:          int32(game.MinesLeft),
		CellsStepped:       int32(game.CellsStepped),
		ResumedAt:          game.ResumedAt,
		ID:                 game.ID,
	})

	if err = tx.Commit(); err != nil {
		log.Println("Error occurred trying to commit tx: ", err)
		return nil, ErrInternal
	}

	return game, nil
}

func calculatePlay(game *model.Game, coord model.Coord, action PlayAction) error {
	if coord.Row >= game.Rows || coord.Col >= game.Columns {
		return ErrOutsideBoardBoundaries
	}

	cell := game.Get(coord)

	if action == FlagPlay {
		// flag acts as a toggle
		if cell.Action == model.NoAction {
			cell.Action = model.FlagAction
			game.MinesLeft--
		} else if cell.Action == model.FlagAction {
			cell.Action = model.NoAction
			game.MinesLeft++
		}
	} else if action == StepPlay {
		stepPlay(game, coord)
	} else {
		return ErrUnrecognizedAction
	}

	if game.CellAmount-game.CellsStepped == game.Mines {
		game.Status = model.GameWon
		game.AccumulatedSeconds += int(time.Now().UTC().Sub(game.ResumedAt).Seconds())
	}

	return nil
}

func stepPlay(game *model.Game, coord model.Coord) {
	cell := game.Get(coord)
	if cell.Action == model.NoAction || cell.Action == model.FlagAction {
		cell.Action = model.StepAction
		game.CellsStepped++
		if cell.Type == model.MineCell {
			game.Status = model.GameLost
		} else if cell.Value == 0 {
			for _, sc := range game.Surrounding(coord) {
				sCell := game.Get(sc)
				if sCell.Action == model.NoAction {
					stepPlay(game, sc)
				}
			}
		}
	}
}
