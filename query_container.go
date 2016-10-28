package main

func GetPossibilityList(nineGroup []*Cell) map[string]int {
	list := make(map[string]int)
	for _, v := range nineGroup {
		for k := range v.possibilities {
			list[k] += 1
		}
	}
	return list
}