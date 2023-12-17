package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x int
	y int
}

func main() {
	fmt.Println("Part 1 Ex: ", P1("ex.txt"))
	fmt.Println("Part 1 : ", P1("input1.txt"))
	fmt.Println("Part 2 Ex1: ", P2("ex1.txt"))
	fmt.Println("Part 2 Ex2: ", P2("ex2.txt"))
	fmt.Println("Part 2 Ex3: ", P2("ex3.txt"))
	fmt.Println("Part 2 : ", P2("input1.txt"))
}

func P1(input string) int {
	var matrix = make([][]string, 0)
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	var posS Position
	for scanner.Scan() {
		line := scanner.Text()
		var curLine = make([]string, 0)
		for i, c := range line {
			if c == 'S' {
				posS = Position{count, i}
			}
			curLine = append(curLine, string(c))
		}
		matrix = append(matrix, curLine)
		count++
	}
	var isLoop = false
	var firstNeighb = posS
	var secondNeighb = posS
	if posS.x > 0 {
		var char = matrix[posS.x-1][posS.y]
		if char == "7" || char == "F" || char == "|" {
			if checkLoop(matrix, Position{posS.x - 1, posS.y}, posS) {
				firstNeighb = Position{posS.x - 1, posS.y}
				isLoop = true
			}
		}
	}
	if posS.x < len(matrix)-1 {
		var char = matrix[posS.x+1][posS.y]
		if char == "J" || char == "L" || char == "|" {
			if checkLoop(matrix, Position{posS.x + 1, posS.y}, posS) {
				if isLoop {
					secondNeighb = Position{posS.x + 1, posS.y}
				} else {
					firstNeighb = Position{posS.x + 1, posS.y}
				}
				isLoop = true
			}
		}
	}
	if posS.y > 0 {
		var char = matrix[posS.x][posS.y-1]
		if char == "L" || char == "F" || char == "-" {
			if checkLoop(matrix, Position{posS.x, posS.y - 1}, posS) {
				if isLoop {
					secondNeighb = Position{posS.x, posS.y - 1}
				} else {
					firstNeighb = Position{posS.x, posS.y - 1}
				}
				isLoop = true
			}
		}
	}
	if posS.y < len(matrix[0])-1 {
		var char = matrix[posS.x][posS.y+1]
		if char == "J" || char == "7" || char == "-" {
			if checkLoop(matrix, Position{posS.x, posS.y + 1}, posS) {
				if isLoop {
					secondNeighb = Position{posS.x, posS.y + 1}
				} else {
					firstNeighb = Position{posS.x, posS.y + 1}
				}
				isLoop = true
			}
		}
	}
	var distMatrix = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		distMatrix[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			distMatrix[i][j] = -1
		}
	}
	distMatrix = updateDistMatrix(matrix, posS, distMatrix, firstNeighb, secondNeighb)
	var maxVal = 0
	for i := 0; i < len(distMatrix); i++ {
		for j := 0; j < len(distMatrix[0]); j++ {
			if distMatrix[i][j] > maxVal {
				maxVal = distMatrix[i][j]
			}
		}
	}
	return maxVal
}

