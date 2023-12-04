package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// Ouvrir le fichier en lecture
	file, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()
	var input = make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	var oldO2 = input
	var oldCO2 = input
	var O2 float64 = 0
	var CO2 float64 = 0
	//for i in range 0 to len of the first element of input
	for i := 0; i < len(input[0]); i++ {
		var newO2 = findBitCriteria(oldO2, i, false)
		if len(newO2) == 1 {
			for j := 0; j < len(newO2[0]); j++ {
				if newO2[0][j] == '1' {
					O2 += math.Pow(2, float64(len(newO2[0])-1-j))
				}
			}
			break
		}
		oldO2 = newO2

	}
	for i := 0; i < len(input[0]); i++ {
		var newCO2 = findBitCriteria(oldCO2, i, true)
		if len(newCO2) == 1 {
			for j := 0; j < len(newCO2[0]); j++ {
				if newCO2[0][j] == '1' {
					CO2 += math.Pow(2, float64(len(newCO2[0])-1-j))
				}
			}
			break
		}
		oldCO2 = newCO2
	}
	fmt.Println(O2*CO2, O2, CO2)
}

func findBitCriteria(input []string, pos int, fewer bool) []string {
	var score = 0
	var bitValue = 0
	for i := 0; i < len(input); i++ {
		if input[i][pos] == '1' {
			score++
		} else {
			score--
		}
	}
	if score > 0 && !fewer {
		bitValue = 1
	}
	if score < 0 && fewer {
		bitValue = 1
	}
	if !fewer && score == 0 {
		bitValue = 1
	}
	var newInput = make([]string, 0)
	for i := 0; i < len(input); i++ {
		var s = string(input[i][pos])
		if s == strconv.Itoa(int(bitValue)) {
			newInput = append(newInput, input[i])
		}
	}
	return newInput
}
