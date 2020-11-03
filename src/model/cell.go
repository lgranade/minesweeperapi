package model

// Cell represents a cell in the board
type Cell struct {
	Type   CellType   `json:"t"`
	Action CellAction `json:"a"`
	Value  int        `json:"v"`
}

// CellType is the cell type
type CellType string

const (
	// MineCell indicates this cell has a mine under it
	MineCell CellType = "m"

	// NumberCell indicates this cell has a number under it
	NumberCell CellType = "n"
)

// CellAction is the action the user executed over this cell
type CellAction string

const (
	// NoAction indicates the user hasn't done anything with this cell
	NoAction CellAction = "n"
	// StepAction indicates that the user has stepped on this cell or that it is uncovered
	StepAction CellAction = "s"
	// FlagAction indicates that the user has flagged this cell
	FlagAction CellAction = "f"
)
