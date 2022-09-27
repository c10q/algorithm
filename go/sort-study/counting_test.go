package sort_study

import "testing"

func TestCountingSort(t *testing.T) {
	list := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	result := CountingSort(list, 10)
	t.Log(result)
}
