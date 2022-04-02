package main

import "fmt"

func solution1(dist [][]int) [][]int {
	fmt.Println(dist)
	return [][]int{{}}
}

func solution2(grid []string) int {
	var block []string
	for i, str := range grid {
		for j, char := range str {
			if string(char) == "?" {
				if i-1 > 0 {
					block = append(block, string(str[i-1]))
				}
				if i+1 < len(grid) {
					block = append(block, string(str[i+1]))
				}
				if j-1 > 0 {
					block = append(block, string(str[j-1]))
				}
				if j+1 < len(str) {
					block = append(block, string(str[j+1]))
				}
			}
		}
	}
	fmt.Println(block)

	return -1
}

func progress(maps [][]int, start int, end int, limit int, result *[]int, stack []int) {
	if len(stack) >= limit {
		for _, s := range stack {
			if s == end {
				for _, v := range stack {
					bf := false
					for _, v1 := range *result {
						if v == v1 {
							bf = true
							break
						}
					}
					if bf {
						continue
					}
					*result = append(*result, v)
				}
				break
			}
		}
		return
	}

	for _, node := range maps[start] {
		newStack := make([]int, len(stack)+1)
		for i, v := range stack {
			newStack[i] = v
		}
		newStack[len(stack)] = node
		progress(maps, node, end, limit, result, newStack)
	}
}

func solution3(n int, edges [][]int, k int, a int, b int) int {
	maps := make([][]int, n)
	var resultArr []int
	var result = 0

	for _, edge := range edges {
		e0 := edge[0]
		e1 := edge[1]
		maps[e0] = append(maps[e0], e1)
		maps[e1] = append(maps[e1], e0)
	}
	progress(maps, a, b, k, &resultArr, []int{})
	resultArr = append(resultArr, a)
	resultArr = append(resultArr, b)

	for _, edge := range edges {
		e0 := edge[0]
		e1 := edge[1]
		b0 := false
		b1 := false
		for _, node := range resultArr {
			if e0 == node {
				b0 = true
			}
			if e1 == node {
				b1 = true
			}
		}
		if b0 && b1 {
			result += 1
		}
	}

	return result
}

func main() {
	//var dist [][]int
	//dist = [][]int{{0, 5, 2, 4, 1}, {5, 0, 3, 9, 6}, {2, 3, 0, 6, 3}, {4, 9, 6, 0, 3}, {1, 6, 3, 3, 0}}

	//solution1(dist)

	// var grids [][]string
	// grids = [][]string{{"??b", "abc", "cc?"}, {"abcabcab", "????????"}, {"aa?"}}

	// for _, grid := range grids {
	// solution2(grid)
	// }

	var edges [][]int
	edges = [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 0}, {5, 1}, {6, 1}, {7, 2}, {7, 3}, {4, 5}, {5, 6}, {6, 7}}

	fmt.Println(solution3(8, edges, 4, 0, 3))
}
