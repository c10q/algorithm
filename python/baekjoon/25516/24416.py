def fibonacci():
    n = int(input())
    f = [0 for _ in range(n + 1)]
    f[1] = 1
    f[2] = 1
    for i in range(3, n + 1):
        f[i] = f[i - 1] + f[i - 2]
    print(f[n], n - 2)


fibonacci()