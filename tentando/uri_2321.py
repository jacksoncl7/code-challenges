# https://www.urionlinejudge.com.br/judge/en/problems/view/2321
# at the real life I'll make a object or a dict, but...
def evaluation(x2, y2, x3, y3):
    if x2 >= x3 and y2 >= y3:
        return 1
    elif x2 >= x3 and y2 == y3:
        return 1
    elif x2 == x3 and y2 <= y3:
        return 1
    return 0


x1, y1, x2, y2 = list(map(int, str(input()).split()))
x3, y3, x4, y4 = list(map(int, str(input()).split()))

response = evaluation(x2, y2, x3, y3)
if not(response):
    response = evaluation(x4, y4, x1, y1)
print(response)
