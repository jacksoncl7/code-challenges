# https://www.urionlinejudge.com.br/judge/pt/problems/view/2343
'''
size = int(input())
response = 0
registries = []

for i in range(size):
    registry = tuple(map(int, str(input()).split()))
    if registry in registries: response = 1
    registries.append(registry)

print(response)
'''

size = int(input())
response = 0
registries = []

for i in range(size):
    registry = tuple(map(int, str(input()).split()))
    registries.append(registry)

length = len(set(registries))
response = 1 if length != len(registries) else 0
print(response)
