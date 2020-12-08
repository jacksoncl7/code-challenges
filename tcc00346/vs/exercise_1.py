repeat = int(input())
pssw = []
for _ in range(repeat):
    pssw.append(input().split())

for ar in pssw:
    resp = ''
    for word in ar:
        resp += word[0]
    print(resp)
