import numpy as np
import re

file = 'input.txt'
ocean_floor = np.zeros((1000,1000))
overlap = 0

for line in open(file):
    regex = '\d+'            
    match = re.findall(regex, line)
    x1 = int(match[0])
    y1 = int(match[1])
    x2 = int(match[2])
    y2 = int(match[3])

    max_x = max(x1,x2)
    min_x = min(x1,x2)
    max_y = max(y1,y2)
    min_y = min(y1,y2)


    if x1==x2:
        for i in range(min_y, max_y + 1):
            ocean_floor[i][x1] +=1
    elif y1==y2: 
        for i in range(min_x, max_x + 1):
            ocean_floor[y1][i] +=1



    
#print(ocean_floor)

for i in range(1000000):
    if ocean_floor.item(i) > 1:
        overlap += 1

print("there are:", overlap, "overlappings")
            
    



      