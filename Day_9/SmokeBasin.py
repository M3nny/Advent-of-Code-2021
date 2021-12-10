file = open('input.txt')

def SmallestNearby(column,row):
    near = []
    if column+1 < len(basin[0]):
        near.append(basin[row][column+1])
    if column-1 >= 0:
        near.append(basin[row][column-1])
    if row+1 < len(basin):
        near.append(basin[row+1][column])
    if row-1 >= 0:
        near.append(basin[row-1][column])
    return min(near)


basin = []
for line in file:
    line = line.strip()
    basin.append([ int(char) for char in line ])

result = 0
for row in range(len(basin)):
    for column in range(len(basin[0])):
        if basin[row][column] < SmallestNearby(column,row):
            result += 1 + basin[row][column]

print(result)