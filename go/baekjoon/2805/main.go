package main

import (
	"bufio"
	"fmt"
	"os"
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

	N := NextUint64() // 1 ~ 1000000
	M := NextUint64() // 1 ~ 2000000000

	var min uint64
	var max uint64

	trees := make([]uint64, N)
	for i := uint64(0); i < N; i++ {
		trees[i] = NextUint64()
		if trees[i] > max {
			max = trees[i]
		}
	}

	for min+1 < max {
		mid := (min + max) / 2

		height := uint64(0)
		for i := range trees {
			if trees[i] > mid {
				height += trees[i] - mid
			}
		}

		if height < M {
			max = mid
		} else {
			min = mid
		}
	}

	wr.WriteString(fmt.Sprintf("%d", min))
}
