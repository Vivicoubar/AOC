package main

import (
	"bufio"
	"fmt"
	"math"
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
	var minFuel float64 = 999999999999999
	var maxPos = 2
	var initPos []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var stringVal = strings.Split(line, ",")
		for i := 0; i < len(stringVal); i++ {
			var number, _ = strconv.Atoi(stringVal[i])
			initPos = append(initPos, number)
			if maxPos < number {
				maxPos = number
			}
		}
	}
	for i := 0; i < maxPos; i++ {
		var sum float64 = 0
		for j := 0; j < len(initPos); j++ {
			var perSoSum float64 = 0
			var k float64 = 0
			for k = 0; k < math.Abs(float64(initPos[j]-i)); k++ {
				perSoSum += k + 1
			}
			sum += perSoSum
		}
		if sum < minFuel {
			minFuel = sum
		}
	}
	fmt.Println(minFuel)
}
