# https://www.urionlinejudge.com.br/judge/pt/problems/view/2373

trays = int(input())
broken_glasses = 0
for i in range(trays):
    can, glass = list(map(int, str(input()).split()))
    #waiter_list = list(map(int, str(input()).split()))
    #can, glass = waiter_list[0], waiter_list[1]
    if can > glass:
        broken_glasses += glass
print(broken_glasses)
