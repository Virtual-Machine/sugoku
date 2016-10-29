package main

type Cell struct {
	row		int
	col     int
	box     int
	val     string
	possibilities Possibility
}

func (c *Cell) RemovePossibility(val string){
	delete(c.possibilities, val)
}

func (c *Cell) GetPossibility() string {
	var keys []string
	for k := range c.possibilities {
		keys = append(keys, k)
	}
	return keys[0]
}

func (c *Cell) SetPossibilities(possibilities ...string) {
	c.possibilities = EmptyPossibility()
	for _, i := range possibilities {
		c.possibilities[i] = true
	}
}

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

func (c *Cell) HasPossibility(val string) bool{
	return c.possibilities[val]
}

func (c *Cell) PossibilityCount() int{
	return len(c.possibilities)
}

func (c *Cell) SetValue(val string){
	c.val = val
	c.possibilities = EmptyPossibility()
}

func ContainsCompatibleSet(cells ...Cell) bool{

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