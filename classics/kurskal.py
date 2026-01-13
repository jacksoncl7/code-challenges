import heapq

n, m = input().split()         # ler numero de vertices e arestas
n = int(n)
m = int(m)

H = []

for j in range(m):                # ler as m arestas do digrafo
    a, b, c = input().split()     # ler aresta de a para b com custo c
    a = int(a)
    b = int(b)
    c = int(c)
    heapq.heappush(H, (c, a, b))  # colocar aresta no heap

C = [[] * n for i in range(n)]    # criar n conjuntos

for i in range(n):
    C[i].append(i)                # cada C[i] é inicializado com {i}

S = []
for i in range(n):
    S.append(i)                   # S[i] é o conjunto ao qual o vértice i pertence

cont = 0
custo = 0

print(C)
print(S)

while cont < n-1:                 # bastam n-1 arestas
    c, a, b = heapq.heappop(H)    # remover a próxima aresta do heap
    if S[a] != S[b]:              # se as arestas unem árvores diferentes...
        custo = custo + c
        p = S[a]
        q = S[b]
        if q < p:
            p, q = q, p
        for j in C[q]:           # para cada j no conjunto C[q]
            S[j] = p
        C[p].extend(C[q])        # unir C[p] e C[q] (a união fica em C[p])
        C[q] = []                # esvaziar C[q]
        cont = cont + 1
        print(C)
        print(S)
print(custo)

'''
#entrada de dados

9 14
0 1 4
0 7 8
1 2 8
1 7 11
2 3 7
2 5 4
2 8 2
3 4 9
3 5 14
4 5 10
5 6 2
6 7 1
6 8 6
7 8 7

'''
