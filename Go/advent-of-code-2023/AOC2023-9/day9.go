package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(P1("ex.txt"))
	fmt.Println(P1("input1.txt"))
	fmt.Println(P2("ex.txt"))
	fmt.Println(P2("input1.txt"))
}

func P1(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	var toAddLastValues []int = make([]int, 0)
	var curDiff = make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		curDiff = make([][]int, 0)
		var elements = strings.Split(line, " ")
		var numElements = make([]int, 0)
		for _, element := range elements {
			var num, _ = strconv.Atoi(element)
			numElements = append(numElements, num)
		}
		//We begin to process the line
		curDiff = append(curDiff, numElements)
		count := 0
		isGood := false
		for !isGood {
			var newDif = make([]int, 0)
			for i := 0; i < len(curDiff[count]); i++ {
				if i > 0 {
					newDif = append(newDif, curDiff[count][i]-curDiff[count][i-1])
				}
			}
			isGood = true
			{
				for i := 0; i < len(newDif); i++ {
					if newDif[i] != 0 {
						isGood = false
						break
					}
				}
			}
			curDiff = append(curDiff, newDif)
			count++
		}
		var lastLine = curDiff[count]
		lastLine = append(lastLine, 0)
		curDiff[count] = lastLine
		var prevElement = 0
		for i := 1; i <= count; i++ {
			var curLine = curDiff[count-i]
			var newVal = prevElement + curLine[len(curLine)-1]
			curLine = append(curLine, newVal)
			curDiff[count-i] = curLine
			prevElement = newVal
		}
		var lastValue = curDiff[0][len(curDiff[0])-1]
		toAddLastValues = append(toAddLastValues, lastValue)
	}
	var p1 = 0
	for _, value := range toAddLastValues {
		p1 += value
	}
	return p1
}

func P2(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	var toAddFirstValues []int = make([]int, 0)
	var curDiff = make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		curDiff = make([][]int, 0)
		var elements = strings.Split(line, " ")
		var numElements = make([]int, 0)
		for _, element := range elements {
			var num, _ = strconv.Atoi(element)
			numElements = append(numElements, num)
		}
		//We begin to process the line
		curDiff = append(curDiff, numElements)
		count := 0
		isGood := false
		for !isGood {
			var newDif = make([]int, 0)
			for i := 0; i < len(curDiff[count]); i++ {
				if i > 0 {
					newDif = append(newDif, curDiff[count][i]-curDiff[count][i-1])
				}
			}
			isGood = true
			{
				for i := 0; i < len(newDif); i++ {
					if newDif[i] != 0 {
						isGood = false
						break
					}
				}
			}
			curDiff = append(curDiff, newDif)
			count++
		}
		var lastLine = curDiff[count]
		lastLine = append(lastLine[0:1], lastLine[0:]...)
		lastLine[0] = 0
		curDiff[count] = lastLine
		var prevElement = 0
		for i := 1; i <= count; i++ {
			var curLine = curDiff[count-i]
			var newVal = curLine[0] - prevElement
			curLine = append(curLine[:1], curLine[0:]...)
			curLine[0] = newVal
			curDiff[count-i] = curLine
			prevElement = newVal
		}
		toAddFirstValues = append(toAddFirstValues, curDiff[0][0])
	}
	var p2 = 0
	for _, value := range toAddFirstValues {
		p2 += value
	}
	return p2
}
