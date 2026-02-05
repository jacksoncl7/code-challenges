# https://www.urionlinejudge.com.br/judge/pt/problems/view/1397

test = int(input())

while test != 0:
    left = 0
    right = 0
    for i in range(test):
        game = list(map(int, str(input()).split()))
        if game[0] > game[1]:
            left += 1
        elif game[0] < game[1]:
            right += 1

    print(str(left) + ' ' + str(right))
    test = int(input())
