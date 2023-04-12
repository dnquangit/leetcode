package solution

import "fmt"

func AddNeighbor(node *Node, neighbor *Node) {
	if node == nil || neighbor == nil {
		return
	}
	if node.Neighbors == nil {
		node.Neighbors = make([]*Node, 0)
	}
	node.Neighbors = append(node.Neighbors, &Node{
		Val:       neighbor.Val,
		Neighbors: nil,
	})
}

func AddNeighbor2(node *Node, neighbor *Node) {
	if node == nil || neighbor == nil {
		return
	}
	if node.Neighbors == nil {
		node.Neighbors = make([]*Node, 0)
	}
	//fmt.Printf("add from  %v to %v \n", node.Val, neighbor.Val)
	node.Neighbors = append(node.Neighbors, neighbor)
	for i := 0; i < len(node.Neighbors); i++ {
		//fmt.Printf("Neighbors : i: %v - v %v \n", i, node.Neighbors[i].Val)
	}
}

func Travel(node *Node) {
	queue := make([]*Node, 0)
	visited := make(map[int]bool)

	queue = append(queue, node)
	visited[node.Val] = true

	/* loop though neighborhood */
	//debug := 0
	for len(queue) > 0 {
		PrintSub(queue[0])
		firstItem := queue[0]
		queue = queue[1:]

		neighborhood := firstItem.Neighbors
		for i := 0; i < len(neighborhood); i++ {
			//fmt.Printf("neighborhood: i-%v - value: %v \n", i, neighborhood[i].Val)
			_, ok := visited[neighborhood[i].Val]
			if !ok {
				queue = append(queue, neighborhood[i])
				visited[neighborhood[i].Val] = true
				fmt.Println(neighborhood[i].Val)
			}
		}

	}
}

func PrintSub(node *Node) {
	fmt.Printf("%v- sub: ", node.Val)
	neighborhood := node.Neighbors
	for i := 0; i < len(neighborhood); i++ {
		fmt.Printf("%v ", neighborhood[i].Val)
	}
	fmt.Println()
}
