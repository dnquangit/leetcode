package solution

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	/*
		the main point: basically, using bfs for searching from source and source's neighborhoods
	*/

	/* read graph, numVertex, adjacencyList */

	/*
		use queue to hold the processed node
		use visited to hold the visited node

		searching from source node, and move to neighborhood nodes and go on, go on ... when queue is not empty
		if found target, break loop and return result.
		otherwise, return false at outside loop.
	*/
	graph := readGraph(n, edges)
	queue := make([]int, 0)
	visited := make(map[int]bool)

	/* mark start */
	queue = append(queue, source)
	visited[source] = true

	/* loop though neighborhood */
	//debug := 0
	for len(queue) > 0 {

		firstItem := queue[0]
		if firstItem == destination {
			return true
		}
		queue = queue[1:]

		neighborhood := graph.AdjacencyList[firstItem]
		for i := 0; i < len(neighborhood); i++ {
			isVisited, ok := visited[neighborhood[i]]
			if !ok || !isVisited {
				queue = append(queue, neighborhood[i])
				visited[neighborhood[i]] = true
			}
		}

	}

	return false
}
