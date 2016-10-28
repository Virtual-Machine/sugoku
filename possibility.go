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

func EqualQuadruplets(pos1 Possibility, pos2 Possibility, pos3 Possibility, pos4 Possibility) bool {
	ret := true
	var keys []string
	for k := range pos1 {
		keys = append(keys, k)
	}
	for k := range pos2 {
		keys = append(keys, k)
	}
	for k := range pos3 {
		keys = append(keys, k)
	}
	for k := range pos4 {
		keys = append(keys, k)
	}
	keys = removeDuplicates(keys)
	if len(keys) != 4 {
		ret = false
	}
	return ret
}

func GetQuadruplet(pos1 Possibility, pos2 Possibility, pos3 Possibility, pos4 Possibility) []string{
	var keys []string
	for k := range pos1 {
		keys = append(keys, k)
	}
	for k := range pos2 {
		keys = append(keys, k)
	}
	for k := range pos3 {
		keys = append(keys, k)
	}
	for k := range pos4 {
		keys = append(keys, k)
	}
	keys = removeDuplicates(keys)
	return keys
}

func EqualTriplets(pos1 Possibility, pos2 Possibility, pos3 Possibility) bool {
	ret := true
	var keys []string
	for k := range pos1 {
		keys = append(keys, k)
	}
	for k := range pos2 {
		keys = append(keys, k)
	}
	for k := range pos3 {
		keys = append(keys, k)
	}
	keys = removeDuplicates(keys)
	if len(keys) != 3 {
		ret = false
	}
	return ret
}

func GetTriplet(pos1 Possibility, pos2 Possibility, pos3 Possibility) []string{
	var keys []string
	for k := range pos1 {
		keys = append(keys, k)
	}
	for k := range pos2 {
		keys = append(keys, k)
	}
	for k := range pos3 {
		keys = append(keys, k)
	}
	keys = removeDuplicates(keys)
	return keys
}

func removeDuplicates(xs []string) []string {
	found := make(map[string]bool)
	var newArray []string
	for _, x := range xs {
		if !found[x] {
			found[x] = true
			newArray = append(newArray, x)
		}
	}
	return newArray
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