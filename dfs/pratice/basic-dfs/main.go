package main

import (
	"go-module/dfs"
)

func main() {
	g := dfs.ReadGraph("graph.txt")
	path := g.TravelDFS(0, 5)
	dfs.PrintShortestPath(path)
}
