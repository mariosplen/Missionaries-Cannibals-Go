package main

func bfs(initialState state) state {
	if initialState.isGoal() {
		return initialState
	}

	var closed stateSlice
	var frontier stateSlice
	frontier = append(frontier, initialState)
	for len(frontier) != 0 {
		s := frontier[0]        // Get top element
		frontier = frontier[1:] // Discard top element
		if !closed.contains(s) {
			closed = append(closed, s)
		}
		neighbours := s.spawnNeighbours()
		for _, neighbour := range neighbours {
			if !closed.contains(neighbour) || !frontier.contains(neighbour) {
				frontier = append(frontier, neighbour)
				if neighbour.isGoal() {
					return neighbour
				}
			}
		}
	}
	return initialState
}
