# https://www.urionlinejudge.com.br/judge/en/problems/view/2304 

init, repeat = tuple(map(int, str(input()).split()))
ctrl = { 'D': init, 'E': init, 'F': init }
for _ in range(repeat):
    turn = str(input()).split()
    if turn[0] == 'C':
        ctrl[turn[1]] -= int(turn[2])
    elif turn[0] == 'V':
        ctrl[turn[1]] += int(turn[2])
    else:
        ctrl[turn[2]] -= int(turn[3])
        ctrl[turn[1]] += int(turn[3])

print('{} {} {}'.format(ctrl['D'], ctrl['E'], ctrl['F']))
