package sort_study

func CountingSort(slice []int, max int) []int {
	count := make([]int, max+1)
	for _, v := range slice {
		count[v]++
	}
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}
	result := make([]int, len(slice))
	for _, v := range slice {
		result[count[v]-1] = v
		count[v]--
	}
	return result
}
