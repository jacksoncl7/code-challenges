# https://www.urionlinejudge.com.br/judge/pt/problems/view/2420

size = int(input())
territory = list(map(int, str(input()).split()))
limit = sum(territory) / 2
counter = 1
acc = 0
for number in territory:
    acc += number
    if acc >= limit:
        print(counter)
        break
    counter += 1
