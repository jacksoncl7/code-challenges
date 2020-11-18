players, turns = list(map(int, str(input()).split()))
game = list(map(int, str(input()).split()))
winner = 0
winner_score = 0

for player in range(players):
    current_score = 0
    for turn in range(turns):
        current_score += game[turn + (turn * (players - 1)) + player]

    if winner_score <= current_score:
        winner_score = current_score
        winner = player + 1

print(winner)
