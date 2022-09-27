from datetime import datetime, timedelta


# 모든 달은 28일까지 있다고 가정합니다
# 2000.01.01 ~ 2022.12.28

def solution(today, terms, privacies):
    answer = []
    terms_map = {}

    today = today.split('.')

    for i in range(len(terms)):
        term = terms[i].split(' ')
        terms_map[term[0]] = term[1]

    for i in range(len(privacies)):
        privacy = privacies[i].split(' ')
        date = privacy[0].split('.')
        term = privacy[1]

        if diff_date(today, add_month(date, int(terms_map[term]))) <= 0:
            answer.append(i + 1)

    return answer


def add_month(date, m):
    year = int(date[0])
    month = int(date[1])
    day = int(date[2])

    month += m
    if month > 12:
        year += month // 12
        month = month % 12

    return [str(year), str(month), str(day)]


def diff_date(date1, date2):
    date1 = date1
    date2 = date2

    year_diff = int(date2[0]) - int(date1[0])
    month_diff = int(date2[1]) - int(date1[1])
    day_diff = int(date2[2]) - int(date1[2])

    return year_diff * 365 + month_diff * 28 + day_diff


if __name__ == '__main__':
    print(solution("2020.01.01", ["Z 3", "D 5"], ["2019.01.01 D", "2019.11.15 Z", "2019.08.02 D", "2019.07.01 D", "2018.12.28 Z"]))
    print(solution("2020.01.01", ["Z 3", "D 100"],
                   ["2000.01.01 D", "2019.11.15 Z", "2019.08.02 D", "2019.07.01 D", "2018.12.28 Z"]))
