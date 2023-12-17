package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Line struct {
	chars []string
}

type Pos struct {
	line int
	col  int
}

func main() {
	fmt.Println("P1 ex: ", P1("AOC2023-14/ex.txt"))
	fmt.Println("P1: ", P1("AOC2023-14/input1.txt"))
	fmt.Println("P2 ex:", P2("AOC2023-14/ex.txt"))
	fmt.Println("P2: ", P2("AOC2023-14/input1.txt"))
}

func P1(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	var matrixCol = make([]Line, 0)
	var matrixRow = make([]Line, 0)
	var blockPos = make([]Pos, 0)
	scanner := bufio.NewScanner(file)
	countLine := 0
	for scanner.Scan() {
		line := scanner.Text()
		var lineChar = make([]string, 0)
		countCol := 0
		for _, char := range line {
			lineChar = append(lineChar, string(char))
			if string(char) == "#" {
				blockPos = append(blockPos, Pos{countLine, countCol})
			}
			countCol++
		}
		matrixRow = append(matrixRow, Line{lineChar})
		countLine++
	}
	for i := 0; i < len(matrixRow[0].chars); i++ {
		var lineChar = make([]string, 0)
		for j := 0; j < len(matrixRow); j++ {
			lineChar = append(lineChar, matrixRow[j].chars[i])
		}
		matrixCol = append(matrixCol, Line{lineChar})
	}
	var p1 = 0
	for _, line := range matrixCol {
		var curScore = len(line.chars)
		for i, char := range line.chars {
			if char == "O" {
				p1 += curScore
				curScore--
			} else if char == "#" {
				curScore = len(line.chars) - i - 1
			}
		}
	}
	return p1
}

func P2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	var matrixCol = make([]Line, 0)
	var matrixRow = make([]Line, 0)
	var blockPos = make([]Pos, 0)
	scanner := bufio.NewScanner(file)
	countLine := 0
	for scanner.Scan() {
		line := scanner.Text()
		var lineChar = make([]string, 0)
		countCol := 0
		for _, char := range line {
			lineChar = append(lineChar, string(char))
			if string(char) == "#" {
				blockPos = append(blockPos, Pos{countLine, countCol})
			}
			countCol++
		}
		matrixRow = append(matrixRow, Line{lineChar})
		countLine++
	}
	for i := 0; i < len(matrixRow[0].chars); i++ {
		var lineChar = make([]string, 0)
		for j := 0; j < len(matrixRow); j++ {
			lineChar = append(lineChar, matrixRow[j].chars[i])
		}
		matrixCol = append(matrixCol, Line{lineChar})
	}
	var numCycle = 1000000000
	var p2 = doSolveCycleNum(matrixCol, numCycle)
	return p2
}

func doSolveCycleNum(col []Line, cycle int) int {
	var matrixCol = col
	var mapHistory = make(map[string]int)
	var originalString = matrixToLineChar(matrixCol)
	mapHistory[originalString] = 0
	var end = false
	var maxIter = cycle
	var i = 0
	for !end {
		i++
		matrixCol = doEast(doSouth(doWest(doNorth(matrixCol))))
		var matrixChar = matrixToLineChar(matrixCol)
		if _, ok := mapHistory[matrixChar]; ok {
			//We found the period, we can calculate the rightConfig after cycle cycles
			return calcLoad(findRightMatrix(mapHistory, cycle, mapHistory[matrixChar]))
		}
		mapHistory[matrixChar] = len(mapHistory)
		if i > maxIter {
			end = true
		}
	}
	fmt.Println("No period found :(")
	return 0
}

func showMatrixCol(col []Line) {
	var matrixLine = swapRowAndCols(col)
	for _, line := range matrixLine {
		fmt.Println(line)
	}
}

func calcLoad(matrixCol []Line) int {
	var load = 0
	for _, line := range matrixCol {
		for k, char := range line.chars {
			if char == "O" {
				load += len(line.chars) - k
			}
		}
	}
	return load
}

func findRightMatrix(mapHistory map[string]int, cycle int, periodIndex int) []Line {
	var rightInd = 0
	if cycle < len(mapHistory) {
		//Find the right value checking all the keys of the map
		rightInd = cycle
	} else {
		rightInd = ((cycle - len(mapHistory)) % (len(mapHistory) - periodIndex)) + periodIndex

	}
	var rightMatrixChar = ""
	for key, value := range mapHistory {
		if value == rightInd {
			rightMatrixChar = key
			break
		}
	}
	var matrixCol = charToMatrixCol(rightMatrixChar)
	return matrixCol
}

