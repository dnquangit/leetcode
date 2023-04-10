package main

import "go-module/bfs"

func main() {
	graph := bfs.ReadGraph("graph.txt")
	shortestPath := graph.FindShortestPath(0, 5)
	if shortestPath != nil {
		bfs.PrintShortestPath(shortestPath)
	}
}
