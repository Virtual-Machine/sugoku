package main

import ( 
	"fmt"
	"strconv"
)

type Table struct {
	cells []Cell
	containers Container
}

func (t *Table) PrintBoard(){
	counter := 1
	for _, v := range t.cells {
		if v.val == "X" {
			fmt.Print("\033[31m", v.val, "\033[39m ")
		} else {
			fmt.Print("\033[32m", v.val, "\033[39m ")
		}
		counter++
		if counter == 10 {
			counter = 1
			fmt.Println("")
		}
	}
	fmt.Println("")
}

func (t *Table) PrintPossibilities(){
	counter := 1
	for _, v := range t.cells {
		if v.val == "X" {
			fmt.Print("\033[31m", v.possibilities, "\033[39m ")
		} else {
			fmt.Print("\033[32m", v.possibilities, "\033[39m ")
		}
		counter++
		if counter == 10 {
			counter = 1
			fmt.Println("")
		}
	}
	fmt.Println("")
}

func (t *Table) RemoveOptions(){
	for _, v := range t.cells {
		if v.val != "X" {
			t.ClearOption(v.row, v.col, v.box, v.val)
		}
	}
}

func (t *Table) ClearOption(row int, col int, box int, val string){
	rowCells := t.containers.GetRow(row)
	colCells := t.containers.GetCol(col)
	boxCells := t.containers.GetBox(box)
	for _, cell := range rowCells {
		cell.RemovePossibility(val)
	}
	for _, cell := range colCells {
		cell.RemovePossibility(val)
	}
	for _, cell := range boxCells {
		cell.RemovePossibility(val)
	}
}

func (t *Table) SimplifyBoard(mode string) bool {
	// This is a modular algorithm, that will only take one action per round.
	// First algorithm that is able to simplify wins.
	// Board remembers effects of previous rounds to allow more complex deduction logic.
	simplified := false
	defer t.PrintBoard()
	
	// 1. Check for exclusive singles.
	simplified = t.CheckExclusiveSingles(mode)
	if simplified { return simplified }

	// 2. Hidden singles.
	simplified = t.CheckHiddenSingles(mode)
	if simplified { return simplified }

	// 3. Exclusive pairs.
	simplified = t.CheckExclusivePairs(mode)
	if simplified { return simplified }

	// 4. Hidden pairs.
	simplified = t.CheckHiddenPairs(mode)
	if simplified { return simplified }




	return simplified
}

func Explain(mode string, message string){
	if mode == "EXPLAIN" {
		fmt.Println(message)
	}
}

func (t *Table) CheckExclusiveSingles(mode string) bool {
	simplified := false
	for i, cell := range t.cells {
		if cell.PossibilityCount() == 1 {
			simplified = true
			value := cell.GetPossibility()
			t.cells[i].SetValue(value)
			t.ClearOption(cell.row, cell.col, cell.box, value)
			Explain(mode, "Found Exclusive Single... Row: " + strconv.Itoa(cell.row) + " Col: " + strconv.Itoa(cell.col) + " Val: " + value)
		}
	}
	return simplified
}

func (t *Table) CheckHiddenSingles(mode string) bool {
	simplified := false
	for _, row := range t.containers.rows {
		list := GetPossibilityList(row)
		for k, v := range list {
			if v == 1 {
				for _, cell := range row {
					if cell.HasPossibility(k){
						simplified = true
						cell.SetValue(k)
						t.ClearOption(cell.row, cell.col, cell.box, k)
						Explain(mode, "Found Hidden Row Single... Row: " + strconv.Itoa(cell.row) + " Col: " + strconv.Itoa(cell.col) + " Val: " + k)
					}
				}
			}
		} 
	}
	for _, col := range t.containers.cols {
		list := GetPossibilityList(col)
		for k, v := range list {
			if v == 1 {
				for _, cell := range col {
					if cell.HasPossibility(k){
						simplified = true
						cell.SetValue(k)
						t.ClearOption(cell.row, cell.col, cell.box, k)
						Explain(mode, "Found Hidden Col Single... Row: " + strconv.Itoa(cell.row) + " Col: " + strconv.Itoa(cell.col) + " Val: " + k)
					}
				}
			}
		} 
	}
	for _, box := range t.containers.boxes {
		list := GetPossibilityList(box)
		for k, v := range list {
			if v == 1 {
				for _, cell := range box {
					if cell.HasPossibility(k){
						simplified = true
						cell.SetValue(k)
						t.ClearOption(cell.row, cell.col, cell.box, k)
						Explain(mode, "Found Hidden Box Single... Row: " + strconv.Itoa(cell.row) + " Col: " + strconv.Itoa(cell.col) + " Val: " + k)
					}
				}
			}
		} 
	}
	return simplified
}

