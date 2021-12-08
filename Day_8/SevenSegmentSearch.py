file = open('input.txt', 'r')

cont = 0
for line in file:
    line = line.split("| ")[1].split("\n")[0]
    for word in line.split():
        if len(word) == 2 or len(word) == 3 or len(word) == 4 or len(word) == 7:
            cont += 1

print ("1, 4, 7 or 8 appear", cont, "times")