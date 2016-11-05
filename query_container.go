package main

// GetPossibilityList takes a virtual row, column, or box container and returns a count of each numerical possibility.
func GetPossibilityList(nineGroup []*Cell) map[string]int {
	list := make(map[string]int)
	for _, v := range nineGroup {
		for k := range v.possibilities {
			list[k]++
		}
	}
	return list
}

// GetPossibilityListPartial a slice of cells and returns a count of each numerical possibility.
func GetPossibilityListPartial(cells ...Cell) map[string]int {
	list := make(map[string]int)
	for _, v := range cells {
		for k := range v.possibilities {
			list[k]++
		}
	}
	return list
}
