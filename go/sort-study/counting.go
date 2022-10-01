package sort_study

func CountingSortA(arr []int, k int) []int {
	count := make([]int, k+1)
	for _, v := range arr { // n 회
		count[v]++
	}

	sorted := make([]int, 0, len(arr))
	for i := 0; i < k+1; i++ {
		for j := 0; j < count[i]; j++ { // n 회
			sorted = append(sorted, i)
		}
	}

	return sorted
}

func CountingSortB(arr []int, k int) []int {
	count := make([]int, k+1)
	for _, v := range arr { // n 회
		count[v]++
	}

	for i := 1; i <= k; i++ { // k 회
		count[i] += count[i-1]
	}

	sorted := make([]int, len(arr))
	for _, v := range arr { // n 회
		sorted[count[v]-1] = v
		count[v]--
	}

	return sorted
}
