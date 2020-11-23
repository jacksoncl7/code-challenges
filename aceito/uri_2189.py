# https://www.urionlinejudge.com.br/judge/pt/problems/view/2189

repeat = int(input())
test = 1

while repeat != 0:
    counter = 1
    entry = list(map(int, str(input()).split()))

    for num in entry:
        if num == counter:
            print('Teste ' + str(test))
            print(str(num) + '\n')
            break
        counter += 1

    test += 1
    repeat = int(input())
