package main

// Cell is the representation of one cell in the 81 cell matrix.
type Cell struct {
	row           int
	col           int
	box           int
	val           string
	possibilities Possibility
}

// RemovePossibility allows a cell to exclude a possibility from its set.
func (c *Cell) RemovePossibility(val string) {
	delete(c.possibilities, val)
}

// GetPossibility returns the first possibility in the set, used to retrieve final value.
func (c *Cell) GetPossibility() string {
	var keys []string
	for k := range c.possibilities {
		keys = append(keys, k)
	}
	return keys[0]
}

// SetPossibilities forces a cells possibility to the given set of possibilities.
func (c *Cell) SetPossibilities(possibilities ...string) {
	c.possibilities = EmptyPossibility()
	for _, i := range possibilities {
		c.possibilities[i] = true
	}
}

// OnlyPossibilities forces a cells possibility to the given set of possibilities but does not add to the set if the possibility is already excluded.
func (c *Cell) OnlyPossibilities(possibilities ...string) {
	for k := range c.possibilities {
		hit := false
		for _, i2 := range possibilities {
			if i2 == k {
				hit = true
			}
		}
		if !hit {
			c.RemovePossibility(k)
		}
	}
}

// GetPossibilities returns the combined set of possibilities in a group.
func (c *Cell) GetPossibilities(otherCells ...Cell) []string {
	var keys []string
	for k := range c.possibilities {
		keys = append(keys, k)
	}
	for _, cell := range otherCells {
		for k := range cell.possibilities {
			hit := false
			for _, key := range keys {
				if key == k {
					hit = true
				}
			}
			if !hit {
				keys = append(keys, k)
			}
		}
	}
	return keys
}

// HasPossibility returns true if possibility is possible in a cell.
func (c *Cell) HasPossibility(val string) bool {
	return c.possibilities[val]
}

// PossibilityCount returns the total number of possibilities left for a cell.
func (c *Cell) PossibilityCount() int {
	return len(c.possibilities)
}

// SetValue removes all remaining possibilities and sets the value for the cell.
func (c *Cell) SetValue(val string) {
	c.val = val
	c.possibilities = EmptyPossibility()
}

// ContainsCompatibleSet determines if a group of cells contain a combined possibility set that is equal to the number of cells, indicating that they are an exclusive set.
func ContainsCompatibleSet(cells ...Cell) bool {

	count := len(cells)
	possibilityList := make(map[string]bool)
	for _, cell := range cells {
		for k := range cell.possibilities {
			possibilityList[k] = true
		}
	}
	if len(possibilityList) == count {
		return true
	}
	return false
}
