def validation_pack(packs):
    packs.sort()
    aux = packs[0]
    del packs[0]

    if aux > 8:
        return 'N'
    for val in packs:
        if abs(val - aux) > 8:
            return 'N'
        aux = val
    return 'S'

num_packs = int(input())
packs = list(map(int, input().split()))
print(validation_pack(packs))
