package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type gearPos struct {
	row int
	col int
}
type gearInfo struct {
	num1     int
	num2     int
	numFound int
}

func main() {
	// Ouvrir le fichier en lecture
	fmt.Println("Ex1 :", Part1("AOC2023-3/ex.txt"))
	fmt.Println("Out1 :", Part1("AOC2023-3/input1.txt"))
	fmt.Println("Ex2 :", Part2("AOC2023-3/ex.txt"))
	fmt.Println("Out2 :", Part2("AOC2023-3/input1.txt"))
}

func Part1(filePath string) int {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for i, c := range line {
			row[i] = string(c)
		}
		grid = append(grid, row)
	}

	var sum = 0
	for i := 0; i < len(grid); i++ {
		var curNum = 0
		var curPos = make([]int, 0)
		var isPart = false
		for j := 0; j < len(grid[0]); j++ {
			if isNumber(grid[i][j]) {
				var element, _ = strconv.Atoi(grid[i][j])
				curNum = 10*curNum + element
				curPos = append(curPos, j)
				var foundPart = false

				foundPart, _, _ = checkPart(grid, curPos, i)
				if foundPart {
					isPart = true
				}
			} else {
				if isPart {
					sum += curNum
				}
				curNum = 0
				curPos = make([]int, 0)
				isPart = false
			}
		}
		// IT IS IMPORTANT TO CHECK AGAIN AT THE END OF THE ROW BECAUSE THE LAST NUMBER MIGHT BE A PART
		if isPart {
			sum += curNum
			curNum = 0
			curPos = make([]int, 0)
			isPart = false
		}

	}
	return sum
}

func Part2(filePath string) int {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for i, c := range line {
			row[i] = string(c)
		}
		grid = append(grid, row)
	}

	var sum = 0
	var gears = make(map[gearPos]gearInfo)
	for i := 0; i < len(grid); i++ {
		var curNum = 0
		var curPos = make([]int, 0)
		var isPart = false
		var foundChar = ""
		var foundCharPos = gearPos{}
		for j := 0; j < len(grid[0]); j++ {
			if isNumber(grid[i][j]) {
				var element, _ = strconv.Atoi(grid[i][j])
				curNum = 10*curNum + element
				curPos = append(curPos, j)
				var foundPart = false

				foundPart, foundChar, foundCharPos = checkPart(grid, curPos, i)
				if foundPart {
					isPart = true
				}
			} else {
				if isPart {
					sum += curNum
					if foundChar == "*" {
						if foundCharPos.col != -1 && foundCharPos.row != -1 { //Might be always true
							var value, exists = gears[foundCharPos]
							if exists {
								gears[foundCharPos] = gearInfo{num1: value.num1, num2: curNum, numFound: value.numFound + 1}
							} else {
								gears[foundCharPos] = gearInfo{num1: curNum, numFound: 1}
							}
						}
					}
				}
				curNum = 0
				curPos = make([]int, 0)
				isPart = false
			}
		}
		// IT IS IMPORTANT TO CHECK AGAIN AT THE END OF THE ROW BECAUSE THE LAST NUMBER MIGHT BE A PART
		if isPart {
			sum += curNum
			if foundChar == "*" {
				if foundCharPos.col != -1 && foundCharPos.row != -1 { //Might be always true
					var value, exists = gears[foundCharPos]
					if exists {
						gears[foundCharPos] = gearInfo{num1: value.num1, num2: curNum, numFound: value.numFound + 1}
					} else {
						gears[foundCharPos] = gearInfo{num1: curNum, numFound: 1}
					}
				}
			}
			curNum = 0
			curPos = make([]int, 0)
			isPart = false
		}

	}
	var sum2 = 0
	for _, value := range gears {
		if value.numFound == 2 {
			sum2 += value.num1 * value.num2
		}
	}
	return sum2
}

func isNumber(char string) bool {
	if char == "0" || char == "1" || char == "2" || char == "3" || char == "4" || char == "5" || char == "6" || char == "7" || char == "8" || char == "9" {
		return true
	}
	return false
}

func checkPart(grid [][]string, curPos []int, row int) (bool, string, gearPos) {
	for i := row - 1; i <= row+1; i++ {
		var firstPos = curPos[0]
		var secondPos = curPos[len(curPos)-1]
		for j := firstPos - 1; j <= secondPos+1; j++ {
			if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
				if grid[i][j] != "." && !isNumber(grid[i][j]) {
					return true, grid[i][j], gearPos{i, j}
				}
			}
		}
	}
	return false, "", gearPos{-1, -1}
}
