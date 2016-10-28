package main


type Container struct {
	rows [9][]*Cell
	cols [9][]*Cell
	boxes [9][]*Cell
}

func (c *Container) GetRow(row int) []*Cell {
	return c.rows[row - 1]
}

func (c *Container) GetCol(col int) []*Cell {
	return c.cols[col - 1]
}

func (c *Container) GetBox(box int) []*Cell {
	return c.boxes[box - 1]
}