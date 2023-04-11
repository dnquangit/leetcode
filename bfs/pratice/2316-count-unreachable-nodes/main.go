package main

import (
	"fmt"
	"go-module/solution"
)

/*
	solution;
	- tân dụng bfs để chia và đếm các điểm trong graph thành các component.
	- sau khi có được số lượng node ở component, tính số pair theo xác xuất thống kê
		count = bfs(graph, start)
		pair += remain * (remain - count)
		remain -= remain - count
*/
func main() {
	//fmt.Println(solution.CountUnreachableNodes(3, [][]int{[]int{0, 1}, []int{0, 2}, []int{1, 2}}))
	fmt.Println(solution.CountUnreachableNodes(7, [][]int{[]int{0, 2}, []int{0, 5}, []int{2, 4}, []int{1, 6}, []int{5, 4}}))
}
