# https://www.urionlinejudge.com.br/judge/pt/problems/view/2456

def crescente(lista):
    aux = lista[0]
    for num in lista:
        if num < aux:
            return False
        aux = num

    return True

lista = list(map(int, str(input()).split()))
if crescente(lista):
    print('C')
elif crescente(lista[::-1]):
    print('D')
else:
    print('N')
