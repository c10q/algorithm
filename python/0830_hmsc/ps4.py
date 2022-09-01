def solution(n):
    answer = 0
    c = bin(n)[2:]
    l = c.count('1')
    return len(c) * l

print(solution(100000000))