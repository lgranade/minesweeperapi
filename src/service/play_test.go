// +build unit

package service

import (
	"testing"

	"github.com/lgranade/minesweeperapi/model"
)

func TestPlayStepEmptyCell(t *testing.T) {

	game := *game6By6
	game.Board = [][]model.Cell{}
	setBoard(&game, mines6By6)

	calculatePlay(&game, 0, 3, StepPlay)

	stepped := 0

	for i := range game.Board {
		for j := range game.Board[i] {
			cell := &game.Board[i][j]
			if cell.Action == model.StepAction {
				stepped++
			}
		}
	}

	// log.Println("stepped: ", stepped)

	if stepped != 21 {
		t.Errorf("Bad stepped amount")
		return
	}
}
