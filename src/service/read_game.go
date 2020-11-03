package service

import (
	"encoding/json"
	"log"

	"github.com/lgranade/minesweeperapi/dao/minesweeper"
	"github.com/lgranade/minesweeperapi/model"
)

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
