#https://www.urionlinejudge.com.br/judge/en/problems/view/2293

def max_sum_column(matrix, length_col):
    col = 0
    maximun = 0
    while col != length_col:
        partial = 0
        for line in matrix:
            partial += line[col] 
        col += 1
        if partial > maximun:
            maximun = partial
    return maximun

def max_sum_line(matrix):
    maximun = 0
    for line in matrix:
        partial = sum(line)
        if partial > maximun:
            maximun = partial
    return maximun


line, col = tuple(map(int, str(input()).split()))
matrix = []
for _ in range(line):
    matrix.append(list(map(int, str(input()).split())))

lines = max_sum_line(matrix)
columns = max_sum_column(matrix, col)
if lines > columns:
    print(lines)
else:
    print(columns)
