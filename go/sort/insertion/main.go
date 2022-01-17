package main

import "fmt"

type iList []int

func (list *iList) Swap(i int, j int) {
	fmt.Printf("[%d, %d] Swap ", (*list)[i], (*list)[j])
	(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
}

func main() {
	arr := iList{4, 0, 3, 7, 2, 1, 9, 6, 5, 8}

	for i := 1; i < len(arr); i++ {
		fmt.Printf("%d회차: ", i)
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr.Swap(j, j-1)
			} else {
				break
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println(arr)
}
