package sort_study

func InsertionSort(arr IList) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}
