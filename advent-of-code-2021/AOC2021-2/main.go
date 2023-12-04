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

	scanner := bufio.NewScanner(file)
	var length = 0
	var depth = 0
	var aim = 0
	for scanner.Scan() {
		line := scanner.Text()
		//Split each line into a slice of strings with " " as separator
		split := strings.Split(line, " ")
		//Cast the second element into an int
		val, _ := strconv.Atoi(split[1])
		if split[0] == "forward" {
			length += val
			depth += aim * val
		}
		if split[0] == "down" {
			aim += val
		}
		if split[0] == "up" {
			aim -= val
		}
	}
	fmt.Println(length, depth, length*depth)

}
