package solution

import "fmt"

func CountUnreachableNodes(n int, edges [][]int) int64 {
	pairs := countPairs(n, edges)
	return pairs
}

func countPairs(n int, edges [][]int) int64 {

	// build graph
	graph := readGraph(n, edges)

	// queue, visited
	queue := make([]int, 0)
	visited := make(map[int]bool, 0)
	var pairs int64
	remains := graph.NumVertex

	for i := 0; i < graph.NumVertex; i++ {

		labelVertex := i
		pairInComponent := 0

		if isAlreadyVisitedNode(visited, labelVertex) == true {
			continue
		}
		// bfs search
		queue = append(queue, labelVertex)
		visited[labelVertex] = true
		pairInComponent++

		for len(queue) > 0 {
			firstItem := queue[0]
			queue = queue[1:]

			neighborhood, ok := graph.AdjacencyList[firstItem]
			if ok {
				for j := 0; j < len(neighborhood); j++ {
					if isAlreadyVisitedNode(visited, neighborhood[j]) == false {
						visited[neighborhood[j]] = true
						pairInComponent++
						queue = append(queue, neighborhood[j])
					}
				}
			}
		}

		// pairInComponent ??
		//fmt.Printf("i: %v - node: %v - pairInComponent: %v - visited :%v \n", i, labelVertex, pairInComponent, isAlreadyVisitedNode(visited, labelVertex))
		pairs += int64(pairInComponent * (remains - pairInComponent))
		remains -= pairInComponent
		fmt.Printf("i: %v - node: %v - pairInComponent: %v - pairs: %v - remains: %v \n", i, labelVertex, pairInComponent, pairs, remains)
	}

	return pairs
}

func isAlreadyVisitedNode(list map[int]bool, key int) bool {
	_, found := list[key]
	return found
}
