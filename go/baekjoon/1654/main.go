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

	K := NextUint64() // 1 ~ 10000
	N := NextUint64() // 1 ~ 1000000

	lines := make([]uint64, K)
	for i := uint64(0); i < K; i++ {
		lines[i] = NextUint64()
	}

	min := uint64(0)
	max := uint64(1) << 32

	for {
		half := (min + max) / 2

		if min > max {
			wr.WriteString(fmt.Sprintf("%d", max))
			break
		}

		var cableCount uint64
		for _, l := range lines {
			cableCount += l / half
		}

		if cableCount < N {
			max = half - 1
		} else {
			min = half + 1
		}
	}

}
