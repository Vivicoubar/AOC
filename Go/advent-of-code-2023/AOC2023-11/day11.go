package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type galaxyPos struct {
	x, y float64
}

func main() {
	var input = "input1.txt"
	var ex = "ex.txt"
	fmt.Println("Part1 Ex:", P1(ex))
	fmt.Println("Part1 :", P1(input))
	fmt.Println("Part2 :", P2(input, 1e6))
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
	var grid = make([][]rune, 0)
	var galaxyList = make([]galaxyPos, 0)
	var lineToDouble = make([]float64, 0)
	var columnToDouble = make([]float64, 0)
	var countY float64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		var hasGalaxy = false
		var curLine = make([]rune, 0)
		for i, c := range line {
			curLine = append(curLine, c)
			if c == '#' {
				hasGalaxy = true
				galaxyList = append(galaxyList, galaxyPos{float64(i), countY})
			}
		}
		if !hasGalaxy {
			lineToDouble = append(lineToDouble, countY)
		}
		grid = append(grid, curLine)
		countY++
	}
	for j := 0; j < len(grid[0]); j++ {
		var hasGalaxy = false
		for _, line := range grid {
			if line[j] == '#' {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			columnToDouble = append(columnToDouble, float64(j))
		}
	}
	var num float64 = 1
	for j, galaxy := range galaxyList {
		var newY = galaxyList[j].y
		for _, line := range lineToDouble {
			if galaxy.y > line {
				newY += num
			}
		}
		galaxyList[j].y = newY
	}
	for j, galaxy := range galaxyList {
		var newX = galaxyList[j].x
		for _, column := range columnToDouble {
			if galaxy.x > column {
				newX += num
			}
		}
		galaxyList[j].x = newX
	}
	var distance = make([]int, 0)
	for i, galaxy1 := range galaxyList {
		for j, galaxy2 := range galaxyList {
			if i < j {
				var distY = int(math.Abs(float64(galaxy1.y - galaxy2.y)))
				var distX = int(math.Abs(float64(galaxy1.x - galaxy2.x)))
				distance = append(distance, distX+distY)
			}
		}
	}
	var sum = 0
	for _, dist := range distance {
		sum += dist
	}
	return sum
}

func P2(input string, num float64) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var grid = make([][]rune, 0)
	var galaxyList = make([]galaxyPos, 0)
	var lineToDouble = make([]float64, 0)
	var columnToDouble = make([]float64, 0)
	var countY float64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		var hasGalaxy = false
		var curLine = make([]rune, 0)
		for i, c := range line {
			curLine = append(curLine, c)
			if c == '#' {
				hasGalaxy = true
				galaxyList = append(galaxyList, galaxyPos{float64(i), countY})
			}
		}
		if !hasGalaxy {
			lineToDouble = append(lineToDouble, countY)
		}
		grid = append(grid, curLine)
		countY++
	}
	for j := 0; j < len(grid[0]); j++ {
		var hasGalaxy = false
		for _, line := range grid {
			if line[j] == '#' {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			columnToDouble = append(columnToDouble, float64(j))
		}
	}
	for j, galaxy := range galaxyList {
		var newY = galaxyList[j].y
		for _, line := range lineToDouble {
			if galaxy.y > line {
				newY += num - 1
			}
		}
		galaxyList[j].y = newY
	}
	for j, galaxy := range galaxyList {
		var newX = galaxyList[j].x
		for _, column := range columnToDouble {
			if galaxy.x > column {
				newX += num - 1
			}
		}
		galaxyList[j].x = newX
	}
	var distance = make([]int, 0)
	for i, galaxy1 := range galaxyList {
		for j, galaxy2 := range galaxyList {
			if i < j {
				var distY = int(math.Abs(float64(galaxy1.y - galaxy2.y)))
				var distX = int(math.Abs(float64(galaxy1.x - galaxy2.x)))
				distance = append(distance, distX+distY)
			}
		}
	}
	var sum = 0
	for _, dist := range distance {
		sum += dist
	}
	return sum
}
