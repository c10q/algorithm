package main

import (
	"fmt"
)

type iList []int

var cnt = 0

func (list *iList) Merge() {
	cnt += 1
	l := len(*list)
	r := (l / 2) + (l % 2)

	left := (*list)[0:r]
	right := (*list)[r:l]

	fmt.Printf("left: %d, right: %d\n", left, right)

	if l > 2 {
		left.Merge()
		right.Merge()
	}

	temp := make([]int, l)

	li := 0
	ri := 0

	for i := 0; i < l; i++ {
		if li >= len(left) {
			temp[i] = right[ri]
			ri += 1
		} else if ri >= len(right) {
			temp[i] = left[li]
			li += 1
		} else {
			if left[li] <= right[ri] {
				temp[i] = left[li]
				li += 1
			} else {
				temp[i] = right[ri]
				ri += 1
			}
		}
	}

	fmt.Printf("merged: %d\n", temp)

	*list = temp
}

func main() {
	arr := iList{4, 0, 3, 10, 7, 2, 1, 9, 6, 5, 8}

	arr.Merge()

	fmt.Println(arr)
}
