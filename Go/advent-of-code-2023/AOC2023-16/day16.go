package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	row int
	col int
	dir string
}

func main() {
	var ex = "AOC2023-16/ex.txt"
	var input = "AOC2023-16/input1.txt"
	fmt.Println("Partie 1 exemple: ", P1(ex))
	fmt.Println("Partie 1 : ", P1(input))
	fmt.Println("Partie 2 exemple: ", P2(ex))
	fmt.Println("Partie 2 : ", P2(input))
}

func P1(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix = make([][]string, 0)
	var heatMatrix = make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var row = make([]string, 0)
		var rowInt = make([]int, 0)
		for _, char := range line {
			row = append(row, string(char))
			rowInt = append(rowInt, 0)
		}
		matrix = append(matrix, row)
		heatMatrix = append(heatMatrix, rowInt)
	}
	var beamPos = make([]Pos, 0)
	if matrix[0][0] == "/" {
		beamPos = append(beamPos, Pos{row: 0, col: 0, dir: "up"})
	} else if matrix[0][0] == "\\" {
		beamPos = append(beamPos, Pos{row: 0, col: 0, dir: "down"})
	} else if matrix[0][0] == "|" {
		beamPos = append(beamPos, Pos{row: 0, col: 0, dir: "down"})
	} else {
		beamPos = append(beamPos, Pos{row: 0, col: 0, dir: "right"})
	}
	for len(beamPos) != 0 {
		var curPos = beamPos[0]
		beamPos = beamPos[1:]
		if curPos.row >= len(matrix) || curPos.col >= len(matrix[0]) || curPos.row < 0 || curPos.col < 0 {
			continue
		}
		heatMatrix[curPos.row][curPos.col]++
		if heatMatrix[curPos.row][curPos.col] > 7 { //Not the best, change it to a var that represent the direction that have been explored
			continue
		}
		if curPos.dir == "right" {
			if curPos.col+1 >= len(matrix[0]) {
				continue
			}
			if matrix[curPos.row][curPos.col+1] == "." || matrix[curPos.row][curPos.col+1] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "right"})
			} else if matrix[curPos.row][curPos.col+1] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "up"})
			} else if matrix[curPos.row][curPos.col+1] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "down"})
			} else if matrix[curPos.row][curPos.col+1] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "up"})
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "down"})
			}
		} else if curPos.dir == "up" {
			if curPos.row-1 < 0 {
				continue
			}
			if matrix[curPos.row-1][curPos.col] == "." || matrix[curPos.row-1][curPos.col] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "up"})
			} else if matrix[curPos.row-1][curPos.col] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "right"})
			} else if matrix[curPos.row-1][curPos.col] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "left"})
			} else if matrix[curPos.row-1][curPos.col] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "right"})
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "left"})
			}
		} else if curPos.dir == "left" {
			if curPos.col-1 < 0 {
				continue
			}
			if matrix[curPos.row][curPos.col-1] == "." || matrix[curPos.row][curPos.col-1] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "left"})
			} else if matrix[curPos.row][curPos.col-1] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "down"})
			} else if matrix[curPos.row][curPos.col-1] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "up"})
			} else if matrix[curPos.row][curPos.col-1] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "up"})
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "down"})
			}
		} else if curPos.dir == "down" {
			if curPos.row+1 >= len(matrix) {
				continue
			}
			if matrix[curPos.row+1][curPos.col] == "." || matrix[curPos.row+1][curPos.col] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "down"})
			} else if matrix[curPos.row+1][curPos.col] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "left"})
			} else if matrix[curPos.row+1][curPos.col] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "right"})
			} else if matrix[curPos.row+1][curPos.col] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "right"})
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "left"})
			}
		}
	}
	var p1 = countHeat(heatMatrix)
	//write in output.txt the heat matrix
	file, err = os.Create("output.txt")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return 0
	}
	defer file.Close()
	for _, row := range heatMatrix {
		for _, heat := range row {
			file.WriteString(fmt.Sprintf("%d", heat))
		}
		file.WriteString("\n")
	}
	return p1
}

func countHeat(matrix [][]int) int {
	var count = 0
	for _, row := range matrix {
		for _, heat := range row {
			if heat > 0 {
				count++
			}
		}
	}
	return count
}

