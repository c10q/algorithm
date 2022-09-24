from collections import deque


def solution(cap, n, deliveries, pickups):
    answer = 0
    q = deque(deliveries)
    p = deque(pickups)
    go_back = False
    temp_q = cap
    distance = n
    while q:
        if not go_back:
            temp_q -= q[-1]
            if temp_q < 0:
                answer += distance * 2
                distance = len(q)
                go_back = True
                continue
            temp_q -= p[-1]
            print('-', q.pop())
            print('+', p.pop())
        else:
            go_back = False
            print('goback')
            temp_q = cap
    print(answer)
    return answer


if __name__ == '__main__':
    print(solution(4, 5, [1, 0, 3, 1, 2], [0, 3, 0, 4, 0]) == 16)
    print(solution(2, 7, [1, 0, 2, 0, 1, 0, 2], [0, 2, 0, 1, 0, 2, 0]) == 30)
