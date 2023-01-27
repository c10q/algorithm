package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NextLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 500000), 500000)

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var temp [100002]int

	firstLine := strings.Split(NextLine(sc), " ")

	N, _ := strconv.Atoi(firstLine[0])
	M, _ := strconv.Atoi(firstLine[1])

	secondLine := strings.Split(NextLine(sc), " ")
	temp[1], _ = strconv.Atoi(secondLine[0])
	for i := 1; i < N; i++ {
		v, _ := strconv.Atoi(secondLine[i])
		temp[i+1] = temp[i] + v
	}

	for s := 0; s < M; s++ {
		section := strings.Split(NextLine(sc), " ")
		i, _ := strconv.Atoi(section[0])
		j, _ := strconv.Atoi(section[1])

		wr.WriteString(fmt.Sprintf("%d\n", temp[j]-temp[i-1]))
	}
}
