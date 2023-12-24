package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HailStone struct {
	px, py, pz int
	vx, vy, vz int
}

func main() {
	var input string = "AOC2023-24/input1.txt"
	fmt.Println(P1(input))
	fmt.Println(P2(input))
}

func getHailStone(line string) []HailStone {
	file, err := os.Open(line)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()
	var stones []HailStone
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = scanner.Text()
		var elem = strings.Split(line, " @ ")
		var posElem = strings.Split(elem[0], ",")
		for i := 0; i < len(posElem); i++ {
			posElem[i] = strings.TrimSpace(posElem[i])
		}
		var velElem = strings.Split(elem[1], ", ")
		for i := 0; i < len(velElem); i++ {
			velElem[i] = strings.TrimSpace(velElem[i])
		}
		var stone HailStone
		stone.px, _ = strconv.Atoi(posElem[0])
		stone.py, _ = strconv.Atoi(posElem[1])
		stone.pz, _ = strconv.Atoi(posElem[2])
		stone.vx, _ = strconv.Atoi(velElem[0])
		stone.vy, _ = strconv.Atoi(velElem[1])
		stone.vz, _ = strconv.Atoi(velElem[2])
		stones = append(stones, stone)
	}
	return stones
}

func P1(input string) int {

	var stones []HailStone = getHailStone(input)
	var marginBottom float64 = 2e14
	var marginTop float64 = 4e14
	var p1 = 0

	//Linear system is:
	// xa(t) = apx + avx*t
	// ya(t) = apy + avy*t
	// xb(t) = bpx + bvx*t
	// yb(t) = bpy + bvy*t

	//We want to solve the system:
	// xa(t) = xb(t)
	// ya(t) = yb(t)
	// There is a unique solution if and only if the determinant of the matrix is not 0
	// Assuming that the determinant is not 0, we can solve the system with the cramer rule
	// t = ((bpy-apy)*bvx - (bpx-apx)*bvy) / (bvx*avy - bvy*avx)
	// s = ((bpy-apy)*avx - (bpx-apx)*avy) / (bvx*avy - bvy*avx)
	//

	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			var apx, apy, avx, avy float64 = float64(stones[i].px), float64(stones[i].py), float64(stones[i].vx), float64(stones[i].vy)
			var bpx, bpy, bvx, bvy float64 = float64(stones[j].px), float64(stones[j].py), float64(stones[j].vx), float64(stones[j].vy)
			//We use Cramer's rule to solve the linear equation system
			deno := bvx*avy - bvy*avx
			if deno == 0 {
				continue
			}
			var t = ((bpy-apy)*bvx - (bpx-apx)*bvy) / deno
			var s = ((bpy-apy)*avx - (bpx-apx)*avy) / deno
			if t < 0 || s < 0 {
				continue
			}
			//We check if the stones collide in the targetArea, so we move the stones to its position at time s
			var x = bpx + bvx*s
			var y = bpy + bvy*s
			if marginBottom < x && x < marginTop && marginBottom < y && y < marginTop {
				p1++
			}
		}
	}
	return p1
}
