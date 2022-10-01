from collections import deque


def solution(players):
    max_value = 0
    for i in range(len(players)):
        if players[i] != 1:
            r = tournament(players.copy(), i)

            max_value = max(max_value, r)

    return max_value


def tournament(players, pro):
    players[pro] = 1
    queue = deque(players)
    special_matches = 0

    while queue:
        p1 = queue.popleft()
        if not queue:
            break
        p2 = queue.popleft()

        if p1 == 1 or p2 == 1:
            special_matches += 1
            queue.append(1)
        else:
            queue.append(0)

    return special_matches


if __name__ == '__main__':
    print(solution([1, 0, 0, 1, 0, 0, 1, 0]) == 7)
    print(solution([0, 0, 0, 1]) == 3)
    print(solution([1, 0, 1, 0, 1, 0, 1, 0]) == 7)
