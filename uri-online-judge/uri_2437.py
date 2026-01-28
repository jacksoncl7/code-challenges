# https://www.urionlinejudge.com.br/judge/en/problems/view/2437

x0, y0, x1, y1 = tuple(map(int, str(input()).split()))
print(abs(x0 -x1) + abs(y0 - y1))
