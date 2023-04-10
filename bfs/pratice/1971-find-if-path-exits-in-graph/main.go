package main

import (
	"fmt"
	"go-module/solution"
)

/*
	Topic:
	There is a bi-directional graph with n vertices, where each vertex is labeled from 0 to n - 1 (inclusive).
	The edges in the graph are represented as a 2D integer array edges, where each edges[i] = [ui, vi] denotes a bi-directional edge between vertex ui and vertex vi.
	Every vertex pair is connected by at most one edge, and no vertex has an edge to itself.

	You want to determine if there is a valid path that exists from vertex source to vertex destination.

	Given edges and the integers n, source, and destination,
	return true if there is a valid path from source to destination, or false otherwise.

	Solution template:
	func validPath(n int, edges [][]int, source int, destination int) bool {

	}

	i.e
	topic.png
	n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
	Output: true
	Explanation: There are two paths from vertex 0 to vertex 2:
	- 0 → 1 → 2
	- 0 → 2

	topic2.png
	Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
	Output: false
	Explanation: There is no path from vertex 0 to vertex 5.

*/
func main() {
	fmt.Println(solution.ValidPath(3, [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 0}}, 0, 2))
	fmt.Println(solution.ValidPath(6, [][]int{[]int{0, 1}, []int{0, 2}, []int{3, 5}, []int{5, 4}, []int{4, 3}}, 0, 5))

}
