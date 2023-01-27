package scan

import (
	"math"
	"strconv"
)

func bytesToIntV1(b []byte) int {
	r := 0
	for i := range b {
		r += int(b[len(b)-i-1]-byte(48)) * int(math.Pow10(i))
	}
	return r
}

func bytesToIntV2(b []byte) int {
	i, _ := strconv.Atoi(string(b))
	return i
}

func bytesToIntV3(b []byte) int {
	r := 0
	for _, c := range b {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func bytesToIntV4(b []byte) int {
	r := 0
	for _, c := range b {
		r *= 10
		r += int(c - byte(48))
	}
	return r
}
