package service

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"

	"github.com/lgranade/minesweeperapi/dao"
	"github.com/lgranade/minesweeperapi/dao/minesweeper"
	"github.com/lgranade/minesweeperapi/model"
)

// CreateGame creates a new game
func CreateGame(ctx context.Context, userID uuid.UUID, rows int, columns int, mines int) (*model.Game, error) {
	q, tx, err := dao.GetQuerier().WithTx()
	if err != nil {
		log.Println("An error occurred trying to establish db connection")
		return nil, ErrInternal
	}

	defer tx.Rollback()

	// TODO: check the user actually exists in db, if not return forbidden

	game := buildGame(userID, rows, columns, mines)

	boardRaw, err := json.Marshal(&game.Board)
	if err != nil {
		log.Println("An error occurred trying to store in db, err: ", err)
		return nil, ErrInternal
	}
	boardStr := string(boardRaw)

	_, err = q.CreateGame(ctx, minesweeper.CreateGameParams{
		ID:                 game.ID,
		AccountID:          game.UserID,
		RowAmount:          int32(game.Rows),
		ColumnAmount:       int32(game.Columns),
		AccumulatedSeconds: int32(game.AccumulatedSeconds),
		Board:              boardStr,
		Mines:              int32(mines),
		MinesLeft:          int32(mines),
		GameStatus:         string(model.GameCreated),
	})
	if err != nil {
		log.Println("Error occurred creating game in local db: ", err)
		return nil, ErrInternal
	}

	if err = tx.Commit(); err != nil {
		log.Println("Error occurred trying to commit tx: ", err)
		return nil, ErrInternal
	}

	return game, nil
}

func buildGame(userID uuid.UUID, rows int, columns int, mines int) *model.Game {
	game := &model.Game{
		ID:                 uuid.New(),
		UserID:             userID,
		Status:             model.GameCreated,
		Rows:               rows,
		Columns:            columns,
		Mines:              mines,
		AccumulatedSeconds: 0,
		MinesLeft:          mines,
		Board:              nil,
		CellAmount:         rows * columns,
	}
	setBoard(game, createMines(game.Rows, game.Columns, game.Mines))
	return game
}

func setBoard(game *model.Game, createdMines map[int]bool) {
	game.Board = make([][]model.Cell, game.Rows)
	mineCellsCache := make([]*model.Cell, game.Mines)
	k := 0
	for i := range game.Board {
		game.Board[i] = make([]model.Cell, game.Columns)
		for j := range game.Board[i] {
			ci := getCellIndex(game.Rows, game.Columns, i, j)
			cell := model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Row:    i,
				Column: j,
				Game:   game,
			}
			if _, ok := createdMines[ci]; ok {
				cell.Type = model.MineCell
				mineCellsCache[k] = &cell
				k++
			}
			game.Board[i][j] = cell
		}
	}

	// increment value on all cells surrounding each mine
	for _, mc := range mineCellsCache {
		for _, sc := range mc.Surrounding() {
			if sc.Type == model.NumberCell {
				sc.Value++
			}
		}
	}
}

func createMines(rows int, columns int, mines int) map[int]bool {
	amount := rows * columns

	createdMines := make(map[int]bool)
	for i := 0; i < mines; i++ {
		for {
			rand.Seed(time.Now().UnixNano())
			newMine := rand.Intn(amount)
			if _, ok := createdMines[newMine]; !ok {
				createdMines[newMine] = true
				break
			}
		}
	}
	return createdMines
}

func getCellIndex(rows int, columns int, row int, column int) int {
	return columns*row + column
}
