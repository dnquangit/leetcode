package solution

/*
	n = 3, edges = [[0,1],[1,2],[2,0]]
*/
func readGraph(n int, edges [][]int) *Graph {
	var g Graph

	g.NumVertex = n
	g.AdjacencyList = make(map[int][]int, n)

	for i := 0; i < len(edges); i++ {
		if len(edges[i]) < 2 {
			continue
		}
		label1 := edges[i][0]
		label2 := edges[i][1]

		addAdjacencyVertex(g.AdjacencyList, label1, label2)
		addAdjacencyVertex(g.AdjacencyList, label2, label1)
	}

	return &g
}

func addAdjacencyVertex(list map[int][]int, label int, labelLinked int) {
	currentList, ok := list[label]
	if ok {
		currentList = append(currentList, labelLinked)
		list[label] = currentList
	} else {
		list[label] = []int{labelLinked}
	}
}
