package ps2

import (
	"sort"
)

func solution(capacity int, routes [][]int) bool {
	sort.Slice(routes, func(i, j int) bool {
		return routes[i][1] < routes[j][1]
	})

	deliveryList := make(map[int]int)
	var currentCapacity int

	for i := range routes {
		weight := routes[i][0]
		from := routes[i][1]
		to := routes[i][2]

		deliveryList[to] += weight
		currentCapacity -= deleteDeliveryListHasBeforeRoute(deliveryList, from)
		currentCapacity += weight

		if currentCapacity > capacity {
			return false
		}
	}

	return true
}

func deleteDeliveryListHasBeforeRoute(deliveryList map[int]int, now int) int {
	var deleteCnt int
	for k, v := range deliveryList {
		if k <= now {
			deleteCnt += v
			delete(deliveryList, k)
		}
	}
	return deleteCnt
}
