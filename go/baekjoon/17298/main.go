package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []int

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var nInt int
	fmt.Fscanln(rd, &nInt)

	s, _ := rd.ReadString('\n')
	ns := strings.TrimSuffix(s, "\n")
	nums := strings.Split(ns, " ")
	numSlice := make([]int, nInt)

	for i, num := range nums {
		numSlice[i], _ = strconv.Atoi(num)
	}

	result := make([]string, nInt)
	for i := 0; i < len(result); i++ {
		result[i] = "-1"
	}

	temp := make(stack, 0, nInt)

	for i := 0; i < nInt; i++ {
		v := numSlice[i]
		for len(temp) > 0 && numSlice[temp[len(temp)-1]] < v {
			p := temp.pop()
			result[p] = strconv.Itoa(v)
		}

		temp.push(i)
	}

	wr.WriteString(strings.Join(result, " "))
}

func (s *stack) push(i int) {
	*s = append(*s, i)
}

func (s *stack) pop() int {
	lastIdx := len(*s) - 1
	p := (*s)[lastIdx]
	*s = (*s)[:lastIdx]
	return p
}
