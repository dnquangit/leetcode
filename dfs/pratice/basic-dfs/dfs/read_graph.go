package dfs

import (
	"fmt"
	"strconv"
)

func ReadGraph(filename string) *Graph {
	var g Graph
	g.AdjacencyList = make(map[int][]int, 0)

	fileLines := readFile(filename)
	g.NumVertex, _ = strconv.Atoi((*fileLines)[0])
	g.AdjacencyMatrix = make([][]int, g.NumVertex)

	nameVertex := 0
	for i := 0; i < len(*fileLines); i++ {

		if i == 0 { /* headline */
			continue
		}

		temp := make([]int, 0)
		list := make([]int, 0)
		for j := 0; j < len((*fileLines)[i]); j++ {
			intValue, _ := strconv.Atoi(string((*fileLines)[i][j]))
			temp = append(temp, intValue)
			if intValue == 1 {
				list = append(list, j)
			}
		}

		g.AdjacencyMatrix[nameVertex] = temp
		g.AdjacencyList[nameVertex] = list
		nameVertex = nameVertex + 1
	}

	fmt.Println(g.AdjacencyList)

	return &g
}
