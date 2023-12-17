package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ouvrir le fichier en lecture
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstVar = 0
	var secondVar = 0
	var thirdVar = 0
	count := 0
	out := 0
	for scanner.Scan() {
		line := scanner.Text()
		curNum, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Erreur lors de la conversion en entier :", err)
			return // Passer à la ligne suivante en cas d'erreur
		}
		//Init var
		if count < 3 {

			if count == 0 {
				firstVar = curNum
				count += 1
			} else if count == 1 {
				secondVar = curNum
				count += 1
			} else if count == 2 {
				thirdVar = curNum
				count += 1
			}
			continue
		}
		//We compare the third var (which was set the soonest) one and the current line and update the soonest to the current
		curCount := count % 3
		if curCount == 0 {
			if firstVar < curNum {
				out += 1
			}
			firstVar = curNum
		} else if curCount == 1 {
			if secondVar < curNum {
				out += 1
			}
			secondVar = curNum
		} else if curCount == 2 {
			if thirdVar < curNum {
				out += 1
			}
			thirdVar = curNum
		}
		count++
	}
	fmt.Println(out)
}
