def solution(path):
    answer = []
    route = [[]]
    route_turn = [path[0]]
    d_map = {
        "E": {
            "S": "right",
            "N": "left"
        },
        "W": {
            "N": "right",
            "S": "left"
        },
        "S": {
            "W": "right",
            "E": "left"
        },
        "N": {
            "E": "right",
            "W": "left"
        },
    }

    for i in range(len(path) - 1):
        route[-1].append(path[i])
        if path[i] != path[i + 1]:
            route_turn.append(path[i + 1])
            route.append([])

    time = 0
    for r in range(len(route_turn) - 1):
        distance = len(route[r])
        now = route_turn[r]
        nxt = route_turn[r + 1]
        if distance <= 5:
            answer.append(f"Time {time}: Go straight {distance}00m and turn {d_map[now][nxt]}")

        else:
            answer.append(f"Time {time + distance - 5}: Go straight 500m and turn {d_map[now][nxt]}")
        time += distance

    return answer


print(solution("EEESEEEEEENNNN"))
print(solution("SSSSSSWWWNNNNNN"))
