package main

import (
	"encoding/csv"
	"log"
	"io"
	"bufio"
	"os"
)

func main() {
	in, err := os.Open("puzzle.csv")
	
	if err != nil {
		log.Fatal("Could not open puzzle.csv")
	}
	
	r := csv.NewReader(bufio.NewReader(in))

	var board Table
	round := 1
	row := 1
	box := 0
	col := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for i, v := range record {
			col = i + 1
			box = getBox(col, row)
			var cell Cell
			if v == "X" {
				cell = Cell{row, col, box, v, FullPossibility()}
			} else {
				cell = Cell{row, col, box, v, EmptyPossibility()}
			}
			
			board.cells = append(board.cells, cell)
		}
		row++
	}
	for board.SimplifyBoard() {
		log.Println("Simplying round", round)
		round++
	}
	board.PrintBoard()
}

func getBox(col int, row int) int {
	if col < 4 && row < 4 {
		return 1
	}
	if col < 4 && row > 6 {
		return 7
	}
	if col > 6 && row < 4 {
		return 3
	}
	if col > 6 && row > 6 {
		return 9
	}
	if col < 4 {
		return 4
	}
	if row < 4 {
		return 2
	}
	if col > 6 {
		return 6
	}
	if row > 6 {
		return 8
	}
	return 5
}