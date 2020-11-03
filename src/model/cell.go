package model

// Cell represents a cell in the board
type Cell struct {
	Type   CellType   `json:"t,omitempty"`
	Action CellAction `json:"a,omitempty"`
	Value  int        `json:"v,omitempty"`
	Row    int        `json:"-"`
	Column int        `json:"-"`
	Game   *Game      `json:"-"`
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

// Surrounding returns all surrounding cells
func (c *Cell) Surrounding() []*Cell {
	surrounding := []*Cell{}

	if s := c.Upper(); s != nil {
		surrounding = append(surrounding, s)
	}
	if s := c.UpperRight(); s != nil {
		surrounding = append(surrounding, s)
	}
	if s := c.Right(); s != nil {
		surrounding = append(surrounding, s)
	}
	if s := c.LowerRight(); s != nil {
		surrounding = append(surrounding, s)
	}
	if s := c.Lower(); s != nil {
		surrounding = append(surrounding, s)
	}
	if s := c.LowerLeft(); s != nil {
		surrounding = append(surrounding, s)
	}
	if s := c.Left(); s != nil {
		surrounding = append(surrounding, s)
	}
	if s := c.UpperLeft(); s != nil {
		surrounding = append(surrounding, s)
	}

	return surrounding
}

// Upper gives you the upper cell relative to this one
func (c *Cell) Upper() *Cell {
	if c.Row == 0 {
		return nil
	}
	return &c.Game.Board[c.Row-1][c.Column]
}

// UpperRight gives you the upper right cell relative to this one
func (c *Cell) UpperRight() *Cell {
	if c.Row == 0 {
		return nil
	}
	if c.Column == c.Game.Columns-1 {
		return nil
	}
	return &c.Game.Board[c.Row-1][c.Column+1]
}

// Right gives you the right cell relative to this one
func (c *Cell) Right() *Cell {
	if c.Column == c.Game.Columns-1 {
		return nil
	}
	return &c.Game.Board[c.Row][c.Column+1]
}

// LowerRight gives you the lower right cell relative to this one
func (c *Cell) LowerRight() *Cell {
	if c.Row == c.Game.Rows-1 {
		return nil
	}
	if c.Column == c.Game.Columns-1 {
		return nil
	}
	return &c.Game.Board[c.Row+1][c.Column+1]
}

// Lower gives you the lower cell relative to this one
func (c *Cell) Lower() *Cell {
	if c.Row == c.Game.Rows-1 {
		return nil
	}
	return &c.Game.Board[c.Row+1][c.Column]
}

// LowerLeft gives you the lower left cell relative to this one
func (c *Cell) LowerLeft() *Cell {
	if c.Row == c.Game.Rows-1 {
		return nil
	}
	if c.Column == 0 {
		return nil
	}
	return &c.Game.Board[c.Row+1][c.Column-1]
}

// Left gives you the left cell relative to this one
func (c *Cell) Left() *Cell {
	if c.Column == 0 {
		return nil
	}
	return &c.Game.Board[c.Row][c.Column-1]
}

// UpperLeft gives you the upper left cell relative to this one
func (c *Cell) UpperLeft() *Cell {
	if c.Row == 0 {
		return nil
	}
	if c.Column == 0 {
		return nil
	}
	return &c.Game.Board[c.Row-1][c.Column-1]
}
