package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//Observation:
// 1. The grid is a square
// 2. The edge of the grid is always empty
// 3. There is an alternate for step seen at least one time
//(one time they are reached, one time they are not).
// 4. THERE IS A REPETITION OF THE GRID; THIS IS A PAVEMENT
// 5. Line S and Row S are always empty, so we can always go through theses lines and then finish in the reached grid
// 6. If a grid is on a edge, we can add how many times we can go through the pavement and have the same behaviour than the edge grid.
// 7. Same with the Pavement

type GridPos struct {
	gridRow int
	gridCol int
	row     int
	col     int
}

type seenState struct {
	doubleMargin bool
	step         int
	maxIteration int
}

type GridState struct {
	gridPos GridPos
	step    int
}

func P2(input string) int {
	// We read the input
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
	//We can count on a 7*7 grid, and for the corners we can use Obs 6/7
	var maxTranslate = 3
	var maxIter = 26501365
	var maxR = len(garden)
	var maxC = len(garden[0])
	var BFSGrid = BFS(garden, GridPos{0, 0, startPos.row, startPos.col}, maxR, maxC, maxTranslate)
	var i = 0
	var possibleGrid = make([]int, 0)
	var seen = make(map[seenState]int)
	for j := -maxTranslate + 1; j < maxTranslate; j++ {
		possibleGrid = append(possibleGrid, j)
	}
	for row := 0; row < maxR; row++ {
		for col := 0; col < maxC; col++ {
			var _, ok = BFSGrid[GridPos{0, 0, row, col}]
			if ok {
				for _, rowPG := range possibleGrid {
					for _, colPG := range possibleGrid {
						var step = BFSGrid[GridPos{rowPG, colPG, row, col}]
						if step%2 == maxIter%2 && step <= maxIter {
							//If we can reach the pos
							i++
						}
						if (rowPG == maxTranslate-1 || rowPG == -maxTranslate+1) && (colPG == -maxTranslate+1 || colPG == maxTranslate-1) {
							i += findSteps(seen, step, true, maxIter, maxR)
						} else if rowPG == maxTranslate-1 || rowPG == -maxTranslate+1 || colPG == maxTranslate-1 || colPG == -maxTranslate+1 {
							i += findSteps(seen, step, false, maxIter, maxR)
						}
					}
				}
			}
		}
	}
	return i
}

func BFS(garden [][]string, pos GridPos, maxR int, maxCol int, maxTranslate int) map[GridPos]int {
	var BFSDic = make(map[GridPos]int)
	var frontier = make([]GridState, 0)
	frontier = append(frontier, GridState{pos, 0})
	for len(frontier) > 0 {
		var current = frontier[0]
		frontier = frontier[1:]
		var nextGr, nextGc, nextR, nextC int
		nextGr = current.gridPos.gridRow
		nextGc = current.gridPos.gridCol
		nextR = current.gridPos.row
		nextC = current.gridPos.col
		if nextR < 0 {
			nextGr--
			nextR += maxR
		}
		if nextR >= maxR {
			nextGr++
			nextR -= maxR
		}
		if nextC < 0 {
			nextGc--
			nextC += maxCol
		}
		if nextC >= maxCol {
			nextGc++
			nextC -= maxCol
		}
		if !(nextR >= 0 && nextR < maxR && nextC >= 0 && nextC < maxCol && garden[nextR][nextC] != "#") {
			continue
		}
		if _, ok := BFSDic[GridPos{nextGr, nextGc, nextR, nextC}]; ok {
			continue
		}
		if math.Abs(float64(nextGr)) > float64(maxTranslate) || math.Abs(float64(nextGc)) > float64(maxTranslate) {
			continue
		}
		BFSDic[GridPos{nextGr, nextGc, nextR, nextC}] = current.step
		frontier = append(frontier, GridState{GridPos{nextGr, nextGc, nextR + 1, nextC}, current.step + 1})
		frontier = append(frontier, GridState{GridPos{nextGr, nextGc, nextR - 1, nextC}, current.step + 1})
		frontier = append(frontier, GridState{GridPos{nextGr, nextGc, nextR, nextC + 1}, current.step + 1})
		frontier = append(frontier, GridState{GridPos{nextGr, nextGc, nextR, nextC - 1}, current.step + 1})
	}
	return BFSDic
}

func findSteps(seen map[seenState]int, step int, isDouble bool, maxIter int, maxR int) int {
	if seenValue, seenOk := seen[seenState{isDouble, step, maxIter}]; seenOk {
		return seenValue
	} else {
		var value = 0
		var maxLoop int = (maxIter-step)/maxR + 1 //Number of times we cross a full pavement
		for loop := 1; loop < maxLoop; loop++ {
			if step+loop*maxR <= maxIter && (step+loop*maxR)%2 == maxIter%2 {
				//Find how many tiles can be reached in the pavement
				if isDouble {
					//Loop + 1 because we can go in 2 directions
					value += loop + 1
				} else {
					//1 car we only go in one direction
					value += 1

				}

			}
		}
		seen[seenState{isDouble, step, maxIter}] = value
		return value
	}

}
