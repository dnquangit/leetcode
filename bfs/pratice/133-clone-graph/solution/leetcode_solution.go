package solution

func cloneGraph2(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	var queue []*Node

	queue = append(queue, node)

	//create nodes
	for len(queue) != 0 {
		oldNode := queue[0]
		queue = queue[1:]

		newNode := new(Node)
		newNode.Val = oldNode.Val
		visited[oldNode] = newNode

		for _, v := range oldNode.Neighbors {
			if _, ok := visited[v]; ok {
				continue
			}

			queue = append(queue, v)
		}
	}

	//create edges
	for oldNode, newNode := range visited {
		for _, neighbor := range oldNode.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, visited[neighbor])
		}
	}

	return visited[node]
}
