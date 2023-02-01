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
	defer wr.Flush()

	N := NextUint64()
	K := NextUint64()

	min := uint64(1)
	max := K

	for min < max {
		mid := (min + max) / 2

		c := uint64(0)
		for i := uint64(1); i <= N; i++ {
			if mid/i >= N {
				c += N
			} else {
				c += mid / i
			}
		}

		if K <= c {
			max = mid
		} else {
			min = mid + 1
		}
	}

	wr.WriteString(fmt.Sprintf("%d", min))
}