func (t *Table) CheckExclusivePairs(mode string) bool {
	simplified := false
	for _, row := range t.containers.rows {
		for i1, cell1 := range row {
			for i2, cell2 := range row {
				if i1 != i2 && cell1.PossibilityCount() > 1 && cell2.PossibilityCount() > 1 && ContainsCompatibleSet(*cell1, *cell2){
					possibilities := cell1.GetPossibilities()
					
					for i3, cell3 := range row {
						if i3 != i1 && i3 != i2 {
							if cell3.HasPossibility(possibilities[0]) || cell3.HasPossibility(possibilities[1]){
								cell3.RemovePossibility(possibilities[0])
								cell3.RemovePossibility(possibilities[1])
								simplified = true
							}
						}
					}
				}
			}
		}
	}
	for _, col := range t.containers.cols {
		for i1, cell1 := range col {
			for i2, cell2 := range col {
				if i1 != i2 && cell1.PossibilityCount() > 1 && cell2.PossibilityCount() > 1 && ContainsCompatibleSet(*cell1, *cell2){
					possibilities := cell1.GetPossibilities()
					
					for i3, cell3 := range col {
						if i3 != i1 && i3 != i2 {
							if cell3.HasPossibility(possibilities[0]) || cell3.HasPossibility(possibilities[1]){
								cell3.RemovePossibility(possibilities[0])
								cell3.RemovePossibility(possibilities[1])
								simplified = true
							}
						}
					}
				}
			}
		}
	}
	for _, box := range t.containers.boxes {
		for i1, cell1 := range box {
			for i2, cell2 := range box {
				if i1 != i2 && cell1.PossibilityCount() > 1 && cell2.PossibilityCount() > 1 && ContainsCompatibleSet(*cell1, *cell2){
					possibilities := cell1.GetPossibilities()
					
					for i3, cell3 := range box {
						if i3 != i1 && i3 != i2 {
							if cell3.HasPossibility(possibilities[0]) || cell3.HasPossibility(possibilities[1]){
								cell3.RemovePossibility(possibilities[0])
								cell3.RemovePossibility(possibilities[1])
								simplified = true
							}
						}
					}
				}
			}
		}
	}

	if simplified {
		Explain(mode, "Excluded based on exclusive pairs")
	}
	return simplified
}

func (t *Table) CheckHiddenPairs(mode string) bool {
	simplified := false
	for _, row := range t.containers.rows {
		list := GetPossibilityList(row)

		for k, v := range list {
			for k2, v2 := range list {
				if k != k2 && v == 2 && v2 == 2 {
					for i1, cell := range row {
						for i2, cell2 := range row {
							if i1 != i2 {
								if cell.HasPossibility(k) && cell.HasPossibility(k2) && cell2.HasPossibility(k) && cell2.HasPossibility(k2) {
									if cell.PossibilityCount() > 2 || cell2.PossibilityCount() > 2 {
										simplified = true
										cell.SetPossibilities(k, k2)
										cell2.SetPossibilities(k, k2)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	for _, col := range t.containers.cols {
		list := GetPossibilityList(col)

		for k, v := range list {
			for k2, v2 := range list {
				if k != k2 && v == 2 && v2 == 2 {
					for i1, cell := range col {
						for i2, cell2 := range col {
							if i1 != i2 {
								if cell.HasPossibility(k) && cell.HasPossibility(k2) && cell2.HasPossibility(k) && cell2.HasPossibility(k2) {
									if cell.PossibilityCount() > 2 || cell2.PossibilityCount() > 2 {
										simplified = true
										cell.SetPossibilities(k, k2)
										cell2.SetPossibilities(k, k2)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	for _, box := range t.containers.boxes {
		list := GetPossibilityList(box)

		for k, v := range list {
			for k2, v2 := range list {
				if k != k2 && v == 2 && v2 == 2 {
					for i1, cell := range box {
						for i2, cell2 := range box {
							if i1 != i2 {
								if cell.HasPossibility(k) && cell.HasPossibility(k2) && cell2.HasPossibility(k) && cell2.HasPossibility(k2) {
									if cell.PossibilityCount() > 2 || cell2.PossibilityCount() > 2 {
										simplified = true
										cell.SetPossibilities(k, k2)
										cell2.SetPossibilities(k, k2)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if simplified {
		Explain(mode, "Excluded based on hidden pairs")
	}
	return simplified
}
