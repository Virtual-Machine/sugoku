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
	keys := make([]string, len(c.possibilities))
	for k := range c.possibilities {
		keys = append(keys, k)
	}
	return keys[1]
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