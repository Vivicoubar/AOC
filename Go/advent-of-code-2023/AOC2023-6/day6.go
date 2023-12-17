package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 Ex: ", P1("AOC2023-6/ex.txt"))
	fmt.Println("Part 1 : ", P1("AOC2023-6/input.txt"))
	fmt.Println("Part 2 Ex: ", P2("AOC2023-6/ex.txt"))
	fmt.Println("Part 2 : ", P2("AOC2023-6/input.txt"))
}

func P2(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var times = 0
	var record = 0
	var numBeat = 0
	// Read the first line
	if scanner.Scan() {
		line := scanner.Text()
		var timesRead = strings.TrimSpace(line[5:])
		timesRead = strings.ReplaceAll(timesRead, " ", "")
		times, _ = strconv.Atoi(timesRead)
	}
	// Read the second line
	if scanner.Scan() {
		line := scanner.Text()
		recordRead := strings.TrimSpace(line[9:])
		recordRead = strings.ReplaceAll(recordRead, " ", "")
		record, _ = strconv.Atoi(recordRead)
	}
	for i := 0; i < times; i++ {
		timeLeft := times - i
		if timeLeft*i > record {
			numBeat += 1
		}
	}
	return numBeat
}

func P1(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var times []int
	var record []int
	var numBeat []int
	// Read the first line
	if scanner.Scan() {
		line := scanner.Text()
		timesRead := strings.Split(strings.TrimSpace(line[5:]), " ")

		for _, time := range timesRead {
			if time == "" {
				continue
			}
			timeVal, _ := strconv.Atoi(strings.TrimSpace(time))
			times = append(times, timeVal)
		}
	}

	// Read the second line
	if scanner.Scan() {
		line := scanner.Text()
		recordRead := strings.Split(strings.TrimSpace(line[9:]), " ")
		for _, rec := range recordRead {
			if rec == "" {
				continue
			}
			recVal, _ := strconv.Atoi(strings.TrimSpace(rec))
			record = append(record, recVal)
		}
	}
	numBeat = make([]int, len(record))
	for i := 0; i < len(times); i++ {
		for j := 0; j < times[i]; j++ {
			timeLeft := times[i] - j
			if timeLeft*j > record[i] {
				numBeat[i]++
			}
		}
	}
	var p1 = 1
	for i := 0; i < len(numBeat); i++ {
		p1 *= numBeat[i]
	}
	return p1
}
