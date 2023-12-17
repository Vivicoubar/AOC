
class NecessaryCubes:
    def __init__(self, red, green, blue):
        self.red = red
        self.green = green
        self.blue = blue

    def getPower(self):
        return self.red * self.green * self.blue


def P1(input):
    red = 12
    green = 13
    blue = 14
    with open(input, 'r') as file:
        validGames = []
        for line in file:
            line = line.split(": ")[1].strip()
            games = line.split("; ")
            isValid = True
            for game in games:
                boxes = game.split(", ")
                for i in range(len(boxes)):
                    num = int(boxes[i].split(" ")[0])
                    color = boxes[i].split(" ")[1]
                    if color == "red":
                        isValid = num <= red
                    elif color == "green":
                        isValid = num <= green
                    elif color == "blue":
                        isValid = num <= blue
                    if not isValid:
                        break
                if isValid == False:
                    break
            validGames.append(isValid)
        p1 = 0
        for i in range(len(validGames)):
            if validGames[i]:
                p1 += i+1
        return p1

def P2(input):
    red = 12
    green = 13
    blue = 14
    with open(input, 'r') as file:
        necessaryCubes= []
        for line in file:
            line = line.split(": ")[1].strip()
            games = line.split("; ")
            cube = NecessaryCubes(0, 0, 0)
            for game in games:
                boxes = game.split(", ")
                for i in range(len(boxes)):
                    num = int(boxes[i].split(" ")[0])
                    color = boxes[i].split(" ")[1]
                    if color == "red":
                        necessary = num
                        if necessary > cube.red:
                            cube.red = necessary
                    elif color == "green":
                        necessary = num
                        if necessary > cube.green:
                            cube.green = necessary
                    elif color == "blue":
                        necessary = num
                        if necessary > cube.blue:
                            cube.blue = necessary
            necessaryCubes.append(cube)
        p2 = 0
        for i in range(len(necessaryCubes)):
            p2 += necessaryCubes[i].getPower()
        return p2

print(P1("ex.txt"))
print(P1("input.txt"))
print(P2("ex.txt"))
print(P2("input.txt"))
