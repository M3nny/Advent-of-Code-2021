file = open("input.txt")
syntax = []
score = 0

for line in file:
    syntax.append(line)

for line in syntax:
    container = []

    for char in line:
        if char == "{" or char == "[" or char == "(" or char == "<":
            container.append(char)
            
        elif char == ")" or char == "]" or char == "}" or char == ">":

            if char == ">" and container[-1] == "<":
                container.pop()
            
            elif char == ")" and container[-1] == "(":
                container.pop()

            elif char == "]" and container[-1] == "[":
                container.pop()

            elif char == "}" and container[-1] == "{":
                container.pop()
                
            else:
                if char == ")":
                    score += 3
                elif char == "]":
                    score += 57
                elif char == "}":
                    score += 1197
                elif char == ">":
                    score += 25137
                break
    
print("total syntax score is:", score)

