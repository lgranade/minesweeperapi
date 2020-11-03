// +build unit

package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/lgranade/minesweeperapi/model"
)

var game6By6 = &model.Game{
	ID:                 uuid.New(),
	UserID:             uuid.New(),
	Status:             model.GameCreated,
	Rows:               6,
	Columns:            6,
	Mines:              5,
	AccumulatedSeconds: 0,
	MinesLeft:          5,
	Board: [][]model.Cell{
		[]model.Cell{
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    0,
				Column: 0,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    0,
				Column: 1,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    0,
				Column: 2,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  0,
				Row:    0,
				Column: 3,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    0,
				Column: 4,
			},
			model.Cell{
				Type:   model.MineCell,
				Action: model.NoAction,
				Value:  0,
				Row:    0,
				Column: 5,
			},
		},

		[]model.Cell{
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    1,
				Column: 0,
			},
			model.Cell{
				Type:   model.MineCell,
				Action: model.NoAction,
				Value:  0,
				Row:    1,
				Column: 1,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    1,
				Column: 2,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  0,
				Row:    1,
				Column: 3,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    1,
				Column: 4,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    1,
				Column: 5,
			},
		},

		[]model.Cell{
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    2,
				Column: 0,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    2,
				Column: 1,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    2,
				Column: 2,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  0,
				Row:    2,
				Column: 3,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  0,
				Row:    2,
				Column: 4,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  0,
				Row:    2,
				Column: 5,
			},
		},

		[]model.Cell{
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    3,
				Column: 0,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    3,
				Column: 1,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    3,
				Column: 2,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  0,
				Row:    3,
				Column: 3,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  0,
				Row:    3,
				Column: 4,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  0,
				Row:    3,
				Column: 5,
			},
		},

		[]model.Cell{
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  2,
				Row:    4,
				Column: 0,
			},
			model.Cell{
				Type:   model.MineCell,
				Action: model.NoAction,
				Value:  0,
				Row:    4,
				Column: 1,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  3,
				Row:    4,
				Column: 2,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    4,
				Column: 3,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    4,
				Column: 4,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  0,
				Row:    4,
				Column: 5,
			},
		},

		[]model.Cell{
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  2,
				Row:    5,
				Column: 0,
			},
			model.Cell{
				Type:   model.MineCell,
				Action: model.NoAction,
				Value:  0,
				Row:    5,
				Column: 1,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  3,
				Row:    5,
				Column: 2,
			},
			model.Cell{
				Type:   model.MineCell,
				Action: model.NoAction,
				Value:  0,
				Row:    5,
				Column: 3,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  1,
				Row:    5,
				Column: 4,
			},
			model.Cell{
				Type:   model.NumberCell,
				Action: model.NoAction,
				Value:  0,
				Row:    5,
				Column: 5,
			},
		},
	},
	CellAmount: 6 * 6,
}

var mines6By6 = map[int]bool{
	5:  true,
	7:  true,
	25: true,
	31: true,
	33: true,
}

func init() {
	// reference all cells to the board, this can't be done statically
	for i := range game6By6.Board {
		for j := range game6By6.Board[i] {
			game6By6.Board[i][j].Game = game6By6
		}
	}
}

func TestGetCellIndex(t *testing.T) {
	if getCellIndex(6, 6, 1, 1) != 7 {
		t.Errorf("Bad index")
		return
	}
}

func TestCreateMines(t *testing.T) {
	mines := createMines(6, 6, 4)
	// log.Println("created mines: ", mines)
	if len(mines) != 4 {
		t.Errorf("Wrong amount of mines created")
		return
	}
	for k := range mines {
		if k < 0 || k > 35 {
			t.Errorf("Bar mine cell index")
			return
		}
	}
}

func TestSetBoard(t *testing.T) {

	game := *game6By6
	game.Board = [][]model.Cell{}
	setBoard(&game, mines6By6)

	for i := range game6By6.Board {
		for j := range game6By6.Board[i] {
			if game6By6.Board[i][j].Value != game.Board[i][j].Value {
				t.Errorf("Values don't match, row: %d, col: %d, value: %d", i, j, game.Board[i][j].Value)
				return
			}
		}
	}
}
