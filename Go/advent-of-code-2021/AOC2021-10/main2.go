package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pairs struct {
	i int
	j int
}

func main() {
	// Ouvrir le fichier en lecture
	file, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matrix [][]int = make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var numberString = strings.Split(line, "")
		var numberInt []int = make([]int, 0)
		for _, numberW := range numberString {
			var number, _ = strconv.Atoi(numberW)
			numberInt = append(numberInt, number)
		}
		matrix = append(matrix, numberInt)
	}
	//Create an array of pairs
	var lowPoints = make([]pairs, 0)
	var bassinSize = make([]int, 0)
	var curResult = 0
	var n = len(matrix)
	var m = len(matrix[0])
	var sum = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if check2(matrix, i, j) {
				sum += matrix[i][j] + 1
				lowPoints = append(lowPoints, pairs{i, j})
				curResult, _ = findBassinSize(matrix, i, j, make([]pairs, 0))
				bassinSize = append(bassinSize, curResult)
			}
		}
	}
	//sort bassin size
	for i := 0; i < len(bassinSize); i++ {
		for j := i + 1; j < len(bassinSize); j++ {
			if bassinSize[i] > bassinSize[j] {
				var temp = bassinSize[i]
				bassinSize[i] = bassinSize[j]
				bassinSize[j] = temp
			}
		}
	}
	fmt.Println(bassinSize[len(bassinSize)-1] * bassinSize[len(bassinSize)-2] * bassinSize[len(bassinSize)-3])
	fmt.Println(bassinSize)
	fmt.Println(sum)
}

func findBassinSize(matrix [][]int, i int, j int, found []pairs) (int, []pairs) {
	//IF its 9 or already in, ignore
	for _, e := range found {
		if e.i == i && e.j == j {
			return 0, found
		}
	}
	if i < 0 || i > len(matrix)-1 || j < 0 || j > len(matrix[0]) {
		return 0, found
	}
	if matrix[i][j] == 9 {
		return 0, found
	}
	found = append(found, pairs{i, j})
	var size = 1
	var curSize = 0
	if checkGreater(matrix, i-1, j) {
		curSize, found = findBassinSize(matrix, i-1, j, found)
		size += curSize
	}
	if checkGreater(matrix, i+1, j) {
		curSize, found = findBassinSize(matrix, i+1, j, found)
		size += curSize
	}
	if checkGreater(matrix, i, j-1) {
		curSize, found = findBassinSize(matrix, i, j-1, found)
		size += curSize
	}
	if checkGreater(matrix, i, j+1) {
		curSize, found = findBassinSize(matrix, i, j+1, found)
		size += curSize
	}
	return size, found

}

func checkGreater(matrix [][]int, i int, j int) bool {
	if i < 0 || i > len(matrix)-1 || j < 0 || j > len(matrix[0])-1 {
		return false
	}
	if matrix[i][j] == 9 {
		return false
	}
	return true
}

func check2(matrix [][]int, i int, j int) bool {
	var n = len(matrix)
	var m = len(matrix[0])
	var isMin = true
	if i > 0 {
		if matrix[i][j] >= matrix[i-1][j] {
			isMin = false
		}
	}
	if i < n-1 {
		if matrix[i][j] >= matrix[i+1][j] {
			isMin = false
		}
	}
	if j > 0 {
		if matrix[i][j] >= matrix[i][j-1] {
			isMin = false
		}
	}
	if j < m-1 {
		if matrix[i][j] >= matrix[i][j+1] {
			isMin = false
		}
	}
	return isMin
}
