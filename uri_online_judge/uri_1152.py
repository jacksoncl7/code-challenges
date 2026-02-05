import heapq

def kruskal(n, m):
    H = []
    total = 0

    for j in range(m):
        a, b, c = tuple(map(int, input().split()))
        total += c
        heapq.heappush(H, (c, a, b))

    C = [[] * n for i in range(n)]

    for i in range(n):
        C[i].append(i)

    S = []
    for i in range(n):
        S.append(i)

    cont = 0
    custo = 0

    while cont < n-1:
        c, a, b = heapq.heappop(H)
        if S[a] != S[b]:
            custo = custo + c
            p = S[a]
            q = S[b]
            if q < p:
                p, q = q, p
            for j in C[q]:
                S[j] = p
            C[p].extend(C[q])
            C[q] = []
            cont = cont + 1
    return total - custo


n, m = tuple(map(int, input().split()))
while n != 0 and m != 0:
    print(kruskal(n, m))
    n, m = tuple(map(int, input().split()))
