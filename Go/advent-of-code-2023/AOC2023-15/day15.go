package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Box struct {
	lenses []Lens
}
type Lens struct {
	label string
	focal int
}

func calcHash(input string) int {
	var cur = 0
	var chars = input
	for _, char := range chars {
		cur += int(char)
		cur *= 17
		cur %= 256
	}
	return cur
}

func main() {
	fmt.Println("Partie 1 Ex: ", P1("AOC2023-15/ex.txt"))
	fmt.Println("Partie 1 : ", P1("AOC2023-15/input1.txt"))
	fmt.Println("Partie 2 Ex: ", P2("AOC2023-15/ex.txt"))
	fmt.Println("Partie 2 : ", P2("AOC2023-15/input1.txt"))
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
	var hashPieces []string
	for scanner.Scan() {
		hashPieces = strings.Split(strings.TrimSpace(scanner.Text()), ",")
	}
	var p1 = 0
	for _, piece := range hashPieces {
		p1 += calcHash(piece)
	}
	return p1
}

func P2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hashPieces []string
	for scanner.Scan() {
		hashPieces = strings.Split(strings.TrimSpace(scanner.Text()), ",")
	}
	var boxes = make([]Box, 256)
	for _, piece := range hashPieces {
		var operator = ""
		var label = ""
		if '0' <= piece[len(piece)-1] && piece[len(piece)-1] <= '9' {
			operator = string(piece[len(piece)-2])
			label = piece[:len(piece)-2]
		} else {
			operator = string(piece[len(piece)-1])
			label = piece[:len(piece)-1]
		}
		var hash = calcHash(label)
		var releventBox = boxes[hash]
		if operator == "=" {
			var found = false
			for j, lens := range releventBox.lenses {
				if lens.label == label {
					lens.focal, _ = strconv.Atoi(string(piece[len(piece)-1]))
					releventBox.lenses[j] = lens
					boxes[hash] = releventBox
					found = true
					break
				}
			}
			if !found {
				var focal, _ = strconv.Atoi(string(piece[len(piece)-1]))
				releventBox.lenses = append(releventBox.lenses, Lens{label, focal})
				boxes[calcHash(label)] = releventBox
			}
		} else if operator == "-" {
			var newLenses = make([]Lens, 0)
			for _, lens := range releventBox.lenses {
				if lens.label == label {
					continue
				} else {
					newLenses = append(newLenses, lens)
				}
			}
			releventBox.lenses = newLenses
			boxes[calcHash(label)] = releventBox
		}
	}
	var p2 = 0
	for i, box := range boxes {
		for j, lens := range box.lenses {
			var toAdd = (j + 1) * (i + 1) * lens.focal
			p2 += toAdd
		}
	}
	return p2
}
