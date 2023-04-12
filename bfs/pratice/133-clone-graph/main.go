package main

import (
	"go-module/solution"
)

func main() {
	var node1 solution.Node
	node1.Val = 1

	var node2 solution.Node
	node2.Val = 2

	var node3 solution.Node
	node3.Val = 3

	var node4 solution.Node
	node4.Val = 4

	solution.AddNeighbor2(&node1, &node2)
	solution.AddNeighbor2(&node1, &node4)

	solution.AddNeighbor2(&node2, &node1)
	solution.AddNeighbor2(&node2, &node3)

	solution.AddNeighbor2(&node3, &node2)
	solution.AddNeighbor2(&node3, &node4)

	solution.AddNeighbor2(&node4, &node1)
	solution.AddNeighbor2(&node4, &node3)
	//solution.Travel(&node1)
	//fmt.Println("clone graph", solution.Clone(&node1))
	clone := solution.Clone(&node1)
	solution.Travel(clone)
}
