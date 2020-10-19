repeat = int(input())
test = 1

while repeat != 0:
    player_one = str(input())
    player_two = str(input())
    winners = []

    for i in range(repeat):
        input_games = list(map(int, input().split()))
        winner = player_one if sum(input_games) % 2 == 0 else player_two
        winners.append(winner)

    print("Teste " + str(test))
    for winner in winners:
        print(winner)
    print()

    test += 1
    repeat = int(input())
