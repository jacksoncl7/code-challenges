# https://www.urionlinejudge.com.br/judge/pt/problems/view/2443
def div(num1, num2):
    test = num1 % num2
    if test != 0:
        return num1, num2
    while num2 != 1:
        num1 = num1 / num2
        num2 = num2 / num2
        if not(print1 % print2):
           return num1, num2

    return int(num1), int(num2)

def mul(num1, num2):
   limit = min(num1, num2)

   for i in range(int(limit) - 1, 0, -1):
       if num1 % i == 0 and num2 % i == 0:
           num1 = num1 / i
           num2 = num2 / i

   return int(num1), int(num2)

if __name__ == '__main__':
    num1, num2, num3, num4 = list(map(int, str(input()).split()))
    print1 = num1 * num4 + num2 * num3
    print2 = num2 * num4
    print1, print2 = div(print1, print2)
    print1, print2 = mul(print1, print2)
    print("{} {}".format(print1, print2))


