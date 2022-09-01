package ps2

import "fmt"

func solution(pouches []string) int {
	l := len(pouches)
	fmt.Println(l)
	p := initPouches(pouches)
	return l - dfs([5]int{0, 0, 0, 0, 0}, p, l)
}

func dfs(pouch [5]int, pouches [][5]int, min int) int {
	if len(pouches) == 0 {
		return 0
	}

	for i := range pouches {
		newPouch := addPouch(pouch, pouches[i])
		if isPassable(newPouch) {
			min = minInt(min, dfs(newPouch, copiedRemovedPouch(pouches, i), len(pouches)))
		} else {
			dfs(newPouch, copiedRemovedPouch(pouches, i), min)
		}
	}

	fmt.Println(min)

	return min
}

func initPouches(pouches []string) [][5]int {
	var result [][5]int
	for _, pouch := range pouches {
		var temp [5]int
		for _, v := range pouch {
			if v == 'a' {
				temp[0]++
			}
			if v == 'b' {
				temp[1]++
			}
			if v == 'c' {
				temp[2]++
			}
			if v == 'd' {
				temp[3]++
			}
			if v == 'e' {
				temp[4]++
			}
		}
		result = append(result, temp)
	}
	return result
}

func addPouch(pouch [5]int, otherPouch [5]int) [5]int {
	var result [5]int
	for i := range pouch {
		result[i] = pouch[i] + otherPouch[i]
	}
	return result
}

func isPassable(pouch [5]int) bool {
	idx, val := getMax(pouch)
	var sum int
	for i := range pouch {
		if i == idx {
			continue
		}
		sum += pouch[i]
	}

	if sum < val {
		return true
	}

	return false
}

func getMax(pouch [5]int) (idx, val int) {
	for i, v := range pouch {
		if v > val {
			idx = i
			val = v
		}
	}
	return idx, val
}

func copiedRemovedPouch(pouches [][5]int, idx int) [][5]int {
	var result [][5]int
	for i := range pouches {
		if i == idx {
			continue
		}
		result = append(result, pouches[i])
	}
	return result
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
