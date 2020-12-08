tamanho, frango = map(int, input().split())
choc = tuple(map(int, input().split()))
ctr = list(choc)
dias = 0
rng = range(tamanho)

while frango != 0:
    ctr = list(map(lambda x: x - 1, ctr))
    dias += 1
    while 0 in ctr:
        frango -= 1
        i = ctr.index(0)
        ctr[i] = choc[i]
print(dias)
