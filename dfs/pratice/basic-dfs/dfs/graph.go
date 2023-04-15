package dfs

type Graph struct {
	NumVertex       int
	AdjacencyMatrix [][]int
	AdjacencyList   map[int][]int
}
