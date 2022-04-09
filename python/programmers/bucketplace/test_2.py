from collections import deque


def solution(call):
    answer = ""
    a_map = [[] for _ in range(26)]
    a_cnt = [0 for _ in range(26)]
    l_call = len(call)
    for i in range(l_call):
        char = call[i]
        if char.islower():
            a_idx = int(ord(char.upper())) - 65
        else:
            a_idx = int(ord(char)) - 65
        a_map[a_idx].append(i)
        a_cnt[a_idx] += 1

    to_delete = []
    to_delete_str = []

    for alphas in a_map:
        if len(alphas) != max(a_cnt):
            continue

        target_idx = a_map.index(alphas)
        target_group = alphas

        if call[target_idx] in to_delete:
            continue

        left_available = True
        right_available = True

        if 0 in target_group:
            left_available = False
        if 25 in target_group:
            right_available = False

        min_diff = 25
        for i in range(len(target_group) - 1):
            if target_group[i + 1] - target_group[i] < min_diff:
                min_diff = target_group[i + 1] - target_group[i]

        left = 0
        right = 0
        while left_available or right_available:
            leftAlpha = -1
            rightAlpha = -1

            for i in target_group:
                if left_available and i - 1 - left >= 0:
                    if leftAlpha == -1:
                        leftAlpha = call[i - 1 - left]
                    elif leftAlpha != call[i - 1 - left]:
                        left_available = False
                else:
                    left_available = False

                if right_available and i + 1 + right < l_call:
                    if rightAlpha == -1:
                        rightAlpha = call[i + 1 + right]
                    elif rightAlpha != call[i + 1 + right]:
                        right_available = False
                else:
                    right_available = False

            if left_available:
                left += 1
            if right_available:
                right += 1

        delete_str = call[target_idx]
        to_delete.append(call[target_idx])
        for i in range(1, left + 1):
            to_delete.append(call[target_idx - i])
            delete_str += call[target_idx - i]
        for i in range(1, right + 1):
            to_delete.append(call[target_idx + i])
            delete_str += call[target_idx + i]
        to_delete_str.append(delete_str)

    temp = [i for i in range(l_call)]
    for _str in to_delete_str:
        queue = deque(_str)
        for i in range(l_call):
            char = call[i]
            if queue.popleft().lower() != char.lower():
                queue = deque(_str)
                continue
            if len(queue) == 0:
                for x in range(len(_str)):
                    temp.remove(i - x)
                queue = deque(_str)

    for i in temp:
        answer += call[i]

    return answer


print(solution("abcabcdefabc"))
print(solution("abxdeydeabz"))
print(solution("abcabca"))
print(solution("ABCabcA"))
