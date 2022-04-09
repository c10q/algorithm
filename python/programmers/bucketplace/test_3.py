from collections import deque


def solution(tstring, variables):
    answer = tstring
    m = {}
    for var in variables:
        k = var[0]
        v = var[1]
        s = v.replace('{', '')
        s = s.replace('}', '')
        if s in m and m[s] == '{' + k + '}':
            variables.remove([s, '{' + k + '}'])
            continue
        if v[0] == '{' and v[-1] == '}':
            m[k] = s

    remove_var = []

    for k in m:
        c = k
        n = m[c]
        queue = deque([[c, n]])
        visited = [[c, n]]
        while queue:
            x1, x2 = queue.popleft()
            if x2 in m:
                if x2.find(k):
                    for v in visited:
                        r = [v[0], '{' + v[1] + '}']
                        if r not in remove_var:
                            remove_var.append([v[0], '{' + v[1] + '}'])
                    break
                visited.append([x2, m[x2]])
                queue.append([x2, m[x2]])
            else:
                break

    for var in remove_var:
        variables.remove(var)

    while True:
        bf = True
        for var in variables:
            k = var[0]
            v = var[1]

            answer = answer.replace('{' + k + '}', v)

        for var in variables:
            if var[0] in answer:
                bf = False

        if bf:
            break

    return answer


print(solution("this is {template} {template} is {state}", [["template", "string"], ["state", "changed"]]))
print(solution("this is {template} {template} is {state}", [["template", "string"], ["state", "{template}"]]))
print(solution("this is {template} {template} is {state}", [["template", "{state}"], ["state", "{template}"]]))
print(solution("this is {template} {template} is {state}", [["template", "{state}"], ["state", "{templates}"]]))
print(solution("{a} {b} {c} {d} {i}",
               [["b", "{c}"], ["a", "{b}"], ["e", "{f}"], ["h", "i"], ["d", "{e}"], ["f", "{d}"], ["c", "d"]]))