func P2(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix = make([][]string, 0)
	var heatMatrix = make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var row = make([]string, 0)
		var rowInt = make([]int, 0)
		for _, char := range line {
			row = append(row, string(char))
			rowInt = append(rowInt, 0)
		}
		matrix = append(matrix, row)
		heatMatrix = append(heatMatrix, rowInt)
	}
	var beamPos = make([]Pos, 0)
	if matrix[0][0] == "/" {
		beamPos = append(beamPos, Pos{row: 0, col: 0, dir: "up"})
	} else if matrix[0][0] == "\\" {
		beamPos = append(beamPos, Pos{row: 0, col: 0, dir: "down"})
	} else if matrix[0][0] == "|" {
		beamPos = append(beamPos, Pos{row: 0, col: 0, dir: "down"})
	} else {
		beamPos = append(beamPos, Pos{row: 0, col: 0, dir: "right"})
	}
	for len(beamPos) != 0 {
		var curPos = beamPos[0]
		beamPos = beamPos[1:]
		if curPos.row >= len(matrix) || curPos.col >= len(matrix[0]) || curPos.row < 0 || curPos.col < 0 {
			continue
		}
		heatMatrix[curPos.row][curPos.col]++
		if heatMatrix[curPos.row][curPos.col] > 7 {
			continue
		}
		if curPos.dir == "right" {
			if curPos.col+1 >= len(matrix[0]) {
				continue
			}
			if matrix[curPos.row][curPos.col+1] == "." || matrix[curPos.row][curPos.col+1] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "right"})
			} else if matrix[curPos.row][curPos.col+1] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "up"})
			} else if matrix[curPos.row][curPos.col+1] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "down"})
			} else if matrix[curPos.row][curPos.col+1] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "up"})
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "down"})
			}
		} else if curPos.dir == "up" {
			if curPos.row-1 < 0 {
				continue
			}
			if matrix[curPos.row-1][curPos.col] == "." || matrix[curPos.row-1][curPos.col] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "up"})
			} else if matrix[curPos.row-1][curPos.col] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "right"})
			} else if matrix[curPos.row-1][curPos.col] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "left"})
			} else if matrix[curPos.row-1][curPos.col] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "right"})
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "left"})
			}
		} else if curPos.dir == "left" {
			if curPos.col-1 < 0 {
				continue
			}
			if matrix[curPos.row][curPos.col-1] == "." || matrix[curPos.row][curPos.col-1] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "left"})
			} else if matrix[curPos.row][curPos.col-1] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "down"})
			} else if matrix[curPos.row][curPos.col-1] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "up"})
			} else if matrix[curPos.row][curPos.col-1] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "up"})
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "down"})
			}
		} else if curPos.dir == "down" {
			if curPos.row+1 >= len(matrix) {
				continue
			}
			if matrix[curPos.row+1][curPos.col] == "." || matrix[curPos.row+1][curPos.col] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "down"})
			} else if matrix[curPos.row+1][curPos.col] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "left"})
			} else if matrix[curPos.row+1][curPos.col] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "right"})
			} else if matrix[curPos.row+1][curPos.col] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "right"})
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "left"})
			}
		}
	}
	var startPoss = make([][]Pos, 0)
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == "/" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: i, col: 0, dir: "up"})
			startPoss = append(startPoss, startBeam)
		} else if matrix[i][0] == "\\" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: i, col: 0, dir: "down"})
			startPoss = append(startPoss, startBeam)
		} else if matrix[i][0] == "|" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: i, col: 0, dir: "down"})
			startBeam = append(startBeam, Pos{row: i, col: 0, dir: "up"})
			startPoss = append(startPoss, startBeam)
		} else {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: i, col: 0, dir: "right"})
			startPoss = append(startPoss, startBeam)
		}
		if matrix[i][len(matrix[0])-1] == "/" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: i, col: len(matrix[0]) - 1, dir: "down"})
			startPoss = append(startPoss, startBeam)
		} else if matrix[i][len(matrix[0])-1] == "\\" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: i, col: len(matrix[0]) - 1, dir: "up"})
			startPoss = append(startPoss, startBeam)
		} else if matrix[i][len(matrix[0])-1] == "|" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: i, col: len(matrix[0]) - 1, dir: "down"})
			startBeam = append(startBeam, Pos{row: i, col: len(matrix[0]) - 1, dir: "up"})
			startPoss = append(startPoss, startBeam)
		} else {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: i, col: len(matrix[0]) - 1, dir: "left"})
			startPoss = append(startPoss, startBeam)
		}
	}
	for j := 1; j < len(matrix[0])-1; j++ {
		if matrix[0][j] == "/" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: 0, col: j, dir: "left"})
			startPoss = append(startPoss, startBeam)
		} else if matrix[0][j] == "\\" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: 0, col: j, dir: "right"})
			startPoss = append(startPoss, startBeam)
		} else if matrix[0][j] == "-" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: 0, col: j, dir: "right"})
			startBeam = append(startBeam, Pos{row: 0, col: j, dir: "left"})
			startPoss = append(startPoss, startBeam)
		} else {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: 0, col: j, dir: "down"})
			startPoss = append(startPoss, startBeam)
		}
		if matrix[len(matrix)-1][j] == "/" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: len(matrix) - 1, col: j, dir: "right"})
			startPoss = append(startPoss, startBeam)
		} else if matrix[len(matrix)-1][j] == "\\" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: len(matrix) - 1, col: j, dir: "left"})
			startPoss = append(startPoss, startBeam)
		} else if matrix[len(matrix)-1][j] == "-" {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: len(matrix) - 1, col: j, dir: "right"})
			startBeam = append(startBeam, Pos{row: len(matrix) - 1, col: j, dir: "left"})
			startPoss = append(startPoss, startBeam)
		} else {
			var startBeam = make([]Pos, 0)
			startBeam = append(startBeam, Pos{row: len(matrix) - 1, col: j, dir: "up"})
			startPoss = append(startPoss, startBeam)
		}

	}

	var p2 = 0
	for _, startPos := range startPoss {
		var curCount = startPosCountHeat(startPos, matrix)
		if curCount > p2 {
			p2 = curCount
		}
	}
	return p2
}

