package main

import "fmt"

type state struct {
	ml      int  // Missionaries on the left
	cl      int  // Canibals on the left
	bIsLeft bool // Boat is on the left

	parentS *state // Parent state
}

// Returns true if the state is the final state
func (s state) isGoal() bool {
	return s.ml == 0 && s.cl == 0
}

func (s state) isValid() bool {
	return (s.ml == 0 || s.ml >= s.cl) && // Canibals should not be able to eat the missionaries on the left
		(3-s.ml == 0 || 3-s.ml >= 3-s.cl) && // Canibals should not be able to eat the missionaries on the right. Missionaries on the right are 3-ml and canibals on the right 3-cl
		(s.cl >= 0 && s.ml >= 0 && s.cl <= 3 && s.ml <= 3) // Their number should be beetween 0 and 3
}

// Using a move, creates an example state without checking if it is valid yet
func newState(m move, cState state) state {
	var nState state

	// Checks the position of boat in current state and creates the states acordingly
	// Can't move from left towards left, or from right towards right and
	// if moving towards left number of miss&can on the left increases and decreases by moving on the right. All acording to the m move
	if m.leftwards {
		nState = state{cState.ml + m.nMiss, cState.cl + m.nCann, true, nil}
	} else {
		nState = state{cState.ml - m.nMiss, cState.cl - m.nCann, false, nil}
	}
	return nState
}

// Returns a slice of children nodes, some may be invalid, we don't check for validity here
// Must take frontier and closedset as parameters, because we need to check if the supposed node is already explored, otherwise the node cannot be a child
func (s state) spawnChildren(frontier []state, closed []state) []state {
	var children []state
	var nState state
	for _, m := range moves {
		if m.leftwards != s.bIsLeft {
			nState = newState(m, s) // supposed node
			nState.parentS = &s
			if !contains(nState, closed) && !contains(nState, frontier) { // Checks if the supposed node is already explored by looking at closed and frontier
				children = append(children, nState)
			}
		}
	}
	return children
}

// toString() functionality
func (s state) String() string {
	if s.bIsLeft == true {
		return fmt.Sprintf("(%v,%v,L)", s.ml, s.cl)
	} else {
		return fmt.Sprintf("(%v,%v,R)", s.ml, s.cl)
	}
}
