# https://www.urionlinejudge.com.br/judge/pt/problems/view/2344

score = int(input())

if score == 0:
    print('E')
elif score < 36:
    print('D')
elif score < 61:
    print('C')
elif score < 86:
    print('B')
elif score > 85:
    print('A')
