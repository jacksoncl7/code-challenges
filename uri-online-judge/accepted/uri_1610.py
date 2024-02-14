# https://www.urionlinejudge.com.br/judge/pt/problems/view/1610 

def find_loop():
    n, m = list(map(int, str(input()).split()))
    out = [[] * n for i in range(n)]
    for j in range(m):
        a, b = list(map(int, str(input()).split()))
        out[a-1].append(b-1)
    marca = n * [0]
    ciclo = False
    for a in range(n):
        if marca[a] == 0:
            pilha = [a]
            while pilha != [] and not ciclo:
                v = pilha[len(pilha)-1]
                marca[v] = 1
                if out[v] != []:
                    for w in out[v]:
                        del out[v][0]
                        if marca[w] == 0:
                            pilha.append(w)
                            break
                        elif w in pilha:
                             ciclo = True
                else:
                    pilha.pop()
    if ciclo:
        return 'SIM'
    else:
        return 'NAO'

repeat = int(input())
for _ in range(repeat):
    print(find_loop())

