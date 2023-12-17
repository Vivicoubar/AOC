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
	fmt.Println("Ex1 :", Part1("ex.txt"))
	fmt.Println("Out1 :", Part1("input1.txt"))
	fmt.Println("Ex2 :", Part2("ex2.txt"))
	fmt.Println("Out2 :", Part2("input1.txt"))
}

func Part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var sum = 0
	for scanner.Scan() {
		line := scanner.Text()
		var lineString []string
		//Enlever les lettres de la ligne
		var toRemove = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		for i := 0; i < len(toRemove); i++ {
			var letter = toRemove[i]
			line = strings.ReplaceAll(line, letter, "")
		}
		lineString = strings.Split(line, "")
		//Récupérer les 2 derniers caractères
		var value1, _ = strconv.Atoi(lineString[0])
		var value2, _ = strconv.Atoi(lineString[len(lineString)-1])
		sum += value1*10 + value2
	}
	return sum
}

func Part2(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var sum = 0
	for scanner.Scan() {
		line := scanner.Text()
		var lineString []string
		var toRemove = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		var curLine = ""
		for i := 0; i < len(line); i++ {
			//WATCHOUT FOR "one1twoone" => "1121"
			curLine += string(line[i])
			curLine = strings.ReplaceAll(curLine, "one", "1e")
			curLine = strings.ReplaceAll(curLine, "two", "2o")
			curLine = strings.ReplaceAll(curLine, "three", "3e")
			curLine = strings.ReplaceAll(curLine, "four", "4")
			curLine = strings.ReplaceAll(curLine, "five", "5e")
			curLine = strings.ReplaceAll(curLine, "six", "6")
			curLine = strings.ReplaceAll(curLine, "seven", "7n")
			curLine = strings.ReplaceAll(curLine, "eight", "8t")
			curLine = strings.ReplaceAll(curLine, "nine", "9e")
		}
		line = curLine
		for i := 0; i < len(toRemove); i++ {
			var letter = toRemove[i]
			line = strings.ReplaceAll(line, letter, "")
		}
		lineString = strings.Split(line, "")
		var value1, _ = strconv.Atoi(lineString[0])
		var value2, _ = strconv.Atoi(lineString[len(lineString)-1])
		sum += value1*10 + value2
	}
	return sum
}
