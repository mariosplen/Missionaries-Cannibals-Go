package main

import "fmt"

type state struct {
	ml      int
	cl      int
	bIsLeft bool

	parentS *state
}

func (s state) isGoal() bool {
	return s.ml == 0 && s.cl == 0
}

func (s state) isValid() bool {
	return (s.ml == 0 || s.ml >= s.cl) && (3-s.ml == 0 || 3-s.ml >= 3-s.cl) && (s.cl >= 0 && s.ml >= 0 && s.cl <= 3 && s.ml <= 3)
}

func newState(m move, cState state) state {
	var nState state

	if m.leftwards {
		nState = state{cState.ml + m.nMiss, cState.cl + m.nCann, true, nil}
	} else {
		nState = state{cState.ml - m.nMiss, cState.cl - m.nCann, false, nil}
	}
	return nState
}

func (s state) spawnChildren(frontier []state, closed []state) []state {
	var children []state
	var nState state
	for _, m := range moves {
		if m.leftwards != s.bIsLeft {
			nState = newState(m, s)
			nState.parentS = &s
			if !contains(nState, closed) && !contains(nState, frontier) {
				children = append(children, nState)
			}

		}
	}
	return children
}

func (s state) String() string {
	if s.bIsLeft == true {
		return fmt.Sprintf("(%v,%v,L)", s.ml, s.cl)
	} else {
		return fmt.Sprintf("(%v,%v,R)", s.ml, s.cl)
	}

}
