package main

// Possibility is the representation of valid possibilities of a cell.
type Possibility map[string]bool

// EmptyPossibility creates a possibility set with no valid possibilities.
func EmptyPossibility() Possibility {
	m := make(Possibility)
	return m
}

// FullPossibility creates a possibility set with all valid possibilities.
func FullPossibility() Possibility {
	m := make(Possibility)
	m["1"] = true
	m["2"] = true
	m["3"] = true
	m["4"] = true
	m["5"] = true
	m["6"] = true
	m["7"] = true
	m["8"] = true
	m["9"] = true
	return m
}
