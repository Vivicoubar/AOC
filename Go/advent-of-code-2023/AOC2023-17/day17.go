package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Pos struct {
	row int
	col int
	dir int // 0: up, 1: right, 2: down, 3: left
}

func main() {
	fmt.Println(P1("./AOC2023-17/ex.txt"))
	fmt.Println(P1("./AOC2023-17/input1.txt"))
	fmt.Println(P2("./AOC2023-17/ex.txt"))
	fmt.Println(P2("./AOC2023-17/input1.txt"))
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
	var matrix = make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var row = make([]int, 0)
		for _, char := range line {
			var num, _ = strconv.Atoi(string(char))
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	//We use Djikstra's algorithm to find the shortest path
	var maxR = len(matrix)
	var maxC = len(matrix[0])
	var dist = make(map[Pos]int)
	//Initialisation
	var queue = make([]Pos, 0)
	//We add every starting point to the queue without its distance, except if we come back
	for i := 0; i < 4; i++ {
		var pos = Pos{0, 0, i}
		dist[pos] = 0
		queue = append(queue, pos)
	}
	//Djikstra's algorithm
	for len(queue) > 0 {
		//fmt.Println(len(queue)) << UNCOMMENT TO SEE IF ITS STUCK IN A INFINITE LOOP
		// Pop the element at index pos, which is the one with the smallest distance
		var pos = findMinPosId(queue, dist)
		var current = queue[pos]
		queue = append(queue[:pos], queue[pos+1:]...)
		var row = current.row
		var col = current.col
		var dir = current.dir
		var rowDir = 0
		var colDir = 0
		//Get the direction
		if dir == 0 {
			rowDir = -1
		} else if dir == 1 {
			colDir = 1
		} else if dir == 2 {
			rowDir = 1
		} else if dir == 3 {
			colDir = -1
		}
		//Handle the possibility to go in straight line up to 3 times
		for i := 1; i <= 3; i++ {
			var newRow = row + rowDir*i
			var newCol = col + colDir*i
			if newRow >= 0 && newRow < maxR && newCol >= 0 && newCol < maxC {
				var newDist = dist[current]
				for j := 1; j < i+1; j++ {
					newDist += matrix[row+rowDir*j][col+colDir*j]
				}
				//Rotate the direction
				// Rotate the direction
				newDir := (dir - 1) % 4 // Ensure newDir is in the range [0, 3]
				newPos := Pos{newRow, newCol, newDir}
				distInMap, ok := dist[newPos]
				if !ok {
					distInMap = 10000000000000000
				}
				if newDist < distInMap {
					dist[newPos] = newDist
					queue = addIfNotPresent(queue, newPos)
				}

				newDir = (dir + 1) % 4 // Ensure newDir is in the range [0, 3]
				newPos = Pos{newRow, newCol, newDir}
				distInMap, ok = dist[newPos]
				if !ok {
					distInMap = 100000000000000000
				}
				if newDist < distInMap {
					dist[newPos] = newDist
					queue = addIfNotPresent(queue, newPos)
				}
			}
		}
	}
	//Get the shortest path
	var minDist, ok = dist[Pos{maxR - 1, maxC - 1, 0}]
	if !ok {
		minDist = 1000000000
	}
	for i := 1; i < 4; i++ {
		var pos = Pos{maxR - 1, maxC - 1, i}
		var distInMap, forok = dist[pos]
		if forok && distInMap < minDist {
			minDist = distInMap
		}
	}
	return minDist

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
	var matrix = make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var row = make([]int, 0)
		for _, char := range line {
			var num, _ = strconv.Atoi(string(char))
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	var maxR = len(matrix)
	var maxC = len(matrix[0])
	var dist = make(map[Pos]int)
	//Initialisation
	var queue = make([]Pos, 0)
	//We add every starting point to the queue without its distance, except if we come back
	for i := 0; i < 4; i++ {
		var pos = Pos{0, 0, i}
		dist[pos] = 0
		queue = append(queue, pos)
	}
	//Djikstra's algorithm
	for len(queue) > 0 {
		//fmt.Println(len(queue)) << UNCOMMENT TO SEE IF ITS STUCK IN A INFINITE LOOP
		// Pop the element at index pos, which is the one with the smallest distance
		var pos = findMinPosId(queue, dist)
		var current = queue[pos]
		queue = append(queue[:pos], queue[pos+1:]...)
		var row = current.row
		var col = current.col
		var dir = current.dir
		var rowDir = 0
		var colDir = 0
		//Get the direction
		if dir == 0 {
			rowDir = -1
		} else if dir == 1 {
			colDir = 1
		} else if dir == 2 {
			rowDir = 1
		} else if dir == 3 {
			colDir = -1
		}
		//Handle the possibility to go in straight line from 4 to 10 times = we explore only the summit within 4 to 10 steps
		for i := 4; i <= 10; i++ {
			var newRow = row + rowDir*i
			var newCol = col + colDir*i
			if newRow >= 0 && newRow < maxR && newCol >= 0 && newCol < maxC {
				var newDist = dist[current]
				for j := 1; j < i+1; j++ {
					newDist += matrix[row+rowDir*j][col+colDir*j]
				}
				//Rotate the direction
				// Rotate the direction
				newDir := (dir - 1) % 4 // Ensure newDir is in the range [0, 3]
				newPos := Pos{newRow, newCol, newDir}
				distInMap, ok := dist[newPos]
				if !ok {
					distInMap = 10000000000000000
				}
				if newDist < distInMap {
					dist[newPos] = newDist
					queue = addIfNotPresent(queue, newPos)
				}

				newDir = (dir + 1) % 4 // Ensure newDir is in the range [0, 3]
				newPos = Pos{newRow, newCol, newDir}
				distInMap, ok = dist[newPos]
				if !ok {
					distInMap = 100000000000000000
				}
				if newDist < distInMap {
					dist[newPos] = newDist
					queue = addIfNotPresent(queue, newPos)
				}
			}
		}
	}
	//Get the shortest path
	var minDist, ok = dist[Pos{maxR - 1, maxC - 1, 0}]
	if !ok {
		minDist = 1000000000
	}
	for i := 1; i < 4; i++ {
		var pos = Pos{maxR - 1, maxC - 1, i}
		var distInMap, forok = dist[pos]
		if forok && distInMap < minDist {
			minDist = distInMap
		}
	}
	return minDist
}

func findMinPosId(queue []Pos, dist map[Pos]int) int {
	var minPos = 0
	var minDist = dist[queue[0]]
	for i, pos := range queue {
		if dist[pos] < minDist {
			minPos = i
			minDist = dist[pos]
		}
	}
	return minPos
}

func addIfNotPresent(queue []Pos, pos Pos) []Pos {
	for _, p := range queue {
		if p == pos {
			return queue
		}
	}
	return append(queue, pos)
}
