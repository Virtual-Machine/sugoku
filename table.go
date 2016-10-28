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
