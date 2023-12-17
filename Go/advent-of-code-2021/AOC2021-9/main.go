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
	var n = len(matrix)
	var m = len(matrix[0])
	var sum = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if check(matrix, i, j) {
				sum += matrix[i][j] + 1
				fmt.Println(i, j, matrix[i][j])
			}
		}
	}
	fmt.Println(sum)
}

func check(matrix [][]int, i int, j int) bool {
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
