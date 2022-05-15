package main

import (
	"errors"
	"log"
)

// General search algorithm from the book : Artificial Intelligence ISBN: 978-618-5196-44-8
// Explanation of the algorithm is also in the book
func general(initialState state) (state, error) {
	var closed []state
	var frontier []state

	frontier = append(frontier, initialState)
	currentState := frontier[0]

	for !currentState.isGoal() {
		frontier = delete(currentState, frontier)
		if !contains(currentState, closed) {
			next := currentState.spawnChildren(frontier, closed)
			frontier = insert(next, frontier)
			frontier = prune(frontier)
			frontier = reorder(frontier)
			closed = append(closed, currentState)
		}
		if len(frontier) == 0 {
			return initialState, errors.New("Solution not found!")
		}
		currentState = frontier[0]

	}
	return currentState, nil
}

func delete(s state, states []state) []state {
	i, err := index(states, s)
	if err != nil {
		log.Fatal(err)
	}
	return remove(states, i)
}

func prune(states []state) []state {
	var prunedStates []state
	for _, state := range states {
		if state.isValid() { // Here we get rid of the invalid states
			prunedStates = append(prunedStates, state)
		}
	}
	return prunedStates
}

func reorder(states []state) []state {
	// No reordering done
	return states
}

func insert(s []state, states []state) []state {
	return append(s, states...)
}

func contains(s state, states []state) bool {
	for _, v := range states {
		if v.cl == s.cl && v.ml == s.ml && v.bIsLeft == s.bIsLeft {
			return true
		}
	}
	return false
}

func index(slice []state, item state) (int, error) {
	for i := range slice {
		if slice[i] == item {
			return i, nil
		}
	}
	return -1, errors.New("State not found in slice")
}

func remove(slice []state, i int) []state {
	return append(slice[:i], slice[i+1:]...)
}
