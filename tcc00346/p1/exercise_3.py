def main():
    p1, p2, p3 =  list(map(int, str(input()).split()))
    if (p1 == p2) or (p2 == p3) or (p3 == p1):
        return 'S'
    elif p1 + p2 - p3 == 0:
        return 'S'
    elif p1 + p3 - p2 == 0:
        return 'S'
    elif p2 + p3 - p1 == 0:
        return 'S'

    elif p1 - p3 - p2 == 0:
        return 'S'
    elif p2 - p3 - p1 == 0:
        return 'S'
    elif p3 - p3 - p2 == 0:
        return 'S'
    return 'N'

response = main()
print(response)
