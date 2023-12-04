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
	var day0, day1, day2, day3, day4, day5, day6, day7, day8 = 0, 0, 0, 0, 0, 0, 0, 0, 0
	var numDays = 256
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var inititalState = strings.Split(line, ",")
		for _, element := range inititalState {
			var numElement, _ = strconv.Atoi(element)
			if numElement == 1 {
				day1++
			} else if numElement == 2 {
				day2++
			} else if numElement == 3 {
				day3++
			} else if numElement == 4 {
				day4++
			} else if numElement == 5 {
				day5++
			} else if numElement == 6 {
				day6++
			} else if numElement == 7 {
				day7++
			} else if numElement == 8 {
				day8++
			}
		}
	}
	for i := 0; i < numDays; i++ {
		var newDay0, newDay1, newDay2, newDay3, newDay4, newDay5, newDay6, newDay7, newDay8 = day1,
			day2,
			day3,
			day4,
			day5,
			day6,
			day7 + day0,
			day8,
			day0
		day0 = newDay0
		day1 = newDay1
		day2 = newDay2
		day3 = newDay3
		day4 = newDay4
		day5 = newDay5
		day6 = newDay6
		day7 = newDay7
		day8 = newDay8
	}
	var sum = day0 + day1 + day2 + day3 + day4 + day5 + day6 + day7 + day8
	fmt.Println(sum)
}
