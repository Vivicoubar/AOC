package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type neededTrio struct {
	red   int
	green int
	blue  int
}

func main() {
	fmt.Println("Ex1 :", Part1("AOC2023-2/ex.txt"))
	fmt.Println("Out1 :", Part1("AOC2023-2/input1.txt"))
	fmt.Println("Ex2 :", Part2("AOC2023-2/ex.txt"))
	fmt.Println("Out2 :", Part2("AOC2023-2/input1.txt"))
}

func Part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var red = 12
	var green = 13
	var blue = 14
	var sum = 0
	var possibleId = make([]int, 0)
	var powers = make([]int, 0)
	for scanner.Scan() {
		var currentNeed = neededTrio{0, 0, 0}
		line := scanner.Text()
		var stringElement = strings.Split(line, ":")
		var numGame = stringElement[0]
		numGame = strings.TrimSpace(numGame)
		numGame = strings.TrimPrefix(numGame, "Game ")
		var numGameId, _ = strconv.Atoi(numGame)
		stringElement = stringElement[1:]
		//strip the space
		var canGame = true
		stringElement[0] = strings.TrimSpace(stringElement[0])
		var games = strings.Split(stringElement[0], ";")
		for i := 0; i < len(games); i++ {
			games[i] = strings.TrimSpace(games[i])
		}
		for i := 0; i < len(games); i++ {

			var game = strings.Split(games[i], ",")
			var gameString = ""
			for j := 0; j < len(game); j++ {
				game[j] = strings.TrimSpace(game[j])
				gameString += game[j] + " "
			}
			gameString = strings.TrimSpace(gameString)
			red = 12
			green = 13
			blue = 14
			var element = strings.Split(gameString, " ")
			for k := 0; k < len(element); k++ {
				if element[k] == "red" && k != 0 {
					var val, _ = strconv.Atoi(element[k-1])
					red -= val
				}
				if element[k] == "green" && k != 0 {
					var val, _ = strconv.Atoi(element[k-1])
					green -= val
				}
				if element[k] == "blue" && k != 0 {
					var val, _ = strconv.Atoi(element[k-1])
					blue -= val
				}
				if !(red >= 0 && green >= 0 && blue >= 0) {
					canGame = false
				}
			}
			//Check neededCubes
			if red >= 0 {
				red = 12 - red
			}
			if green >= 0 {
				green = 13 - green
			}
			if blue >= 0 {
				blue = 14 - blue
			}
			if red < 0 {
				red = 12 - red
			}
			if green < 0 {
				green = 13 - green
			}
			if blue < 0 {
				blue = 14 - blue
			}

			red = int(math.Abs(float64(red)))
			green = int(math.Abs(float64(green)))
			blue = int(math.Abs(float64(blue)))
			if currentNeed.red < red {
				currentNeed.red = red
			}
			if currentNeed.green < green {
				currentNeed.green = green
			}
			if currentNeed.blue < blue {
				currentNeed.blue = blue
			}
		}
		if canGame {
			possibleId = append(possibleId, numGameId)
			sum += numGameId
		}
		powers = append(powers, currentNeed.red*currentNeed.blue*currentNeed.green)
	}
	var newSum = 0
	for i := 0; i < len(powers); i++ {
		newSum += powers[i]
	}
	return sum
}

func Part2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var red = 12
	var green = 13
	var blue = 14
	var sum = 0
	var possibleId = make([]int, 0)
	var powers = make([]int, 0)
	for scanner.Scan() {
		var currentNeed = neededTrio{0, 0, 0}
		line := scanner.Text()
		var stringElement = strings.Split(line, ":")
		var numGame = stringElement[0]
		numGame = strings.TrimSpace(numGame)
		numGame = strings.TrimPrefix(numGame, "Game ")
		var numGameId, _ = strconv.Atoi(numGame)
		stringElement = stringElement[1:]
		//strip the space
		var canGame = true
		stringElement[0] = strings.TrimSpace(stringElement[0])
		var games = strings.Split(stringElement[0], ";")
		for i := 0; i < len(games); i++ {
			games[i] = strings.TrimSpace(games[i])
		}
		for i := 0; i < len(games); i++ {

			var game = strings.Split(games[i], ",")
			var gameString = ""
			for j := 0; j < len(game); j++ {
				game[j] = strings.TrimSpace(game[j])
				gameString += game[j] + " "
			}
			gameString = strings.TrimSpace(gameString)
			red = 12
			green = 13
			blue = 14
			var element = strings.Split(gameString, " ")
			for k := 0; k < len(element); k++ {
				if element[k] == "red" && k != 0 {
					var val, _ = strconv.Atoi(element[k-1])
					red -= val
				}
				if element[k] == "green" && k != 0 {
					var val, _ = strconv.Atoi(element[k-1])
					green -= val
				}
				if element[k] == "blue" && k != 0 {
					var val, _ = strconv.Atoi(element[k-1])
					blue -= val
				}
				if !(red >= 0 && green >= 0 && blue >= 0) {
					canGame = false
				}
			}
			//Check neededCubes
			if red >= 0 {
				red = 12 - red
			}
			if green >= 0 {
				green = 13 - green
			}
			if blue >= 0 {
				blue = 14 - blue
			}
			if red < 0 {
				red = 12 - red
			}
			if green < 0 {
				green = 13 - green
			}
			if blue < 0 {
				blue = 14 - blue
			}

			red = int(math.Abs(float64(red)))
			green = int(math.Abs(float64(green)))
			blue = int(math.Abs(float64(blue)))
			if currentNeed.red < red {
				currentNeed.red = red
			}
			if currentNeed.green < green {
				currentNeed.green = green
			}
			if currentNeed.blue < blue {
				currentNeed.blue = blue
			}
		}
		if canGame {
			possibleId = append(possibleId, numGameId)
			sum += numGameId
		}
		powers = append(powers, currentNeed.red*currentNeed.blue*currentNeed.green)
	}
	var newSum = 0
	for i := 0; i < len(powers); i++ {
		newSum += powers[i]
	}

	return newSum
}
