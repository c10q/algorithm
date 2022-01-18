package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rcv(p [2]int, w int, board [][]int) (int, int) {
	pX := p[0]
	pY := p[1]

	blue := 0
	white := 0

	for i := pX; i < pX+w; i++ {
		for j := pY; j < pY+w; j++ {
			if board[i][j] == 0 {
				white += 1
			} else {
				blue += 1
			}
		}
	}

	if blue == w*w || white == w*w {
		if board[pX][pY] == 0 {
			return 1, 0
		} else {
			return 0, 1
		}
	} else {
		h := w / 2

		w1, b1 := rcv([2]int{pX, pY}, h, board)
		w2, b2 := rcv([2]int{pX + h, pY}, h, board)
		w3, b3 := rcv([2]int{pX, pY + h}, h, board)
		w4, b4 := rcv([2]int{pX + h, pY + h}, h, board)

		return w1 + w2 + w3 + w4, b1 + b2 + b3 + b4
	}

}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n int
	fmt.Fscanln(rd, &n)

	board := make([][]int, n)

	for i := 0; i < n; i++ {
		readString, err := rd.ReadString('\n')
		rs := strings.TrimSuffix(readString, "\n")
		if err != nil {
			return
		}

		tempSlice := make([]int, n)

		for j, s := range strings.Split(rs, " ") {
			tempSlice[j], _ = strconv.Atoi(s)
		}

		board[i] = tempSlice
	}

	w, b := rcv([2]int{0, 0}, n, board)
	wr.WriteString(strconv.Itoa(w) + "\n")
	wr.WriteString(strconv.Itoa(b))
}
