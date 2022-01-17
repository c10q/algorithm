package main

import "fmt"

type iList []int

func (list iList) Min() (int, int) {
	var idx int
	var min int

	for i, v := range list {
		if i == 0 {
			idx = i
			min = v
		}
		if v < min {
			idx = i
			min = v
		}
	}

	fmt.Printf("[%d] Min ", min)

	return idx, min
}

func (list *iList) Swap(i int, j int) {
	fmt.Printf("[%d, %d] Swap ", (*list)[i], (*list)[j])
	(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
}

func main() {
	arr := iList{4, 0, 3, 7, 2, 1, 9, 6, 5, 8}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d회차: ", i)
		idx, _ := arr[i:].Min()
		arr.Swap(i, idx+i)
		fmt.Printf("\n")
	}

	fmt.Println(arr)
}
