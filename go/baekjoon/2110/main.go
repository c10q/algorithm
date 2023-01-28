package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func NextUint64() uint64 {
	sc.Scan()
	var r uint64
	for _, c := range sc.Bytes() {
		r *= 10
		r += uint64(c - byte(48))
	}
	return r
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	N := NextUint64() // 1 ~ 200000
	M := NextUint64() // 1 ~ 1000000000

	routers := make([]uint64, N)
	for i := uint64(0); i < N; i++ {
		router := NextUint64()
		routers[i] = router
	}

	sort.Slice(routers, func(i, j int) bool {
		return routers[i] < routers[j]
	})

	min := uint64(1)
	max := routers[N-1]

	for min < max {
		mid := (min + max) / 2

		rtCnt := uint64(1)
		tpIdx := uint64(0)
		for i := uint64(1); i < N; i++ {
			if routers[i]-routers[tpIdx] >= mid {
				tpIdx = i
				rtCnt += 1
			}
		}

		if rtCnt < M {
			max = mid
		} else {
			min = mid + 1
		}

	}

	wr.WriteString(fmt.Sprintf("%d", min-1))
}
