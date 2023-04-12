package solution

func Clone(node *Node) *Node {
	result := cloneGraph2(node)
	return result
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var cloneNode Node
	cloneNode.Val = node.Val

	queue := make([]*Node, 0)
	visited := make(map[int]bool)
	mapOldClone := make(map[int]*Node)

	queue = append(queue, node)
	visited[node.Val] = true
	mapOldClone[node.Val] = &cloneNode

	for len(queue) > 0 {

		firstItem := queue[0]
		queue = queue[1:]
		neighbors := firstItem.Neighbors

		for i := 0; i < len(neighbors); i++ {

			_, found := visited[neighbors[i].Val]
			if found == false {
				queue = append(queue, neighbors[i])
				visited[neighbors[i].Val] = true
				mapOldClone[neighbors[i].Val] = &Node{Val: neighbors[i].Val}
			}

			mapOldClone[firstItem.Val].Neighbors = append(
				mapOldClone[firstItem.Val].Neighbors,
				mapOldClone[neighbors[i].Val],
			)
		}

	}

	return &cloneNode
}
