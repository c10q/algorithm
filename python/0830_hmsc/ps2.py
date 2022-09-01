def solution(movie):
    answer = []
    temp = {}
    for i in range(len(movie)):
        if movie[i] not in temp:
            temp[movie[i]] = 1
        else:
            temp[movie[i]] += 1

    movies = {}

    for i in temp:
        if temp[i] not in movies:
            movies[temp[i]] = [i]
        else:
            movies[temp[i]].append(i)


    movies = sorted(movies.items(), key=lambda x: x[0], reverse=True)
    for i in movies:
        answer += sorted(i[1])

    return answer


print(solution(["classic", "pop", "classic", "classic", "pop"]))