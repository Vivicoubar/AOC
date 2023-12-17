package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	isFirst := true
	var length = 0
	//Créé un tableau d'entiers sans l'initialiser, de taille 8
	var expression []int
	var gamma float64
	var epsilon float64
	for scanner.Scan() {
		line := scanner.Text()
		if isFirst {
			isFirst = false
			length = len(line)
			expression = make([]int, length)
		}
		for i, char := range line {
			if char == '1' {
				expression[i] += 1
			} else {
				expression[i] -= 1
			}
		}
	}
	var gammaList = make([]int, length)
	var epsilonList = make([]int, length)
	//for each element in expression,
	for i, val := range expression {
		b := val > 0
		if b {
			gammaList[i] = 1
			epsilonList[i] = 0
			gamma += math.Pow(2, float64(length-1-i))
		} else {
			gammaList[i] = 0
			epsilonList[i] = 1
			epsilon += math.Pow(2, float64(length-1-i))
		}
	}

	fmt.Println(gamma*epsilon, gamma, epsilon)
	fmt.Println(gammaList, epsilonList)
}
