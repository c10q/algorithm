package hacker_rank

import "sort"

func main() {
	arr := []int32{1, 2, 3, 4, 8, 5, 7}
	sort.Slice(arr, func(i, j int) bool {
		return false
	})
}
