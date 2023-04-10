package bfs

import "fmt"

func resolvePath(path map[int]int, start int, target int) []int {
	result := make([]int, 0)
	result = append(result, target)
	notFound := false

	for true {
		previous, ok := path[target]

		if !ok || previous == -1 {
			notFound = true
		}

		if start == previous {
			result = append(result, previous)
			break
		}
		result = append(result, previous)
		target = previous
	}

	if notFound {
		return nil
	}

	return result
}

func PrintShortestPath(path []int) {
	for i := len(path) - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%v", path[i])
		} else {
			fmt.Printf("%v -> ", path[i])
		}
	}
}
