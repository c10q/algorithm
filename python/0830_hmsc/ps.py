def solution(arrA, arrB):
    answer = False
    for i in range(len(arrA)):
        if arrA[i:] + arrA[:i] == arrB:
            return True

    return answer


print(solution([1, 2, 3, 4], [1, 2, 3, 4]))
print(solution([7, 8, 10], [10, 7, 8]))
print(solution([7, 8, 10], [8, 10, 7]))