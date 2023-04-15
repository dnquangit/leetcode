package dfs

func (g *Graph) TravelDFS(start int, target int) []int {
	if start == target {
		return nil
	}

	stack := make([]int, 0)
	visited := make(map[int]bool, g.NumVertex)
	path := make(map[int]int, g.NumVertex)

	// first
	stack = append(stack, start)
	visited[start] = true

	for len(stack) > 0 {

		// lastItem
		lastItem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		neighborhood := g.AdjacencyList[lastItem]
		for i := 0; i < len(neighborhood); i++ {
			if visited[neighborhood[i]] == false {
				visited[neighborhood[i]] = true
				stack = append(stack, neighborhood[i])
				path[neighborhood[i]] = lastItem
			}
		}
	}

	result := resolvePath(path, start, target)
	return result
}
