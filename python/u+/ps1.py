def counting_sort(string):
    count = [0] * 10
    for s in string:
        count[int(s)] += 1

    return ''.join([str(i) * count[i] for i in range(10)])


def solution(arr):
    a_map = {}

    for i in arr:
        a = counting_sort(str(i))
        if a not in a_map:
            a_map[a] = 1
        else:
            a_map[a] += 1

    return len(a_map)


if __name__ == '__main__':
    print(solution([112, 1814, 121, 1481, 1184]) == 2)
    print(solution([123, 456, 789, 321, 654, 987]) == 3)
    print(solution([1, 2, 3, 1, 2, 3, 4]) == 4)
    print(solution([123, 234, 213, 432, 234, 1234, 2341, 12345, 324]) == 4)
