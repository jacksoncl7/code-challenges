code =  input()
word =  input()
equal = False
num = 0
for i in range(len(code) - len(word) + 1):
#    print('---------------- NEW LOOP --------------------')
    for j in range(len(word)):
#        print('i+j => {} ATAQUE => {} BALBLAB => {}'.format(str(i +j), word[j], code[i + j]))
        if word[j] == code[i + j]:
#            print(code[i + j])
            equal = True
            break
    if equal:
        equal = False
    else:
        num += 1
print(num)
