
def ascii_shift(string, shift):
    new_string = ''
    for character in string:
        new_ascii = ord(character) + shift
        new_string += chr(new_ascii)

    return new_string

string = str(input())
string_passone = ascii_shift(string, 3) 
print(string_passone)
string_passtwo = string_passone[::-1]
print(string_passtwo)
string_passthree = ascii_shift(string, -1) 
print(string_passthree)

