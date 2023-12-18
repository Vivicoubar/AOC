package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	direction string
	value     int
}

type Pos struct {
	row int
	col int
}

func main() {
	fmt.Println("Partie 1 Ex:", P1("AOC2023-18/ex.txt"))
	fmt.Println("Partie 1:", P1("AOC2023-18/input1.txt"))
	fmt.Println("Partie 2 Ex:", P2("AOC2023-18/ex.txt"))
	fmt.Println("Partie 2:", P2("AOC2023-18/input1.txt"))
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
	var moves = make([]Move, 0)
	var directionVector = make(map[string]Pos)
	directionVector["D"], directionVector["R"], directionVector["U"], directionVector["L"] = Pos{1, 0}, Pos{0, 1}, Pos{-1, 0}, Pos{0, -1}
	for scanner.Scan() {
		line := scanner.Text()
		var elements = strings.Split(line, " ")
		var value, _ = strconv.Atoi(elements[1])
		move := Move{elements[0], value}
		moves = append(moves, move)
	}
	var points = make([]Pos, 0)
	var newPoint = Pos{0, 0}
	for _, move := range moves {
		var row, col = newPoint.row, newPoint.col
		var directionVec = directionVector[move.direction]
		newPoint = (Pos{row + directionVec.row*move.value, col + directionVec.col*move.value})
		points = append(points, newPoint)
	}
	var interior = calcPolyVolume(points)
	var border = 0
	for _, move := range moves {
		border += move.value
	}
	//Half of the border is counted in the interior
	var p1 = interior + border/2 + 1
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
	scanner := bufio.NewScanner(file)
	var moves = make([]Move, 0)
	var directionVector = make(map[string]Pos)
	directionVector["D"], directionVector["R"], directionVector["U"], directionVector["L"] = Pos{1, 0}, Pos{0, 1}, Pos{-1, 0}, Pos{0, -1}
	for scanner.Scan() {
		line := scanner.Text()
		var elements = strings.Split(line, " ")
		var numValue = 0
		var hexaNumber = elements[2][2:7]
		var pow = 0
		//Convert hexadecimal to decimal
		for i := len(hexaNumber) - 1; i >= 0; i-- {
			var digit = hexaNumber[i]
			if digit >= '0' && digit <= '9' {
				numValue = numValue + int(math.Pow(16, float64(pow)))*int(digit-'0')
			} else {
				numValue = numValue + int(math.Pow(16, float64(pow)))*int(digit-'a'+10)
			}
			pow++
		}
		var dir = ""
		if elements[2][7] == '0' {
			dir = "R"
		} else if elements[2][7] == '1' {
			dir = "D"
		} else if elements[2][7] == '2' {
			dir = "L"
		} else if elements[2][7] == '3' {
			dir = "U"
		}
		move := Move{dir, numValue}
		moves = append(moves, move)
	}
	var points = make([]Pos, 0)
	var newPoint = Pos{0, 0}
	for _, move := range moves {
		var row, col = newPoint.row, newPoint.col
		var directionVec = directionVector[move.direction]
		newPoint = Pos{row + directionVec.row*move.value, col + directionVec.col*move.value}
		points = append(points, newPoint)
	}
	var interior = calcPolyVolume(points)
	var border = 0
	for _, move := range moves {
		border += move.value
	}
	//Half of the border is counted in the interior
	var p2 = interior + border/2 + 1
	return p2
}

func calcPolyVolume(points []Pos) int {
	//Pick formula
	var volume = 0
	for i := 0; i < len(points)-1; i++ {
		var point = points[i]
		var newPoint = points[i+1]
		volume += (point.row + newPoint.row) * (point.col - newPoint.col)
	}
	return volume / 2
}
