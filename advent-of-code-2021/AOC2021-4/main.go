package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Ouvrir le fichier en lecture
	file, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()
	// Lire le fichier ligne par ligne
	// La première ligne contient les différentes valeurs
	var isFirst = true
	var nums []int = make([]int, 0)
	var matrixSize int
	var count int
	var numMatrix = 0
	var matrix [][]int = make([][]int, 0)
	var lineColumns [][]int = make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isFirst {
			isFirst = false
			//length = len(line)
			//Split the first line
			splitLine := strings.Split(line, ",")
			for _, char := range splitLine {
				//convert string to int
				elem, _ := strconv.Atoi(char)
				nums = append(nums, elem)
			}
			matrixSize = 5
			count = -1
			continue
		}
		//READ THE MATRIX
		if count == matrixSize {
			count = 0
			lineColumns = addMatrixToLineAndColumns2(matrix, lineColumns, numMatrix, matrixSize)
			numMatrix += 1
			matrix = make([][]int, 0)
			continue
		} else {
			if count != -1 {
				var matrixLine = make([]int, 0)
				var previousChar = "-1"
				for _, char := range line {
					if string(char) == " " {
						if previousChar == "-1" {
							continue
						} else {
							elem, _ := strconv.Atoi(previousChar)
							matrixLine = append(matrixLine, elem)
							previousChar = "-1"
						}

					} else {
						if previousChar == "-1" {
							previousChar = ""
						}
						previousChar = previousChar + string(char)
					}

				}
				elem, _ := strconv.Atoi(previousChar)
				matrixLine = append(matrixLine, elem)
				matrix = append(matrix, matrixLine)
			}
			count++
		}

	}
	lineColumns = addMatrixToLineAndColumns2(matrix, lineColumns, numMatrix, matrixSize)
	var posWinner int = -1
	for _, num := range nums {
		posWinner, lineColumns = processNumber1(num, lineColumns)
		if posWinner != -1 {
			var out = (sumUnMarked1(posWinner, lineColumns)) * num
			fmt.Println(posWinner, num, out)
			break
		}
	}
}

func sumUnMarked1(winner int, lines [][]int) int {
	var pos = winner * 10
	var sum = 0
	for i := pos; i < pos+5; i++ {
		if lines[i][0] == winner {
			for j := 1; j < len(lines[i]); j++ {
				if lines[i][j] != -1 {
					sum += lines[i][j]
				}
			}

		}
	}
	return sum
}

func processNumber1(num int, lineColumns [][]int) (int, [][]int) {
	for i, line := range lineColumns {
		var countForLine = 0
		for j := 1; j < len(line); j++ {
			if lineColumns[i][j] != -1 {
				countForLine += 1
			}
			if lineColumns[i][j] == num {
				lineColumns[i][j] = -1
				countForLine -= 1
			}
		}
		if countForLine == 0 {
			return lineColumns[i][0], lineColumns
		}
	}
	return -1, lineColumns
}

func addMatrixToLineAndColumns2(matrix [][]int, lineColumns [][]int, matrixCount int, matrixsize int) [][]int {
	//Create lines
	for _, line := range matrix {
		var newLine = make([]int, 0)
		newLine = append(newLine, matrixCount)
		for _, elem := range line {
			newLine = append(newLine, elem)
		}
		lineColumns = append(lineColumns, newLine)
	}

	for i := 0; i < matrixsize; i++ {
		var newColumn = make([]int, 0)
		newColumn = append(newColumn, matrixCount)
		for j := 0; j < matrixsize; j++ {
			newColumn = append(newColumn, matrix[j][i])
		}
		lineColumns = append(lineColumns, newColumn)
	}
	return lineColumns
}