func startPosCountHeat(startPos []Pos, matrix [][]string) int {
	var beamPos = make([]Pos, 0)
	beamPos = append(beamPos, startPos...)
	var heatMatrix = make([][]int, 0)
	for _, row := range matrix {
		var rowInt = make([]int, 0)
		for _, _ = range row {
			rowInt = append(rowInt, 0)
		}
		heatMatrix = append(heatMatrix, rowInt)
	}

	for len(beamPos) != 0 {
		var curPos = beamPos[0]
		beamPos = beamPos[1:]
		if curPos.row >= len(matrix) || curPos.col >= len(matrix[0]) || curPos.row < 0 || curPos.col < 0 {
			continue
		}
		heatMatrix[curPos.row][curPos.col]++
		if heatMatrix[curPos.row][curPos.col] > 7 {
			continue
		}
		if curPos.dir == "right" {
			if curPos.col+1 >= len(matrix[0]) {
				continue
			}
			if matrix[curPos.row][curPos.col+1] == "." || matrix[curPos.row][curPos.col+1] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "right"})
			} else if matrix[curPos.row][curPos.col+1] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "up"})
			} else if matrix[curPos.row][curPos.col+1] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "down"})
			} else if matrix[curPos.row][curPos.col+1] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "up"})
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col + 1, dir: "down"})
			}
		} else if curPos.dir == "up" {
			if curPos.row-1 < 0 {
				continue
			}
			if matrix[curPos.row-1][curPos.col] == "." || matrix[curPos.row-1][curPos.col] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "up"})
			} else if matrix[curPos.row-1][curPos.col] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "right"})
			} else if matrix[curPos.row-1][curPos.col] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "left"})
			} else if matrix[curPos.row-1][curPos.col] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "right"})
				beamPos = append(beamPos, Pos{row: curPos.row - 1, col: curPos.col, dir: "left"})
			}
		} else if curPos.dir == "left" {
			if curPos.col-1 < 0 {
				continue
			}
			if matrix[curPos.row][curPos.col-1] == "." || matrix[curPos.row][curPos.col-1] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "left"})
			} else if matrix[curPos.row][curPos.col-1] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "down"})
			} else if matrix[curPos.row][curPos.col-1] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "up"})
			} else if matrix[curPos.row][curPos.col-1] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "up"})
				beamPos = append(beamPos, Pos{row: curPos.row, col: curPos.col - 1, dir: "down"})
			}
		} else if curPos.dir == "down" {
			if curPos.row+1 >= len(matrix) {
				continue
			}
			if matrix[curPos.row+1][curPos.col] == "." || matrix[curPos.row+1][curPos.col] == "|" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "down"})
			} else if matrix[curPos.row+1][curPos.col] == "/" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "left"})
			} else if matrix[curPos.row+1][curPos.col] == "\\" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "right"})
			} else if matrix[curPos.row+1][curPos.col] == "-" {
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "right"})
				beamPos = append(beamPos, Pos{row: curPos.row + 1, col: curPos.col, dir: "left"})
			}
		}
	}
	return countHeat(heatMatrix)
}
