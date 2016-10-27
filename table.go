package main

import ( 
	"fmt"
)

type Table struct {
	cells []Cell
	rowCells [10][]*Cell
	colCells [10][]*Cell
	boxCells [10][]*Cell
}

func (t *Table) PrintBoard(){
	counter := 1
	for _, v := range t.cells {
		fmt.Print(v.val, " ")

		counter++
		if counter == 10 {
			counter = 1
			fmt.Println("")
		}
	}
	fmt.Println("")
}

func (t *Table) SimplifyBoard() bool {
	simplified := false
	var rowCells [10][]*Cell
	var colCells [10][]*Cell
	var boxCells [10][]*Cell
	var rows [10][]string
	var cols [10][]string
	var boxes [10][]string
	for i, v := range t.cells {
		if v.val != "X" {
			rows[v.row] = append(rows[v.row], v.val)
			cols[v.col] = append(cols[v.col], v.val)
			boxes[v.box] = append(boxes[v.box], v.val)
		} else {
			rowCells[v.row] = append(rowCells[v.row], &t.cells[i])
			colCells[v.col] = append(colCells[v.col], &t.cells[i])
			boxCells[v.box] = append(boxCells[v.box], &t.cells[i])
		}
	}
	for rowNum, row := range rowCells {
		for _, cell := range row {
			for _, elimVal := range rows[rowNum]{
				cell.possibilities.RemovePossibility(elimVal)
			}
		}
	}
	for colNum, col := range colCells {
		for _, cell := range col {
			for _, elimVal := range cols[colNum]{
				cell.possibilities.RemovePossibility(elimVal)
			}
		}
	}
	for boxNum, box := range boxCells {
		for _, cell := range box {
			for _, elimVal := range boxes[boxNum]{
				cell.possibilities.RemovePossibility(elimVal)
			}
		}
	}

	// Exclusive Pair Elimination - (If two cells in a row, column, or box, 
	// exclusively contain the same pair of options, no other cells in that 
	// row, column, or box can possibly be either value)
	// 
	// AND Exclusive Triple Elimination - (If three cells in a row, column, or box, 
	// exclusively contain the same triplets of options, no other cells in that 
	// row, column, or box can possibly be either value)
	for _, row := range rowCells {
		for i1, cell := range row {
			for i2, cell2 := range row {
				if i1 != i2 && len(cell.possibilities) == 2 && len(cell2.possibilities) == 2 {
					if EqualPossibilities(cell.possibilities, cell2.possibilities){
						var keys []string
						for key := range cell.possibilities {
							keys = append(keys, key)
						}
						for _, refs := range row {
							if refs != cell && refs != cell2 {
								refs.possibilities.RemovePossibility(keys[0])
								refs.possibilities.RemovePossibility(keys[1])
							}
						}
					} 
				}
				for i3, cell3 := range row {
					if i1 != i2 && i1 != i3 && i2 != i3 && len(cell.possibilities) > 1 && len(cell.possibilities) < 4 && len(cell2.possibilities) > 1 && len(cell2.possibilities) < 4 && len(cell3.possibilities) > 1 && len(cell3.possibilities) < 4 {
						if EqualTriplets(cell.possibilities, cell2.possibilities, cell3.possibilities){
							triplet := GetTriplet(cell.possibilities, cell2.possibilities, cell3.possibilities)
							for _, refs := range row {
								if refs != cell && refs != cell2 && refs != cell3 {
									refs.possibilities.RemovePossibility(triplet[0])
									refs.possibilities.RemovePossibility(triplet[1])
									refs.possibilities.RemovePossibility(triplet[2])
								}
							}
						} 
					}
				}
			}
		}
	}

	for _, col := range colCells {
		for i1, cell := range col {
			for i2, cell2 := range col {
				if i1 != i2 && len(cell.possibilities) == 2 && len(cell2.possibilities) == 2 {
					if EqualPossibilities(cell.possibilities, cell2.possibilities){
						var keys []string
						for key := range cell.possibilities {
							keys = append(keys, key)
						}
						for _, refs := range col {
							if refs != cell && refs != cell2 {
								refs.possibilities.RemovePossibility(keys[0])
								refs.possibilities.RemovePossibility(keys[1])
							}
						}
					} 
				}
				for i3, cell3 := range col {
					if i1 != i2 && i1 != i3 && i2 != i3 && len(cell.possibilities) > 1 && len(cell.possibilities) < 4 && len(cell2.possibilities) > 1 && len(cell2.possibilities) < 4 && len(cell3.possibilities) > 1 && len(cell3.possibilities) < 4 {
						if EqualTriplets(cell.possibilities, cell2.possibilities, cell3.possibilities){
							triplet := GetTriplet(cell.possibilities, cell2.possibilities, cell3.possibilities)
							for _, refs := range col {
								if refs != cell && refs != cell2 && refs != cell3 {
									refs.possibilities.RemovePossibility(triplet[0])
									refs.possibilities.RemovePossibility(triplet[1])
									refs.possibilities.RemovePossibility(triplet[2])
								}
							}
						} 
					}
				}
			}
		}
	}

	for _, box := range boxCells {
		for i1, cell := range box {
			for i2, cell2 := range box {
				if i1 != i2 && len(cell.possibilities) == 2 && len(cell2.possibilities) == 2 {
					if EqualPossibilities(cell.possibilities, cell2.possibilities){
						var keys []string
						for key := range cell.possibilities {
							keys = append(keys, key)
						}
						for _, refs := range box {
							if refs != cell && refs != cell2 {
								refs.possibilities.RemovePossibility(keys[0])
								refs.possibilities.RemovePossibility(keys[1])
							}
						}
					} 
				}
				for i3, cell3 := range box {
					if i1 != i2 && i1 != i3 && i2 != i3 && len(cell.possibilities) > 1 && len(cell.possibilities) < 4 && len(cell2.possibilities) > 1 && len(cell2.possibilities) < 4 && len(cell3.possibilities) > 1 && len(cell3.possibilities) < 4 {
						if EqualTriplets(cell.possibilities, cell2.possibilities, cell3.possibilities){
							triplet := GetTriplet(cell.possibilities, cell2.possibilities, cell3.possibilities)
							for _, refs := range box {
								if refs != cell && refs != cell2 && refs != cell3 {
									refs.possibilities.RemovePossibility(triplet[0])
									refs.possibilities.RemovePossibility(triplet[1])
									refs.possibilities.RemovePossibility(triplet[2])
								}
							}
						} 
					}
				}
			}
		}
	}	


	for i, v := range t.cells {
		if len(v.possibilities) == 1 {
			simplified = true
			deduction := v.possibilities.GetPossibility()
			t.cells[i].val = deduction
			t.cells[i].possibilities = EmptyPossibility()
		}
	}

	t.PrintBoard()
	return simplified
}
