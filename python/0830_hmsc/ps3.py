def solution(sentence):
    answer = ""
    words = sentence.lower()
    words = words.replace(' ', '')

    alphabets = {
        'a': 0,
        'b': 0,
        'c': 0,
        'd': 0,
        'e': 0,
        'f': 0,
        'g': 0,
        'h': 0,
        'i': 0,
        'j': 0,
        'k': 0,
        'l': 0,
        'm': 0,
        'n': 0,
        'o': 0,
        'p': 0,
        'q': 0,
        'r': 0,
        's': 0,
        't': 0,
        'u': 0,
        'v': 0,
        'w': 0,
        'x': 0,
        'y': 0,
        'z': 0
    }

    for word in words:
        if word in alphabets:
            alphabets[word] += 1

    for alphabet in alphabets:
        if alphabets[alphabet] == 0:
            answer += alphabet

    return answer if answer else "perfect"

print(solution("Jackdaws love my big sphinx of quartz"))