package sort_study

func BubbleSort(arr IList) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr.Swap(j, j+1)
			}
		}
	}
}
