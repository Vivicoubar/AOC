package main

import (
	"AOC2023/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type State struct {
	pos utils.Pos
	dir int
}

func main() {
	fmt.Println(P1("./AOC2023-17/ex.txt"))
	fmt.Println(P1("./AOC2023-17/input1.txt"))
	fmt.Println(P2("./AOC2023-17/ex.txt"))
	fmt.Println(P2("./AOC2023-17/input1.txt"))
}

func P1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matrix = make([][]int, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		var row = make([]int, 0)
		for _, char := range line {
			var num, _ = strconv.Atoi(string(char))
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	var starts = []State{{pos: utils.Pos{0, 0}, dir: 2}, {pos: utils.Pos{0, 0}, dir: 3}}
	var goal = func(s State) bool {
		return s.pos.X == len(matrix)-1 && s.pos.Y == len(matrix[0])-1
	}
	var neighborsFunc = func(s State) []State {
		return neighboors(matrix, s, 1, 3)
	}
	var costFunc = func(from State, to State) int {
		return cost(matrix, from, to)
	}
	var heuristicFunc = func(s State) int {
		return utils.ManhattanDistance(s.pos, utils.Pos{X: len(matrix) - 1, Y: len(matrix[0]) - 1})
	}
	var _, distance = utils.AstarMultipleStart(starts, goal, neighborsFunc, costFunc, heuristicFunc)
	return distance
}

func P2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matrix = make([][]int, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		var row = make([]int, 0)
		for _, char := range line {
			var num, _ = strconv.Atoi(string(char))
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	var starts = []State{{pos: utils.Pos{0, 0}, dir: 2}, {pos: utils.Pos{0, 0}, dir: 3}}
	var goal = func(s State) bool {
		return s.pos.X == len(matrix)-1 && s.pos.Y == len(matrix[0])-1
	}
	var neighborsFunc = func(s State) []State {
		return neighboors(matrix, s, 4, 10)
	}
	var costFunc = func(from State, to State) int {
		return cost(matrix, from, to)
	}
	var heuristicFunc = func(s State) int {
		return utils.ManhattanDistance(s.pos, utils.Pos{X: len(matrix) - 1, Y: len(matrix[0]) - 1})
	}
	var _, distance = utils.AstarMultipleStart(starts, goal, neighborsFunc, costFunc, heuristicFunc)
	return distance
}

func isValidPos(matrix [][]int, pos utils.Pos) bool {
	return pos.X >= 0 && pos.X < len(matrix) && pos.Y >= 0 && pos.Y < len(matrix[0])
}

var direction = []utils.Pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // left, down, right, up

func neighboors(grid [][]int, s State, mini, maxi int) []State {
	var res []State
	dirs := [2]int{(s.dir + 1) % 4, (s.dir + 3) % 4}
	for _, d := range dirs {
		for i := mini; i <= maxi; i++ {
			pos := utils.Pos{s.pos.X + i*direction[d].X, s.pos.Y + i*direction[d].Y}
			if isValidPos(grid, pos) {
				res = append(res, State{pos, d})
			} else {
				break
			}
		}
	}
	return res
}

func cost(matrix [][]int, from State, to State) int {
	r1, c1 := from.pos.X, from.pos.Y
	r2, c2 := to.pos.X, to.pos.Y
	var res int
	if r1 == r2 {
		if c1 < c2 {
			for i := c1 + 1; i <= c2; i++ {
				res += matrix[i][r1]
			}
		} else if c1 > c2 {
			for i := c1 - 1; i >= c2; i-- {
				res += matrix[i][r1]
			}
		}
		return res
	}
	if c1 == c2 {
		if r1 < r2 {
			for i := r1 + 1; i <= r2; i++ {
				res += matrix[c1][i]
			}
		} else if r1 > r2 {
			for i := r1 - 1; i >= r2; i-- {
				res += matrix[c1][i]
			}
		}
		return res
	}
	return res
}
