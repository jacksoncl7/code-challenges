# https://www.urionlinejudge.com.br/judge/pt/problems/view/2366

def route_check(quantum, limit, stations):
    max_route = quantum * limit 
    if max_route < 42195:
        return 'N'
    elif max_route < (42195 - limit):
        return 'N'

    counter = 0
    for station in stations:
        if counter < station:
            return 'N'
        counter += limit

    if counter < 42195: 
        return 'N'
    return 'S'

water, limit = list(map(int, str(input()).split()))
stations = list(map(int, str(input()).split()))
print(route_check(water, limit, stations))
