package go_test

import (
	sortStudy "go/go/sort-study"
	"testing"
)

var list = sortStudy.InitValue(20000)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortStudy.BubbleSort(list)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortStudy.SelectionSort(list)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortStudy.InsertionSort(list)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortStudy.QuickSort(list)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortStudy.MergeSort(list)
	}
}

func BenchmarkGolangSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortStudy.GolangSort(list)
	}
}
