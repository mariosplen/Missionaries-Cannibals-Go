package main

import "fmt"

type state struct {
	ml   int
	cl   int
	mr   int
	cr   int
	bPos pos

	parentS *state
}

type stateSlice []state

func (l stateSlice) contains(s state) bool {
	for _, a := range l {
		return a.cl == s.cl && a.ml == s.ml && a.bPos == s.bPos
	}
	return false
}

func (s state) String() string {
	if s.bPos == left {
		return fmt.Sprintf("(%v,%v,L)", s.ml, s.cl)
	} else {
		return fmt.Sprintf("(%v,%v,R)", s.ml, s.cl)
	}

}

// Constructor for state
func genState(ml int, cl int, bPos pos) *state {
	s := state{
		ml:   ml,
		cl:   cl,
		bPos: bPos,
	}
	s.mr = 3 - ml
	s.cr = 3 - cl
	s.parentS = nil
	return &s
}

func (s state) isGoal() bool {
	return s.ml == 0 && s.cl == 0
}

func (s state) isValid() bool {
	return (s.ml == 0 || s.ml >= s.cl) && (s.mr == 0 || s.mr >= s.cr)
}

func move(mov movement, currentState state) *state {
	var nextState *state
	if mov.towards == right {
		nextState = genState(
			currentState.ml-mov.nMiss,
			currentState.cl-mov.nCann,
			right,
		)
	} else {
		nextState = genState(
			currentState.ml+mov.nMiss,
			currentState.cl+mov.nCann,
			left,
		)
	}
	return nextState
}

func (s state) spawnNeighbours() []state {
	var neighbours []state
	var newState *state
	for _, m := range moves {
		if m.towards != s.bPos {
			newState = move(m, s)
			if newState.isValid() {
				newState.parentS = &s
				neighbours = append(neighbours, *newState)
			}
		}
	}
	return neighbours
}
