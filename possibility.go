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