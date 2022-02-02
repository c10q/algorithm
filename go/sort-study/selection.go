package sort_study

func (list IList) Min() (int, int) {
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

	return idx, min
}

func SelectionSort(arr IList) {
	for i := 0; i < len(arr); i++ {
		idx, _ := arr[i:].Min()
		arr.Swap(i, idx+i)
	}
}
