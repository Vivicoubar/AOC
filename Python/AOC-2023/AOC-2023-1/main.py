

def P1(input):
    values = []
    with open(input, 'r') as file:
        for line in file:
            firstChar = ""
            lastChar = ""
            for char in line:
                if char.isnumeric():
                    if firstChar == "":
                        firstChar = char
                    else:
                        lastChar = char
            if lastChar == "":
                lastChar = firstChar
            values.append(int(firstChar + lastChar))
    return sum(values)

def P2(input):
    values = []
    with open(input, 'r') as file:
        for line in file:
            firstChar = ""
            lastChar = ""
            curline = ""
            for char in line:
                curline += char
                curline = curline.replace("one", "1e")
                curline = curline.replace("two", "2o")
                curline = curline.replace("three", "3e")
                curline = curline.replace("four", "4")
                curline = curline.replace("five", "5e")
                curline = curline.replace("six", "6")
                curline = curline.replace("seven", "7n")
                curline = curline.replace("eight", "8t")
                curline = curline.replace("nine", "9e")
                curline = curline.strip()
            line = curline
            for char in line:
                if char.isnumeric():
                    if firstChar == "":
                        firstChar = char
                    else:
                        lastChar = char
            if lastChar == "":
                lastChar = firstChar
            values.append(int(firstChar + lastChar))
    return sum(values)

print(P1("ex.txt"))
print(P1("input.txt"))
print(P2("ex2.txt"))
print(P2("input.txt"))

