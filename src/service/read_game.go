package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lgranade/minesweeperapi/dao"
	"github.com/lgranade/minesweeperapi/dao/minesweeper"
	"github.com/lgranade/minesweeperapi/model"
)

// ReadGame reads game from db
func ReadGame(ctx context.Context, userID uuid.UUID, gameID uuid.UUID) (*model.Game, error) {
	q := dao.GetQuerier().WithoutTx()

	_, err := q.GetAccountByID(ctx, userID)
	if err == sql.ErrNoRows {
		log.Println("Non existent user in db by id: ", userID)
		return nil, ErrNonexistentUser
	}
	if err != nil {
		log.Println("Error reading user from local db: ", err)
		return nil, ErrInternal
	}

	dbGame, err := q.GetGameByID(ctx, gameID)
	if err == sql.ErrNoRows {
		log.Println("Non existent game in db by id: ", gameID)
		return nil, ErrNonexistentGame
	}
	if err != nil {
		log.Println("Error reading game from local db: ", err)
		return nil, ErrInternal
	}

	game := &model.Game{}
	fillGameFromDB(game, &dbGame)

	if game.UserID != userID {
		return nil, ErrForbidden
	}

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
	mGame.ResumedAt = dbGame.ResumedAt.Unix()

	err := json.Unmarshal([]byte(dbGame.Board), &mGame.Board)
	if err != nil {
		log.Println("Can't load board from db")
		return ErrInternal
	}

	mGame.PlayingSeconds = mGame.AccumulatedSeconds + int(time.Now().Unix()-mGame.ResumedAt)
	return nil
}
