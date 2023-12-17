package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line struct {
	chars []string
}

func main() {
	fmt.Println("P1:", P1("AOC2023-13/ex.txt"))
	fmt.Println("P1:", P1("AOC2023-13/input1.txt"))
	fmt.Println("P2:", P2("AOC2023-13/ex.txt"))
	fmt.Println("P2:", P2("AOC2023-13/input1.txt"))
}

func P1(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	var matrixRow = make([]Line, 0)
	var matrixCol = make([]Line, 0)
	var allMatrixRow = make([][]Line, 0)
	var allMatrixCol = make([][]Line, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for i := 0; i < len(matrixRow[0].chars); i++ {
				var chaCol = make([]string, 0)
				for j := 0; j < len(matrixRow); j++ {
					chaCol = append(chaCol, matrixRow[j].chars[i])
				}
				matrixCol = append(matrixCol, Line{chaCol})
			}
			allMatrixCol = append(allMatrixCol, matrixCol)
			allMatrixRow = append(allMatrixRow, matrixRow)
			matrixCol = make([]Line, 0)
			matrixRow = make([]Line, 0)
		} else {
			var chaRow = make([]string, 0)
			for _, char := range line {
				chaRow = append(chaRow, string(char))
			}
			matrixRow = append(matrixRow, Line{chaRow})
		}
	}
	for i := 0; i < len(matrixRow[0].chars); i++ {
		var chaCol = make([]string, 0)
		for j := 0; j < len(matrixRow); j++ {
			chaCol = append(chaCol, matrixRow[j].chars[i])
		}
		matrixCol = append(matrixCol, Line{chaCol})
	}
	allMatrixCol = append(allMatrixCol, matrixCol)
	allMatrixRow = append(allMatrixRow, matrixRow)
	matrixCol = make([]Line, 0)
	matrixRow = make([]Line, 0)
	var p1 = 0
	for i := range allMatrixRow {
		matrixRow = allMatrixRow[i]
		matrixCol = allMatrixCol[i]
		var foundVertPatter = findReflexion(matrixCol)
		var foundHorizPatter = -1
		if foundVertPatter != -1 {
			p1 += foundVertPatter
		} else {
			foundHorizPatter = findReflexion(matrixRow)
			if foundHorizPatter != -1 {
				p1 += foundHorizPatter * 100
			} else {
				fmt.Println("Found nothing on matrix", i, "!")
			}
		}

	}
	return p1
}

func P2(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	var matrixRow = make([]Line, 0)
	var matrixCol = make([]Line, 0)
	var allMatrixRow = make([][]Line, 0)
	var allMatrixCol = make([][]Line, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for i := 0; i < len(matrixRow[0].chars); i++ {
				var chaCol = make([]string, 0)
				for j := 0; j < len(matrixRow); j++ {
					chaCol = append(chaCol, matrixRow[j].chars[i])
				}
				matrixCol = append(matrixCol, Line{chaCol})
			}
			allMatrixCol = append(allMatrixCol, matrixCol)
			allMatrixRow = append(allMatrixRow, matrixRow)
			matrixCol = make([]Line, 0)
			matrixRow = make([]Line, 0)
		} else {
			var chaRow = make([]string, 0)
			for _, char := range line {
				chaRow = append(chaRow, string(char))
			}
			matrixRow = append(matrixRow, Line{chaRow})
		}
	}
	for i := 0; i < len(matrixRow[0].chars); i++ {
		var chaCol = make([]string, 0)
		for j := 0; j < len(matrixRow); j++ {
			chaCol = append(chaCol, matrixRow[j].chars[i])
		}
		matrixCol = append(matrixCol, Line{chaCol})
	}
	allMatrixCol = append(allMatrixCol, matrixCol)
	allMatrixRow = append(allMatrixRow, matrixRow)
	var p2 = 0
	var found = false
	for i := 0; i < len(allMatrixRow); i++ {
		var index = -1
		//For each matrix, count the number of error to form a reflexion, if its 1,  then we have a reflexion with a smudge
		index = findReflexionWithError(allMatrixCol[i])
		if index != -1 {
			p2 += index
			found = true
		} else {
			index = findReflexionWithError(allMatrixRow[i])
			if index != -1 {
				p2 += index * 100
				found = true
			}
		}
		if !found {
			fmt.Println("Found nothing on matrix", i, "!")
		}
	}
	return p2
}

func verifyRef(matrixLine []Line, index int) bool {
	var equals = true
	var i = 0
	for index-1-i >= 0 && index+i < len(matrixLine) && equals {
		var leftInd = index - 1 - i
		var rightInd = index + i
		if !charEquals(matrixLine[leftInd].chars, matrixLine[rightInd].chars) {
			equals = false
		}
		i++
	}
	return equals
}
func findReflexion(matrixLine []Line) int {
	for i := 1; i < len(matrixLine); i++ {
		var line = matrixLine[i]
		var prevLine = matrixLine[i-1]
		if charEquals(line.chars, prevLine.chars) {
			if verifyRef(matrixLine, i) {
				return i
			}
		}
	}
	return -1
}

func charEquals(chars []string, chars2 []string) bool {
	var equals = true
	for i := range chars {
		if chars[i] != chars2[i] {
			equals = false
		}
	}
	return equals
}

func findReflexionWithError(lines []Line) int {
	var numberError = 0
	for i := 1; i < len(lines); i++ {
		numberError = 0
		for j := 0; j < len(lines[i].chars); j++ {
			if lines[i].chars[j] != lines[i-1].chars[j] {
				numberError++
			}
		}
		if numberError == 1 {
			//We only have one error, if we don't have any other error, that's the smudge and we have the reflexion
			if verifyWithError(lines, i, 0) {
				return i
			}
		} else if numberError == 0 {
			//There is one error to find. If we don't find it, we have a reflexion without smudge, so we return the index
			if verifyWithError(lines, i, 1) {
				return i
			}
		}
	}
	return -1
}

func verifyWithError(cols []Line, index int, errorNum int) bool {
	var numberError = 0
	var authorizedError = errorNum
	for i := 1; index+i < len(cols) && index-1-i >= 0; i++ {
		for j := 0; j < len(cols[i].chars); j++ {
			if cols[index+i].chars[j] != cols[index-i-1].chars[j] {
				numberError++
			}
		}
		if numberError > authorizedError {
			return false
		}
	}
	if numberError == authorizedError {
		return true
	} else {
		return false
	}
}
