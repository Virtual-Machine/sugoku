package main

type Possibility map[string]bool


func EmptyPossibility() Possibility{
	m := make(Possibility)
	return m
}

func FullPossibility() Possibility{
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

func EqualPossibilities(pos1 Possibility, pos2 Possibility) bool {
	ret := true
	for k := range pos1 {
		if pos2[k] != true {
			ret = false
		}
	}
	for k := range pos2 {
		if pos1[k] != true {
			ret = false
		}
	}
	return ret
}

func (p *Possibility) RemovePossibility(val string){
	delete(*p, val)
}

func (p *Possibility) GetPossibility() string {
	keys := make([]string, len(*p))
	for k := range *p {
		keys = append(keys, k)
	}
	return keys[1]
}