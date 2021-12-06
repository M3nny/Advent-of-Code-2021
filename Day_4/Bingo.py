file = 'input.txt'

nCalled = [15,61,32,33,87,17,56,73,27,83,0,18,43,8,86,37,40,6,93,25,14,68,64,57,39,46,55,13,21,72,51,81,78,79,52,65,36,92,97,28,9,24,22,69,70,42,3,4,63,50,91,16,41,94,77,85,49,12,76,67,11,62,99,54,95,1,74,34,88,89,82,48,84,98,58,44,5,53,7,19,29,30,35,26,31,10,60,59,80,71,45,38,20,66,47,2,23,96,90,75]

container = []
bingo_table = []
isCalled = []

for line in open(file):
    line = line.strip()
    if line:
        bingo_table.append([int(i) for i in line.split()])

    else:
        if bingo_table != 0:
            container.append(bingo_table)
        bingo_table = []
container.append(bingo_table)


for n in container:
    isCalled.append([[0 for i in range(5)] for i in range(5)])

winner_table = [0 for i in range(len(container))]
isFirst = 0


for n in nCalled:
    for i in range(len(container)):
        for row in range(5):
            for column in range(5):
                if container[i][row][column] == n:
                    #puts a checkmark in the isCalled table
                    isCalled[i][row][column] = 1

        
        if row == 4:
            #check if the row makes a bingo
            cont = 0
            for row in range(5):
                tmp = 1
                for column in range(5):
                    if column != isCalled[i][row][column]:
                        tmp = 0
                if tmp == 1:
                    cont = 1

            #check if the column makes a bingo
            for column in range(5):
                tmp = 1
                for row in range(5):
                    if isCalled[i][row][column] == 0:
                        tmp = 0
                if tmp == 1:
                    cont = 1
            
            #if this is the first bingo
            if cont == 1 and cont != winner_table[i]:
                    #puts a checkmark in the winner_table
                    winner_table[i] = 1
                    newinner = len([j for j in range(len(container)) if winner_table[j] == 1])
                    if newinner == 1 or newinner == len(container):
                        unmarked = 0
                        for row in range(5):
                            for column in range(5):
                                if isCalled[i][row][column] == 0:
                                    unmarked += container[i][row][column]
                        isFirst += 1
                        if isFirst == 1:
                            print("first winning table: ", unmarked * n)


