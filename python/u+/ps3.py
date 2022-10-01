from collections import deque


def is_okay(train):
    dx = [0, 0, 1, -1]
    dy = [1, -1, 0, 0]
    m = {
        'E': [0, 1],
        'W': [0, -1],
        'S': [1, 0],
        'N': [-1, 0]
    }
    for i in range(len(train)):
        for j in range(len(train[i])):
            for k in range(4):
                nx = i + dx[k]
                ny = j + dy[k]
                if 0 <= nx < len(train) and 0 <= ny < len(train[i]):
                    point_you = m[train[nx][ny]]
                    point_me = m[train[i][j]]

                    if [i, j] == [point_you[0] + nx, point_you[1] + ny] and [nx, ny] == [point_me[0] + i,
                                                                                         point_me[1] + j]:

                        return False, [[i, j], [nx, ny]]

    return True, []


def solution(train):
    queue = deque()

    queue.append([train, 0, []])

    while queue:
        now = queue.popleft()

        t = now[0]
        count = now[1]
        visited = now[2]
        r = is_okay(t)
        if r[0]:
            return count
        errors = r[1]
        for e in errors:
            i, j = e
            for k in ['E', 'W', 'S', 'N']:
                if [i, j, k] not in visited and t[i][j] != k:
                    temp = [x[:] for x in t]
                    temp[i] = temp[i][:j] + k + temp[i][j + 1:]
                    queue.append([temp, count + 1, visited + [[i, j, k]]])


if __name__ == '__main__':
    print(is_okay(["ESS", "EEW", "NNW"]))
    print(is_okay(["ESS", "EEE", "NNW"]))
    print(solution(["ESS", "EEW", "NNW"]) == 1)
    print(solution(["EW", "EW", "EW"]) == 3)
    print(solution(["NSN", "ESW", "ENW", "NNN"]) == 2)
