package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Brick struct {
	X1, Y1, Z1, X2, Y2, Z2 int
}

type ByZ []Brick

func main() {
	var input = "AOC2023-22/input1.txt"
	var ex = "AOC2023-22/ex.txt"
	fmt.Println("Partie 1 ex:", P1(ex))
	fmt.Println("Partie 1:", P1(input))
	fmt.Println("Partie 2 ex:", P2(ex))
	fmt.Println("Partie 2:", P2(input))
}

func P1(input string) int {
	//read file
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	var data []Brick
	for scanner.Scan() {
		line := scanner.Text()
		var elem = strings.Split(line, "~")
		var x1, y1, h1, x2, y2, h2 int
		var beginCoords = strings.Split(elem[0], ",")
		var endCoords = strings.Split(elem[1], ",")
		x1, _ = strconv.Atoi(beginCoords[0])
		y1, _ = strconv.Atoi(beginCoords[1])
		h1, _ = strconv.Atoi(beginCoords[2])
		x2, _ = strconv.Atoi(endCoords[0])
		y2, _ = strconv.Atoi(endCoords[1])
		h2, _ = strconv.Atoi(endCoords[2])
		var brick = Brick{x1, y1, h1, x2, y2, h2}
		data = append(data, brick)
	}
	sort.Sort(ByZ(data))
	var _, fallen = fall(data)
	var p1 = 0
	for i, _ := range fallen {
		var withDeletedBrick = make([]Brick, 0)
		withDeletedBrick = append(withDeletedBrick, fallen[:i]...)
		withDeletedBrick = append(withDeletedBrick, fallen[i+1:]...)
		var count, _ = fall(withDeletedBrick)
		if count == 0 {
			p1 += 1
		}
	}
	return p1
}

func P2(input string) int {
	//read file
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	var scanner = bufio.NewScanner(file)
	var data []Brick
	for scanner.Scan() {
		line := scanner.Text()
		var elem = strings.Split(line, "~")
		var x1, y1, h1, x2, y2, h2 int
		var beginCoords = strings.Split(elem[0], ",")
		var endCoords = strings.Split(elem[1], ",")
		x1, _ = strconv.Atoi(beginCoords[0])
		y1, _ = strconv.Atoi(beginCoords[1])
		h1, _ = strconv.Atoi(beginCoords[2])
		x2, _ = strconv.Atoi(endCoords[0])
		y2, _ = strconv.Atoi(endCoords[1])
		h2, _ = strconv.Atoi(endCoords[2])
		var brick = Brick{x1, y1, h1, x2, y2, h2}
		data = append(data, brick)
	}
	sort.Sort(ByZ(data))
	var _, fallen = fall(data)
	var p2 = 0
	for i, _ := range fallen {
		var withDeletedBrick = make([]Brick, 0)
		withDeletedBrick = append(withDeletedBrick, fallen[:i]...)
		withDeletedBrick = append(withDeletedBrick, fallen[i+1:]...)
		var count, _ = fall(withDeletedBrick)
		p2 += count
	}
	return p2
}

func fallenBrick(higher map[string]int, brick Brick) Brick {
	var top = 0
	for x := brick.X1; x <= brick.X2; x++ {
		for y := brick.Y1; y <= brick.Y2; y++ {
			var curTop = higher[strconv.Itoa(x)+","+strconv.Itoa(y)]
			if curTop > top {
				top = curTop
			}
		}
	}
	var height = brick.Z1 - top - 1
	if height < 0 {
		height = 0
	}
	return Brick{brick.X1, brick.Y1, top + 1, brick.X2, brick.Y2, brick.Z2 - height}
}

func fall(heap []Brick) (int, []Brick) {
	var higher = make(map[string]int, 0)
	var stack = make([]Brick, 0)
	var fallBrick = 0
	for _, brick := range heap {
		var curBrick Brick = fallenBrick(higher, brick)
		if curBrick.Z2 != brick.Z2 {
			fallBrick += 1
		}
		stack = append(stack, curBrick)
		for x := brick.X1; x <= brick.X2; x++ {
			for y := brick.Y1; y <= brick.Y2; y++ {
				higher[strconv.Itoa(x)+","+strconv.Itoa(y)] = curBrick.Z2
			}
		}
	}
	return fallBrick, stack
}

func (a ByZ) Len() int      { return len(a) }
func (a ByZ) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByZ) Less(i, j int) bool {
	return a[i].Z1 < a[j].Z1
}