func P2(input string) int {
	var matrix = make([][]string, 0)
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	var posS Position
	for scanner.Scan() {
		line := scanner.Text()
		var curLine = make([]string, 0)
		for i, c := range line {
			if c == 'S' {
				posS = Position{count, i}
			}
			curLine = append(curLine, string(c))
		}
		matrix = append(matrix, curLine)
		count++
	}
	var isLoop = false
	var firstNeighb = posS
	var secondNeighb = posS
	if posS.x > 0 {
		var char = matrix[posS.x-1][posS.y]
		if char == "7" || char == "F" || char == "|" {
			if checkLoop(matrix, Position{posS.x - 1, posS.y}, posS) {
				firstNeighb = Position{posS.x - 1, posS.y}
				isLoop = true
			}
		}
	}
	if posS.x < len(matrix)-1 {
		var char = matrix[posS.x+1][posS.y]
		if char == "J" || char == "L" || char == "|" {
			if checkLoop(matrix, Position{posS.x + 1, posS.y}, posS) {
				if isLoop {
					secondNeighb = Position{posS.x + 1, posS.y}
				} else {
					firstNeighb = Position{posS.x + 1, posS.y}
				}
				isLoop = true
			}
		}
	}
	if posS.y > 0 {
		var char = matrix[posS.x][posS.y-1]
		if char == "L" || char == "F" || char == "-" {
			if checkLoop(matrix, Position{posS.x, posS.y - 1}, posS) {
				if isLoop {
					secondNeighb = Position{posS.x, posS.y - 1}
				} else {
					firstNeighb = Position{posS.x, posS.y - 1}
				}
				isLoop = true
			}
		}
	}
	if posS.y < len(matrix[0])-1 {
		var char = matrix[posS.x][posS.y+1]
		if char == "J" || char == "7" || char == "-" {
			if checkLoop(matrix, Position{posS.x, posS.y + 1}, posS) {
				if isLoop {
					secondNeighb = Position{posS.x, posS.y + 1}
				} else {
					firstNeighb = Position{posS.x, posS.y + 1}
				}
				isLoop = true
			}
		}
	}
	var distMatrix = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		distMatrix[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			distMatrix[i][j] = -1
		}
	}
	distMatrix = updateDistMatrix(matrix, posS, distMatrix, firstNeighb, secondNeighb)
	for i := 0; i < len(distMatrix); i++ {
		for j := 0; j < len(distMatrix[0]); j++ {
			if distMatrix[i][j] == -1 {
				matrix[i][j] = "."
			}
		}
	}
	var curCrossed = 0
	var countIn = 0
	for i := 0; i < len(matrix); i++ {
		curCrossed = 0
		for j := 0; j < len(matrix[0]); j++ {
			//We need no count how many times we "crossed" the border. If its odd, we are in the loop
			if matrix[i][j] == "|" || matrix[i][j] == "L" || matrix[i][j] == "J" {
				curCrossed++
			} else if matrix[i][j] == "." {
				if curCrossed%2 == 1 {
					countIn++
				}

			}
		}
	}
	return countIn
}

func findAvailablePosition(curChar string, pos Position, matrix [][]string) []Position {
	var possiblePos = make([]Position, 0)
	if curChar == "-" {
		if pos.y > 0 {
			possiblePos = append(possiblePos, Position{pos.x, pos.y - 1})
		}
		if pos.y < len(matrix[0])-1 {
			possiblePos = append(possiblePos, Position{pos.x, pos.y + 1})
		}
	} else if curChar == "|" {
		if pos.x > 0 {
			possiblePos = append(possiblePos, Position{pos.x - 1, pos.y})
		}
		if pos.x < len(matrix)-1 {
			possiblePos = append(possiblePos, Position{pos.x + 1, pos.y})
		}
	} else if curChar == "L" {
		if pos.x > 0 {
			possiblePos = append(possiblePos, Position{pos.x - 1, pos.y})
		}
		if pos.y < len(matrix[0])-1 {
			possiblePos = append(possiblePos, Position{pos.x, pos.y + 1})
		}
	} else if curChar == "F" {
		if pos.x < len(matrix)-1 {
			possiblePos = append(possiblePos, Position{pos.x + 1, pos.y})
		}
		if pos.y < len(matrix[0])-1 {
			possiblePos = append(possiblePos, Position{pos.x, pos.y + 1})
		}
	} else if curChar == "7" {
		if pos.x < len(matrix)-1 {
			possiblePos = append(possiblePos, Position{pos.x + 1, pos.y})
		}
		if pos.y > 0 {
			possiblePos = append(possiblePos, Position{pos.x, pos.y - 1})
		}
	} else if curChar == "J" {
		if pos.x > 0 {
			possiblePos = append(possiblePos, Position{pos.x - 1, pos.y})
		}
		if pos.y > 0 {
			possiblePos = append(possiblePos, Position{pos.x, pos.y - 1})
		}
	}
	return possiblePos
}

