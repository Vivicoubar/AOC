package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Ex1 :", Part1("AOC2023-4/ex.txt"))
	fmt.Println("Out1 :", Part1("AOC2023-4/input1.txt"))
	fmt.Println("Ex2 :", Part2("AOC2023-4/ex.txt"))
	fmt.Println("Out2 :", Part2("AOC2023-4/input1.txt"))
}

func Part1(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var totalScore = 0
	var copyList = make([]int, 100000) //A BIT UGLY BUT WORKS
	for i := 0; i < len(copyList); i++ {
		copyList[i] = 1
	}
	copyList[0] = 0
	var currentCard = 0
	var matchingNum = 0
	for scanner.Scan() {
		currentCard++
		line := scanner.Text()
		var numbers = strings.Split(line, ": ")
		numbers = numbers[1:]
		numbers = strings.Split(numbers[0], "|")
		for i := 0; i < len(numbers); i++ {
			numbers[i] = strings.TrimSpace(numbers[i])
			//replace double spaces by single space
			numbers[i] = strings.Replace(numbers[i], "  ", " ", -1)
		}
		var winningNumbers = strings.Split(numbers[0], " ")

		var myNumbers = strings.Split(numbers[1], " ")
		var score = 0
		matchingNum = 0
		for i := 0; i < len(winningNumbers); i++ {
			for j := 0; j < len(myNumbers); j++ {
				if winningNumbers[i] == myNumbers[j] {
					if score == 0 {
						score = 1
					} else {
						score = score * 2
					}
					matchingNum++
				}
			}
		}
		for i := 0; i < matchingNum; i++ {
			if currentCard+i < 209 {
				copyList[currentCard+1+i] += 1 * copyList[currentCard]
			}
		}
		totalScore = totalScore + score
	}
	var sum2 = 0
	for i := 0; i < currentCard+1; i++ {
		sum2 += copyList[i]
	}
	return totalScore
}

func Part2(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var totalScore = 0
	var copyList = make([]int, 100000) //A BIT UGLY BUT WORKS
	for i := 0; i < len(copyList); i++ {
		copyList[i] = 1
	}
	copyList[0] = 0
	var currentCard = 0
	var matchingNum = 0
	for scanner.Scan() {
		currentCard++
		line := scanner.Text()
		var numbers = strings.Split(line, ": ")
		numbers = numbers[1:]
		numbers = strings.Split(numbers[0], "|")
		for i := 0; i < len(numbers); i++ {
			numbers[i] = strings.TrimSpace(numbers[i])
			//replace double spaces by single space
			numbers[i] = strings.Replace(numbers[i], "  ", " ", -1)
		}
		var winningNumbers = strings.Split(numbers[0], " ")

		var myNumbers = strings.Split(numbers[1], " ")
		var score = 0
		matchingNum = 0
		for i := 0; i < len(winningNumbers); i++ {
			for j := 0; j < len(myNumbers); j++ {
				if winningNumbers[i] == myNumbers[j] {
					if score == 0 {
						score = 1
					} else {
						score = score * 2
					}
					matchingNum++
				}
			}
		}
		for i := 0; i < matchingNum; i++ {
			if currentCard+i < 209 {
				copyList[currentCard+1+i] += 1 * copyList[currentCard]
			}
		}
		totalScore = totalScore + score
	}
	var sum2 = 0
	for i := 0; i < currentCard+1; i++ {
		sum2 += copyList[i]
	}
	return sum2
}