func doNorth(col []Line) []Line {
	var matrixCol = col
	var newMatrixCol = make([]Line, 0)
	for i := 0; i < len(matrixCol); i++ {
		var newCol = make([]string, 0)
		var countPoints = 0
		for j := 0; j < len(matrixCol[i].chars); j++ {
			if matrixCol[i].chars[j] == "." {
				countPoints++
			} else if matrixCol[i].chars[j] == "#" {
				for k := 0; k < countPoints; k++ {
					newCol = append(newCol, ".")
				}
				countPoints = 0
				newCol = append(newCol, "#")
			} else if matrixCol[i].chars[j] == "O" {
				newCol = append(newCol, "O")
			}
		}
		for k := 0; k < countPoints; k++ {
			newCol = append(newCol, ".")
		}
		newMatrixCol = append(newMatrixCol, Line{newCol})
	}
	return newMatrixCol
}

func doSouth(col []Line) []Line {
	var matrixCol = col
	var newMatrixCol = make([]Line, 0)
	for i := 0; i < len(matrixCol); i++ {
		var countPoints = 0
		var newCol = make([]string, 0)
		for j := 1; j <= len(matrixCol[i].chars); j++ {
			var index = len(matrixCol[i].chars) - j
			if matrixCol[i].chars[index] == "." {
				countPoints++
			} else if matrixCol[i].chars[index] == "#" {
				for k := 0; k < countPoints; k++ {
					newCol = append(newCol, ".")
				}
				countPoints = 0
				newCol = append(newCol, "#")
			} else if matrixCol[i].chars[index] == "O" {
				newCol = append(newCol, "O")
			}
		}
		for k := 0; k < countPoints; k++ {
			newCol = append(newCol, ".")
		}
		newMatrixCol = append(newMatrixCol, reverseList(Line{newCol}))
	}
	return newMatrixCol
}

func doWest(col []Line) []Line {
	var matrixRow = swapRowAndCols(col)
	var newMatrixRow = make([]Line, 0)
	for i := 0; i < len(matrixRow); i++ {
		var countPoints = 0
		var newRow = make([]string, 0)
		for j := 0; j < len(matrixRow[i].chars); j++ {
			if matrixRow[i].chars[j] == "." {
				countPoints++
			} else if matrixRow[i].chars[j] == "#" {
				for k := 0; k < countPoints; k++ {
					newRow = append(newRow, ".")
				}
				countPoints = 0
				newRow = append(newRow, "#")
			} else if matrixRow[i].chars[j] == "O" {
				newRow = append(newRow, "O")
			}
		}
		for k := 0; k < countPoints; k++ {
			newRow = append(newRow, ".")
		}
		newMatrixRow = append(newMatrixRow, Line{newRow})
	}
	return swapRowAndCols(newMatrixRow)
}

func doEast(col []Line) []Line {
	var matrixRow = swapRowAndCols(col)
	var newMatrixRow = make([]Line, 0)
	for i := 0; i < len(matrixRow); i++ {
		var countPoints = 0
		var newRow = make([]string, 0)
		for j := 1; j <= len(matrixRow[i].chars); j++ {
			var index = len(matrixRow[i].chars) - j
			if matrixRow[i].chars[index] == "." {
				countPoints++
			} else if matrixRow[i].chars[index] == "#" {
				for k := 0; k < countPoints; k++ {
					newRow = append(newRow, ".")
				}
				countPoints = 0
				newRow = append(newRow, "#")
			} else if matrixRow[i].chars[index] == "O" {
				newRow = append(newRow, "O")
			}
		}
		for k := 0; k < countPoints; k++ {
			newRow = append(newRow, ".")
		}
		newMatrixRow = append(newMatrixRow, reverseList(Line{newRow}))
	}
	return swapRowAndCols(newMatrixRow)
}

func swapRowAndCols(row []Line) []Line {
	var matrixCol = make([]Line, 0)
	for i := 0; i < len(row[0].chars); i++ {
		var lineChar = make([]string, 0)
		for j := 0; j < len(row); j++ {
			lineChar = append(lineChar, row[j].chars[i])
		}
		matrixCol = append(matrixCol, Line{lineChar})
	}
	return matrixCol

}

func matrixToLineChar(matrix []Line) string {
	var matrixChar = ""
	for _, line := range matrix {
		for _, char := range line.chars {
			matrixChar += char
		}
		matrixChar += "\n"
	}
	return matrixChar
}

func charToMatrixCol(matrixChar string) []Line {
	var matrixColChar = strings.Split(matrixChar, "\n")
	var matrixCol = make([]Line, 0)
	for _, line := range matrixColChar {
		if line != "" {
			matrixCol = append(matrixCol, Line{strings.Split(strings.TrimSpace(line), "")})
		}
	}
	return matrixCol
}

func reverseList(col Line) Line {
	length := len(col.chars)
	var newChars = make([]string, 0)
	for i := 1; i <= length; i++ {
		var index = length - i
		newChars = append(newChars, col.chars[index])
	}
	col.chars = newChars
	return col
}
