package sort_study

import (
	"math/rand"
	"testing"
	"time"
)

func TestCountingSortA(t *testing.T) {
	list := make([]int, 100)
	for i := range list {
		list[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })

	k := 100
	result := CountingSortA(list, k)
	t.Log(result)
}

func TestCountingSortB(t *testing.T) {
	list := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	result := CountingSortB(list, 10)
	t.Log(result)
}

func BenchmarkCountingSortA_N10_K10(b *testing.B) {
	list := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	k := 10

	for i := 0; i < b.N; i++ {
		CountingSortA(list, k)
	}
}

func BenchmarkCountingSortB_N10_K10(b *testing.B) {
	list := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	k := 10

	for i := 0; i < b.N; i++ {
		CountingSortB(list, k)
	}
}

func BenchmarkCountingSortA_N100_k100(b *testing.B) {
	list := make([]int, 100)
	for i := range list {
		list[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })

	k := 100

	for i := 0; i < b.N; i++ {
		CountingSortA(list, k)
	}
}

func BenchmarkCountingSortB_N100_k100(b *testing.B) {
	list := make([]int, 100)
	for i := range list {
		list[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })

	k := 100

	for i := 0; i < b.N; i++ {
		CountingSortB(list, k)
	}
}

func BenchmarkCountingSortA_N1000_k1000(b *testing.B) {
	list := make([]int, 1000)
	for i := range list {
		list[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })

	k := 1000

	for i := 0; i < b.N; i++ {
		CountingSortA(list, k)
	}
}

func BenchmarkCountingSortB_N1000_k1000(b *testing.B) {
	list := make([]int, 1000)
	for i := range list {
		list[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })

	k := 1000

	for i := 0; i < b.N; i++ {
		CountingSortB(list, k)
	}
}
