package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	var curNumber = 0
	var sum = 0
	var decodedValues []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var content = strings.Split(line, " | ")
		curNumber += countNumbers(content[1])
		var entryAndOutput = strings.Split(content[1]+" "+content[0], " ")
		decodedValues = findDecodedValues(entryAndOutput)
		sum += decode(content[1], decodedValues)
	}
	fmt.Println(curNumber)
	fmt.Println(sum)
}

func countDif(element string, origin string) int {
	var count = 0
	for i := 0; i < len(origin); i++ {
		var inString = false
		for j := 0; j < len(element); j++ {
			if origin[i] == element[j] {
				inString = true
				break
			}
		}
		if !inString {
			count++
		}
	}
	return count
}

func findDecodedValues(line []string) []string {
	var decodedValues []string = make([]string, 10)
	for _, element := range line {
		s := strings.Split(element, "")
		sort.Strings(s)
		element = strings.Join(s, "")
		if len(element) == 4 {
			decodedValues[4] = element
		} else if len(element) == 7 {
			decodedValues[8] = element
		} else if len(element) == 2 {
			decodedValues[1] = element
		} else if len(element) == 3 {
			decodedValues[7] = element
		}
	}
	for _, element := range line {
		s := strings.Split(element, "")
		sort.Strings(s)
		element = strings.Join(s, "")
		if len(element) == 6 {
			if countDif(element, decodedValues[4]) == 0 {
				decodedValues[9] = element
			} else if countDif(element, decodedValues[7]) == 0 {
				decodedValues[0] = element
			} else {
				decodedValues[6] = element
			}
		} else if len(element) == 5 {
			if countDif(element, decodedValues[1]) == 0 {
				decodedValues[3] = element
			} else if countDif(element, decodedValues[4]) == 2 {
				decodedValues[2] = element
			} else {
				decodedValues[5] = element
			}
		}
	}
	return decodedValues
}

func decode(line string, decodedValues []string) int {
	var count = make([]int, 4)

	var curLine = strings.Split(line, " ")
	//create an array of arrays of strings, with 1 in each substring
	for i, element := range curLine {
		//sort the element in the array
		s := strings.Split(element, "")
		sort.Strings(s)
		element = strings.Join(s, "")
		if element == decodedValues[0] {
			count[i] = 0
		} else if element == decodedValues[1] {
			count[i] = 1
		} else if element == decodedValues[2] {
			count[i] = 2
		} else if element == decodedValues[3] {
			count[i] = 3
		} else if element == decodedValues[4] {
			count[i] = 4
		} else if element == decodedValues[5] {
			count[i] = 5
		} else if element == decodedValues[6] {
			count[i] = 6
		} else if element == decodedValues[7] {
			count[i] = 7
		} else if element == decodedValues[8] {
			count[i] = 8
		} else if element == decodedValues[9] {
			count[i] = 9
		}
	}
	var result float64 = 0
	for i, element := range count {
		result += float64(element) * math.Pow(10, float64(3-i))
	}
	return int(result)
}

func countNumbers(line string) int {
	var count = 0
	var curLine = strings.Split(line, " ")
	for _, element := range curLine {
		if len(element) == 4 || len(element) == 2 || len(element) == 3 || len(element) == 7 {
			count++
		}
	}
	return count
}
