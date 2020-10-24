#trilhas https://www.urionlinejudge.com.br/judge/pt/problems/view/2296

def trail_height(count, laps):
    distance_left = 0
    distance_right = 0
    aux = laps[0]

    for i in range(1, count):
        if aux < laps[i]:
            distance_left += laps[i] - aux
        aux = laps[i]
    
    aux = laps[-1]
    for i in range(count - 2, -1, -1):
        if aux < laps[i]:
            distance_right += laps[i] - aux
        aux = laps[i]

    return min(distance_left, distance_right)

input_lenght = int(input())
trails = []
trail_number = 1
better_trail_number = 0
better_trail_value = 1001

for i in range(input_lenght):
    trail_gps = list(map(int, str(input()).split()))
    quantum = trail_gps[0]
    heights = trail_gps[1:]
    distance = trail_height(quantum, heights)
    if better_trail_value == 1001 or better_trail_value > distance:
        better_trail_value = distance
        better_trail_number = trail_number
    trail_number += 1

print(better_trail_number)
