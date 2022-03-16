package main

func main() {
	start := genState(3, 3, left)
	solution := bfs(*start)
	p := solution.getPath()
	p.printPath()
	p.printDepth()
}
