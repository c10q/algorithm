import sys
import heapq

n = int(sys.stdin.readline())

leftHeap = []
rightHeap = []
for i in range(n):
    num = int(sys.stdin.readline())

    if len(leftHeap) == len(rightHeap):
        heapq.heappush(leftHeap, -num)
    else:
        heapq.heappush(rightHeap, num)

    if rightHeap and rightHeap[0] < -leftHeap[0]:
        left = heapq.heappop(leftHeap)
        right = heapq.heappop(rightHeap)

        heapq.heappush(leftHeap, -right)
        heapq.heappush(rightHeap, -left)

    sys.stdout.write(str(-leftHeap[0]) + "\n")
