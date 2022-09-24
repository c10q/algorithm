# 11111 1 11111 1 11111 1 11111 1 11111 1 11111 1 11111 1 11111
# 1 3 7 15 23 31 47

def solution(numbers):
    answer = []

    for i in numbers:
        b = binary(i)
        print(b)
        h = get_binary_count_by_height(i)
        nodes = h[0]
        height = h[1]
        for i in range(len(b), nodes):
            b = '0' + b
        print(b)
        if b[(nodes // 2)] == '0':
            answer.append(0)
            continue

        if check_has_node((nodes // 2) + 1, b, height - 1):
            answer.append(1)
        else:
            answer.append(0)

    return answer


def get_binary_count_by_height(n):
    b = len(binary(n))
    x = 1
    nodes = 1
    while True:
        if b <= nodes:
            return [nodes, x]
        nodes += 2 ** x
        x += 1


def binary(n):
    return str(bin(n)[2:])


def check_has_node(pivot, b, height):
    if height == 1:
        return True

    print(pivot, b, height)

    middle = ((len(b) - pivot) + 1) // 2
    if b[middle - 1] == '0':
        return False
    else:
        left = check_has_node(middle, b, height - 1)

    if b[pivot + middle - 1] == '0':
        return False
    else:
        right = check_has_node(pivot + middle, b, height - 1)

    return left and right


def bin_to_int(b):
    return int(b, 2)


if __name__ == '__main__':
    print(solution([100000000000000, 5]) == [1, 0])
    # print(solution([63, 111, 95]) == [1, 1, 0])
    print(bin_to_int('0101111101111111111111111111111'))
    print(check_has_node(16, '0101111101111111111111111111111', 4))

