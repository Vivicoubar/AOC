package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	row int
	col int
}

type State struct {
	reached bool
	isStart bool
}

type QueueElement struct {
	pos  Pos
	step int
}

func main() {
	fmt.Println(P1("AOC2023-21/ex.txt", 6))
	fmt.Println(P1("AOC2023-21/input1.txt", 64))
	fmt.Println(P2("AOC2023-21/input1.txt"))
}

func P1(input string, num int) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var startPos Pos
	var garden = make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var row = make([]string, 0)
		for _, c := range line {
			if c == 'S' {
				startPos = Pos{len(garden), len(row)}
			}
			row = append(row, string(c))
		}
		garden = append(garden, row)
	}
	//BFS
	var frontier = make([]QueueElement, 0)
	frontier = append(frontier, QueueElement{startPos, 0})
	var reached = make(map[Pos]State)
	var previousStep = 0
	reached[startPos] = State{true, true}
	var maxSteps = num
	var i = 0
	for len(frontier) > 0 && i < maxSteps {
		var current = frontier[0]
		i = current.step
		frontier = frontier[1:]
		for _, next := range getNeighbors(garden, current.pos) {
			if !reached[next].reached {
				frontier = append(frontier, QueueElement{next, i + 1})
				reached[next] = State{true, false}
			}
		}

		i = frontier[0].step
		if i > previousStep {
			previousStep = i
			//seeMap(garden, reached)
			reached = make(map[Pos]State)
		}
	}
	return len(frontier)
}

func getNeighbors(garden [][]string, pos Pos) []Pos {
	var neighbors = make([]Pos, 0)
	if pos.row > 0 && garden[pos.row-1][pos.col] != "#" {
		neighbors = append(neighbors, Pos{pos.row - 1, pos.col})
	}
	if pos.row < len(garden)-1 && garden[pos.row+1][pos.col] != "#" {
		neighbors = append(neighbors, Pos{pos.row + 1, pos.col})
	}
	if pos.col > 0 && garden[pos.row][pos.col-1] != "#" {
		neighbors = append(neighbors, Pos{pos.row, pos.col - 1})
	}
	if pos.col < len(garden[0])-1 && garden[pos.row][pos.col+1] != "#" {
		neighbors = append(neighbors, Pos{pos.row, pos.col + 1})
	}
	return neighbors
}

func seeMap(garden [][]string, reached map[Pos]State) {
	for xRow, row := range garden {
		for xCol, c := range row {
			if reached[Pos{xRow, xCol}].reached && reached[Pos{xRow, xCol}].isStart {
				fmt.Print("S")
			} else {
				if reached[Pos{xRow, xCol}].reached {
					fmt.Print("0")
				} else {
					fmt.Print(c)
				}
			}
		}
		fmt.Println()
	}
}
