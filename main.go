package main

import (
	"log"
)

func main() {
	initialS := state{3, 3, true, nil}

	finalS, err := general(initialS)
	if err != nil {
		log.Fatal(err)
	}

	var path []state
	s := finalS
	path = append(path, s)
	for s.parentS != nil {
		s = *s.parentS
		path = append(path, s)
	}

	depth := len(path) - 1
	for i := depth; i >= 0; i-- {
		s := path[i]
		if s.isGoal() {
			println(s.String())
		} else {
			print(s.String(), "->")
		}
	}
}
