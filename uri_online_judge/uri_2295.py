#https://www.urionlinejudge.com.br/judge/en/problems/view/2295

a, ra, g, rg = tuple(map(float, str(input()).split()))
al = ra / a, 2
gs = rg / g, 2

if al > gs:
    print('A')
else:
    print('G')
