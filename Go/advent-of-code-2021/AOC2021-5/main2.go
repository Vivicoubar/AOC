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
	var cloudMap = make([][]int, 0)
	for i := 0; i < 1000; i++ {
		var cloudLine = make([]int, 1000)
		cloudMap = append(cloudMap, cloudLine)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//line := scanner.Text()
		var lineLimits = strings.Split(scanner.Text(), " -> ")
		var x1, _ = strconv.Atoi(strings.Split(lineLimits[0], ",")[0])
		var y1, _ = strconv.Atoi(strings.Split(lineLimits[0], ",")[1])
		var x2, _ = strconv.Atoi(strings.Split(lineLimits[1], ",")[0])
		var y2, _ = strconv.Atoi(strings.Split(lineLimits[1], ",")[1])
		if x1 == x2 {
			cloudMap = processVerticalLine2(x1, y1, y2, cloudMap)
		} else if y1 == y2 {
			cloudMap = processHorizontalLine2(x1, x2, y1, cloudMap)
		} else {
			cloudMap = processDiagonal2(x1, y1, x2, y2, cloudMap)
		}
	}
	var count = 0
	for _, e := range cloudMap {
		for _, e2 := range e {
			if e2 >= 2 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}

func processHorizontalLine2(x1 int, x2 int, y1 int, cloudMap [][]int) [][]int {
	if x1 < x2 {
		for i := x1; i <= x2; i++ {
			cloudMap[y1][i] += 1
		}
	} else {
		for i := x2; i <= x1; i++ {
			cloudMap[y1][i] += 1
		}
	}
	return cloudMap
}

func processVerticalLine2(x1 int, y1 int, y2 int, cloudMap [][]int) [][]int {
	if y1 < y2 {
		for i := y1; i <= y2; i++ {
			cloudMap[i][x1] += 1
		}
	} else {
		for i := y2; i <= y1; i++ {
			cloudMap[i][x1] += 1
		}
	}
	return cloudMap
}
func processDiagonal2(x1 int, y1 int, x2 int, y2 int, cloudMap [][]int) [][]int {
	if x1 < x2 {
		for i := x1; i <= x2; i++ {
			cloudMap[y1][i] += 1
			if y1 < y2 {
				y1 += 1
			} else {
				y1 -= 1
			}

		}
	} else {
		for i := x2; i <= x1; i++ {
			cloudMap[y2][i] += 1
			if y2 < y1 {
				y2 += 1
			} else {
				y2 -= 1
			}
		}
	}
	return cloudMap
}
