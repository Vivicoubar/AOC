package main

import (
	"AOC2023/utils"
	"bufio"
	"fmt"
	"math"
	"os"
)

type Pos struct {
	row, col int
}

type Slope struct {
	char       string
	drow, dcol int
}

type State struct {
	row, col, distance int
}

func main() {
	var input = "AOC2023-23/input1.txt"
	var ex = "AOC2023-23/ex.txt"
	fmt.Println("Partie 1 ex:", P1(ex))
	fmt.Println("Partie 1:", P1(input))
	fmt.Println("Partie 2 ex:", P2(ex))
	fmt.Println("Partie 2:", P2(input))
}

func P1(input string) int {
	var corners, startPos, maxR, maxC = getSummitAndCorners(input, true)
	var p1 = 0
	var wentThatWay = make([][]bool, 0)
	for r := 0; r < maxR; r++ {
		var row = make([]bool, maxC)
		for c := 0; c < maxC; c++ {
			row[c] = false
		}
		wentThatWay = append(wentThatWay, row)
	}
	var throughGraph func(summit Pos, distance int)
	throughGraph = func(summit Pos, distance int) {
		if !wentThatWay[summit.row][summit.col] {
			wentThatWay[summit.row][summit.col] = true
			if summit.row == maxR-1 {
				p1 = int(math.Max(float64(p1), float64(distance)))
			}
			for _, state := range corners[summit] {
				throughGraph(Pos{state.row, state.col}, distance+state.distance)
			}
			wentThatWay[summit.row][summit.col] = false
		}
	}
	throughGraph(startPos, 0)
	return p1
}

func P2(input string) int {
	var corners, startPos, maxR, maxC = getSummitAndCorners(input, false)
	var p2 = 0
	var wentThatWay = make([][]bool, 0)
	for r := 0; r < maxR; r++ {
		var row = make([]bool, maxC)
		for c := 0; c < maxC; c++ {
			row[c] = false
		}
		wentThatWay = append(wentThatWay, row)
	}
	var throughGraph func(summit Pos, distance int)
	throughGraph = func(summit Pos, distance int) {
		if !wentThatWay[summit.row][summit.col] {
			wentThatWay[summit.row][summit.col] = true
			if summit.row == maxR-1 {
				p2 = int(math.Max(float64(p2), float64(distance)))
			}
			for _, state := range corners[summit] {
				throughGraph(Pos{state.row, state.col}, distance+state.distance)
			}
			wentThatWay[summit.row][summit.col] = false
		}
	}
	throughGraph(startPos, 0)
	return p2
}

func getGridFromFile(input string) [][]string {
	//Read file
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			os.Exit(1)
		}
	}(file)
	var scanner = bufio.NewScanner(file)
	var grid = make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var row = make([]string, 0)
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}
	return grid
}

func getSummits(grid [][]string, maxR, maxC int) (utils.Set[Pos], Pos) {
	var slopes = []Slope{{"^", -1, 0}, {"v", 1, 0}, {">", 0, 1}, {"<", 0, -1}}
	//We will create a graph with every intersection, because we don't care about the direction, only the distance
	var summits = utils.NewSet[Pos]()
	for r := 0; r < maxR; r++ {
		for c := 0; c < maxC; c++ {
			var neighbour = 0
			for _, slope := range slopes {
				if r+slope.drow >= 0 && r+slope.drow < maxR && c+slope.dcol >= 0 && c+slope.dcol < maxC {
					if grid[r+slope.drow][c+slope.dcol] != "#" {
						neighbour += 1
					}
				}
			}
			if neighbour > 2 && (grid[r][c] == "." || grid[r][c] == "^" || grid[r][c] == "v" || grid[r][c] == "<" || grid[r][c] == ">") {
				summits.Add(Pos{r, c})
			}
		}
	}
	var startPos = Pos{0, 0}
	//We need to add the start and end points in our graph
	for col := 0; col < maxC; col++ {
		if grid[0][col] == "." {
			summits.Add(Pos{0, col})
			startPos = Pos{0, col}
		}
		if grid[maxR-1][col] == "." {
			summits.Add(Pos{maxR - 1, col})
		}
	}
	return summits, startPos
}

func getCorners(summits utils.Set[Pos], grid [][]string, maxR, maxC int, slopes []Slope, isPart1 bool) map[Pos][]State {
	var corners = make(map[Pos][]State)
	for summit := range summits {
		var newSummits = make([]State, 0)
		//We do a BFS to find the distance between every intersection (here, BFS is fine because we don't have any choice of path, so we'll always take the longest (the only one)
		var stack = make([]State, 0)
		var visited = utils.NewSet[Pos]()
		stack = append(stack, State{summit.row, summit.col, 0})
		for len(stack) > 0 {
			var current = stack[0]
			stack = stack[1:]
			if !visited.Contains(Pos{current.row, current.col}) {
				visited.Add(Pos{current.row, current.col})
				if summits.Contains(Pos{current.row, current.col}) && (current.row != summit.row || current.col != summit.col) {
					newSummits = append(newSummits, current)
					continue
				}
				for _, slope := range slopes {
					if current.row+slope.drow >= 0 && current.row+slope.drow < maxR && current.col+slope.dcol >= 0 && current.col+slope.dcol < maxC {
						if grid[current.row+slope.drow][current.col+slope.dcol] != "#" {
							if isPart1 {
								if (grid[current.row][current.col] == "v" || grid[current.row][current.col] == "^" || grid[current.row][current.col] == "<" || grid[current.row][current.col] == ">") && (grid[current.row][current.col] != slope.char) {
									continue
								}
							}
							stack = append(stack, State{current.row + slope.drow, current.col + slope.dcol, current.distance + 1})
						}
					}
				}
			}
		}
		corners[summit] = newSummits
	}
	return corners
}

func getSummitAndCorners(input string, isPart1 bool) (map[Pos][]State, Pos, int, int) {
	var grid = getGridFromFile(input)
	var maxR = len(grid)
	var maxC = len(grid[0])
	var summits, startPos = getSummits(grid, maxR, maxC)
	var corners = getCorners(summits, grid, maxR, maxC, []Slope{{"^", -1, 0}, {"v", 1, 0}, {">", 0, 1}, {"<", 0, -1}}, isPart1)
	return corners, startPos, maxR, maxC
}
