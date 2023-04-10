package bfs

func (g *Graph) FindShortestPath(start int, target int) []int {
	if start == target {
		return nil
	}

	queue := make([]int, 0)
	visited := make(map[int]bool, g.NumVertex)
	path := make(map[int]int, g.NumVertex)

	queue = append(queue, start)
	visited[start] = true
	for len(queue) != 0 {

		/* first item and remove */
		firstItem := queue[0]
		queue = queue[1:]

		/* loop through neighborhood */
		neighborhood := g.AdjacencyList[firstItem]
		for i := 0; i < len(neighborhood); i++ {
			if visited[neighborhood[i]] == false {
				visited[neighborhood[i]] = true
				queue = append(queue, neighborhood[i])
				path[neighborhood[i]] = firstItem
			}
		}
	}

	result := resolvePath(path, start, target)
	return result
}
