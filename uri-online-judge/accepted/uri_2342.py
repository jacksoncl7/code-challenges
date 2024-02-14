# https://www.urionlinejudge.com.br/judge/pt/problems/view/2342

limit = int(input())
equation = str(input()).split()
num1, function, num2 = int(equation[0]), equation[1], int(equation[2])
if function == '*':
    result = num1 * num2
else:
    result = num1 + num2

if result > limit:
    print('OVERFLOW')
else:
    print('OK')
