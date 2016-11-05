package main

// Container is a virtual structure holding pointers to specific groupings in the 81 cell matrix. This allows manipulations on the 9 cell rows, columns, and boxes.
type Container struct {
	rows  [9][]*Cell
	cols  [9][]*Cell
	boxes [9][]*Cell
}

// GetRow is a helper function to get the container of a given row.
func (c *Container) GetRow(row int) []*Cell {
	return c.rows[row-1]
}

// GetCol is a helper function to get the container of a given column.
func (c *Container) GetCol(col int) []*Cell {
	return c.cols[col-1]
}

// GetBox is a helper function to get the container of a given box.
func (c *Container) GetBox(box int) []*Cell {
	return c.boxes[box-1]
}
