def scaling(string):
    result = []
    depth = 0
    count = 0
    temp = ''
    for c in string:
        if c.isdigit():
            if depth == 0:
                count = int(str(count) + c)
            else:
                temp += c
        elif c == '(':
            if depth != 0:
                temp += c
            depth += 1
        elif c == ')':
            depth -= 1
            if depth != 0:
                temp += c
            else:
                if '(' in temp:
                    temp = scaling(temp)

                t = ''
                for i in range(count):
                    t += temp

                result.append(t)
                depth = 0
                count = 0
                temp = ''
        else:
            if depth == 0:
                result.append(c)
            else:
                temp += c

    answer = ''

    for i in result:
        answer += i

    return answer


def solution(compressed):
    return scaling(compressed)


if __name__ == '__main__':
    print(scaling('2(2(hi)2(co))x2(bo)'))  ## 2(hi)2(co)2(hi)2(co)xbobo

    #print(solution("3(hi)") == "hihihi")
    #print(solution("2(3(hi)co)") == "hihihicohihihico")
    print(solution("10(p)") == "pppppppppp")
    #print(solution("2(2(hi)2(co))x2(bo)") == "hihicocohihicocoxbobo")
