package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/lgranade/minesweeperapi/dao"
	"github.com/lgranade/minesweeperapi/dao/minesweeper"
	"github.com/lgranade/minesweeperapi/model"
)

// ReadGame reads game from db
func ReadGame(ctx context.Context, gameID uuid.UUID) (*model.Game, error) {
	q := dao.GetQuerier().WithoutTx()

	dbGame, err := q.GetGameByID(ctx, gameID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error occurred querying game: ", err)
			return nil, ErrInternal
		}
		return nil, ErrNonexistentGame
	}

	game := &model.Game{}
	fillGameFromDB(game, &dbGame)

	return game, nil
}

func fillGameFromDB(mGame *model.Game, dbGame *minesweeper.Game) error {
	mGame.ID = dbGame.ID
	mGame.UserID = dbGame.AccountID
	mGame.Status = model.GameStatus(dbGame.GameStatus)
	mGame.Mines = int(dbGame.Mines)
	mGame.AccumulatedSeconds = int(dbGame.AccumulatedSeconds)
	mGame.MinesLeft = int(dbGame.MinesLeft)
	mGame.CellsStepped = int(dbGame.CellsStepped)
	mGame.CellAmount = mGame.Rows * mGame.Columns
	mGame.Rows = int(dbGame.RowAmount)
	mGame.Columns = int(dbGame.ColumnAmount)
	mGame.CreatedAt = dbGame.CreatedAt.Unix()

	err := json.Unmarshal([]byte(dbGame.Board), &mGame.Board)
	if err != nil {
		log.Println("Can't load board from db")
		return ErrInternal
	}
	return nil
}
