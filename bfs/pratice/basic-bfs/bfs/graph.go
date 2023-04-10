package bfs

type Graph struct {
	NumVertex       int
	AdjacencyMatrix [][]int
	AdjacencyList   map[int][]int
}