func updateDistMatrix(matrix [][]string, s Position, distMatrix [][]int, firstNeigh Position, secondNeigh Position) [][]int {
	var pastPos = s
	var pos = firstNeigh
	var curDist = 2
	distMatrix[firstNeigh.x][firstNeigh.y] = 1
	var maxCount = len(matrix) * len(matrix[0])
	for i := 0; i < maxCount; i++ {
		var possiblePos = make([]Position, 0)
		var curChar = matrix[pos.x][pos.y]
		if curChar == "S" {
			break
		} else {
			possiblePos = findAvailablePosition(curChar, pos, matrix)
		}
		if len(possiblePos) == 1 {
			distMatrix[possiblePos[0].x][possiblePos[0].y] = curDist
			curDist++
			pastPos = pos
			pos = possiblePos[0]
			curChar = matrix[pos.x][pos.y]
		} else if len(possiblePos) == 2 {
			if possiblePos[0] == pastPos {
				distMatrix[possiblePos[1].x][possiblePos[1].y] = curDist
				curDist++
				pastPos = pos
				pos = possiblePos[1]
				curChar = matrix[pos.x][pos.y]
			} else {
				distMatrix[possiblePos[0].x][possiblePos[0].y] = curDist
				curDist++
				pastPos = pos
				pos = possiblePos[0]
				curChar = matrix[pos.x][pos.y]
			}
		}
	}
	pastPos = s
	pos = secondNeigh
	distMatrix[secondNeigh.x][secondNeigh.y] = 1
	curDist = 2
	maxCount = len(matrix) * len(matrix[0])
	for i := 0; i < maxCount; i++ {
		var possiblePos = make([]Position, 0)
		var curChar = matrix[pos.x][pos.y]
		if curChar == "S" {
			break
		} else {
			possiblePos = findAvailablePosition(curChar, pos, matrix)
		}
		if len(possiblePos) == 1 {
			if curDist < distMatrix[possiblePos[0].x][possiblePos[0].y] {
				distMatrix[possiblePos[0].x][possiblePos[0].y] = curDist
			}
			curDist++
			pastPos = pos
			pos = possiblePos[0]
			curChar = matrix[pos.x][pos.y]
		} else if len(possiblePos) == 2 {
			if possiblePos[0] == pastPos {
				if curDist < distMatrix[possiblePos[1].x][possiblePos[1].y] {
					distMatrix[possiblePos[1].x][possiblePos[1].y] = curDist
				}
				curDist++
				pastPos = pos
				pos = possiblePos[1]
				curChar = matrix[pos.x][pos.y]
			} else {
				if curDist < distMatrix[possiblePos[0].x][possiblePos[0].y] {
					distMatrix[possiblePos[0].x][possiblePos[0].y] = curDist
				}
				curDist++
				pastPos = pos
				pos = possiblePos[0]
				curChar = matrix[pos.x][pos.y]
			}
		}
	}
	distMatrix[s.x][s.y] = 0
	return distMatrix
}

func checkLoop(matrix [][]string, pos Position, posStarting Position) bool {
	var maxCount = len(matrix) * len(matrix[0])
	var pastPos = posStarting
	for i := 0; i < maxCount; i++ {
		var possiblePos = make([]Position, 0)
		var curChar = matrix[pos.x][pos.y]
		if curChar == "S" {
			return true
		} else {
			possiblePos = findAvailablePosition(curChar, pos, matrix)
		}
		if len(possiblePos) == 1 {
			pastPos = pos
			pos = possiblePos[0]
			curChar = matrix[pos.x][pos.y]
		} else if len(possiblePos) == 2 {
			if possiblePos[0] == pastPos {
				pastPos = pos
				pos = possiblePos[1]
				curChar = matrix[pos.x][pos.y]
			} else {
				pastPos = pos
				pos = possiblePos[0]
				curChar = matrix[pos.x][pos.y]
			}
		} else {
			return false
		}
	}
	return false
}
