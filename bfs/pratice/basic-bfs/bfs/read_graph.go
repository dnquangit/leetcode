package bfs

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFile(filename string) *[]string {
	readFile, err := os.Open(filename)

	if err != nil {
		return nil
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return &fileLines
}

/*
	graph.txt
	6
	010100
	101101
	010001
	110011
	000100
	011100
*/

func ReadGraph(filename string) *Graph {
	var g Graph
	g.AdjacencyList = make(map[int][]int, 0)

	fileLines := ReadFile(filename)
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

	return &g
}
