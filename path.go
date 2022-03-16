package main

import (
	"fmt"
)

// getPath returns a state slice from state up to root
func (s *state) getPath() stateSlice {
	var path []state
	for s.parentS != nil {
		path = append(path, *s)
		s = s.parentS
	}
	path = append(path, *s)
	return path
}

func (p stateSlice) getDepth() int {
	return len(p) - 1
}

func (p stateSlice) printPath() {
	depth := len(p) - 1
	for i := depth; i >= 0; i-- {
		s := p[i]
		if s.isGoal() {
			fmt.Println(s.String())
		} else {
			fmt.Print(s.String(), "->")
		}
	}
}

func (p stateSlice) printDepth() {
	fmt.Println("Depth:", p.getDepth())
}
