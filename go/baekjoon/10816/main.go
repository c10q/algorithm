package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	arr := make([]int, 0, 500000)
	temp := make([]int, 20000001)

	var n int
	fmt.Fscanln(rd, &n)

	readString, _ := rd.ReadString('\n')
	rs := strings.TrimSuffix(readString, "\n")

	for _, s := range strings.Split(rs, " ") {
		a, _ := strconv.Atoi(s)
		arr = append(arr, a)
		temp[a+10000000] += 1
	}

	sort.Ints(arr)

	var m int
	fmt.Fscanln(rd, &m)

	readString, _ = rd.ReadString('\n')
	rs = strings.TrimSuffix(readString, "\n")

	for _, s := range strings.Split(rs, " ") {
		a, _ := strconv.Atoi(s)

		start := 0
		end := n - 1

		r := "0\n"
		for start <= end {
			v := (start + end) / 2
			if (start+end)%2 != 0 {
				v += 1
			}
			if arr[v] == a {
				r = strconv.Itoa(temp[a+10000000]) + "\n"
				break
			} else {
				if arr[v] < a {
					start = v + 1
				} else {
					end = v - 1
				}
			}
		}
		wr.WriteString(r)
	}

}
