package main

import (
	"log"
)

func main() {
	// InitialState
	initialS := state{3, 3, true, nil}

	// Run general search algorithm
	finalS, err := general(initialS)
	if err != nil {
		log.Fatal(err)
	}

	// Create Path by going backwards from the FinalState up to the InitialState
	var path []state
	s := finalS
	path = append(path, s)
	for s.parentS != nil {
		s = *s.parentS
		path = append(path, s)
	}

	// Print the Path reversed
	depth := len(path) - 1
	for i := depth; i >= 0; i-- {
		s := path[i]
		if s.isGoal() {
			println(s.String())
		} else {
			print(s.String(), " -> ")
		}
	}
}
