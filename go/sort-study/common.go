package sort_study

import (
	"math/rand"
	"time"
)

type IList []int

func (list *IList) Swap(i int, j int) {
	(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
}

func InitValue(l int) IList {
	list := make(IList, l)
	for i := 0; i < l; i++ {
		list[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })

	return list
}
