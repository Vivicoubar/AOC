package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type P2HailStone struct {
	Position [3]int
	Velocity [3]int
}

func positionThrowObliterate(hailStones []P2HailStone) int {
	// We only need three stones to obliterate all the others
	// We can use the first three stones to solve a system of linear equations
	// We can then use the solution to find the time at which the stones will be aligned
	// We can then use the time to find the position of the stones at that time
	stones := hailStones[:3]

	equationMatrix := make([][]float64, 6)
	for i := range equationMatrix {
		equationMatrix[i] = make([]float64, 6)
	}

	vector := make([]float64, 6)

	setMatrixValue := func(i, j int, value float64) {
		equationMatrix[i][j] = value
	}

	//We create the matrix of the linear system
	// if the stone has pos p0 and vel v0
	// and every hailstone i has pos pi and vel vi
	// then p0 + t[i]*v0 == p[i] + t[i]*v[i]
	// which is equivalent to p0 - p[i] == t[i]*(v[i] - v0)
	// so we have 3 equations for each stone
	// we can write the system as

	setMatrixValue(0, 1, float64(stones[0].Velocity[2]-stones[1].Velocity[2]))
	setMatrixValue(0, 2, float64(stones[1].Velocity[1]-stones[0].Velocity[1]))
	setMatrixValue(0, 4, float64(stones[1].Position[2]-stones[0].Position[2]))
	setMatrixValue(0, 5, float64(stones[0].Position[1]-stones[1].Position[1]))

	setMatrixValue(1, 0, float64(stones[1].Velocity[2]-stones[0].Velocity[2]))
	setMatrixValue(1, 2, float64(stones[0].Velocity[0]-stones[1].Velocity[0]))
	setMatrixValue(1, 3, float64(stones[0].Position[2]-stones[1].Position[2]))
	setMatrixValue(1, 5, float64(stones[1].Position[0]-stones[0].Position[0]))

	setMatrixValue(2, 0, float64(stones[0].Velocity[1]-stones[1].Velocity[1]))
	setMatrixValue(2, 1, float64(stones[1].Velocity[0]-stones[0].Velocity[0]))
	setMatrixValue(2, 3, float64(stones[1].Position[1]-stones[0].Position[1]))
	setMatrixValue(2, 4, float64(stones[0].Position[0]-stones[1].Position[0]))

	setMatrixValue(3, 1, float64(stones[0].Velocity[2]-stones[2].Velocity[2]))
	setMatrixValue(3, 2, float64(stones[2].Velocity[1]-stones[0].Velocity[1]))
	setMatrixValue(3, 4, float64(stones[2].Position[2]-stones[0].Position[2]))
	setMatrixValue(3, 5, float64(stones[0].Position[1]-stones[2].Position[1]))

	setMatrixValue(4, 0, float64(stones[2].Velocity[2]-stones[0].Velocity[2]))
	setMatrixValue(4, 2, float64(stones[0].Velocity[0]-stones[2].Velocity[0]))
	setMatrixValue(4, 3, float64(stones[0].Position[2]-stones[2].Position[2]))
	setMatrixValue(4, 5, float64(stones[2].Position[0]-stones[0].Position[0]))

	setMatrixValue(5, 0, float64(stones[0].Velocity[1]-stones[2].Velocity[1]))
	setMatrixValue(5, 1, float64(stones[2].Velocity[0]-stones[0].Velocity[0]))
	setMatrixValue(5, 3, float64(stones[2].Position[1]-stones[0].Position[1]))
	setMatrixValue(5, 4, float64(stones[0].Position[0]-stones[2].Position[0]))

	indepX0 := float64(stones[0].Position[1]*stones[0].Velocity[2] - stones[0].Velocity[1]*stones[0].Position[2])
	indepX1 := float64(stones[1].Position[1]*stones[1].Velocity[2] - stones[1].Velocity[1]*stones[1].Position[2])
	indepX2 := float64(stones[2].Position[1]*stones[2].Velocity[2] - stones[2].Velocity[1]*stones[2].Position[2])

	indepY0 := float64(stones[0].Position[2]*stones[0].Velocity[0] - stones[0].Velocity[2]*stones[0].Position[0])
	indepY1 := float64(stones[1].Position[2]*stones[1].Velocity[0] - stones[1].Velocity[2]*stones[1].Position[0])
	indepY2 := float64(stones[2].Position[2]*stones[2].Velocity[0] - stones[2].Velocity[2]*stones[2].Position[0])

	indepZ0 := float64(stones[0].Position[0]*stones[0].Velocity[1] - stones[0].Velocity[0]*stones[0].Position[1])
	indepZ1 := float64(stones[1].Position[0]*stones[1].Velocity[1] - stones[1].Velocity[0]*stones[1].Position[1])
	indepZ2 := float64(stones[2].Position[0]*stones[2].Velocity[1] - stones[2].Velocity[0]*stones[2].Position[1])

	vector[0] = indepX0 - indepX1
	vector[1] = indepY0 - indepY1
	vector[2] = indepZ0 - indepZ1
	vector[3] = indepX0 - indepX2
	vector[4] = indepY0 - indepY2
	vector[5] = indepZ0 - indepZ2

	result := solveLinearEquations(equationMatrix, vector)
	return int(result[0] + result[1] + result[2])
}

func solveLinearEquations(matrix [][]float64, vector []float64) []float64 {
	// Perform Gaussian elimination
	n := len(matrix)
	for i := 0; i < n; i++ {
		// Find pivot row
		pivot := i
		for j := i + 1; j < n; j++ {
			if abs(matrix[j][i]) > abs(matrix[pivot][i]) {
				pivot = j
			}
		}

		// Swap rows
		matrix[i], matrix[pivot] = matrix[pivot], matrix[i]
		vector[i], vector[pivot] = vector[pivot], vector[i]

		// Make the diagonal elements 1
		scale := 1 / matrix[i][i]
		for j := i; j < n; j++ {
			matrix[i][j] *= scale
		}
		vector[i] *= scale

		// Eliminate other rows
		for j := 0; j < n; j++ {
			if j != i {
				scale := matrix[j][i]
				for k := i; k < n; k++ {
					matrix[j][k] -= scale * matrix[i][k]
				}
				vector[j] -= scale * vector[i]
			}
		}
	}

	return vector
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func P2(input string) int {
	filename := input
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	var hailStones []P2HailStone
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var stone P2HailStone
		var elem = strings.Split(line, " @ ")
		var posElem = strings.Split(elem[0], ",")
		for i := 0; i < len(posElem); i++ {
			posElem[i] = strings.TrimSpace(posElem[i])
		}
		var velElem = strings.Split(elem[1], ", ")
		for i := 0; i < len(velElem); i++ {
			velElem[i] = strings.TrimSpace(velElem[i])
		}
		var position [3]int
		position[0], _ = strconv.Atoi(posElem[0])
		position[1], _ = strconv.Atoi(posElem[1])
		position[2], _ = strconv.Atoi(posElem[2])
		var velocity [3]int
		velocity[0], _ = strconv.Atoi(velElem[0])
		velocity[1], _ = strconv.Atoi(velElem[1])
		velocity[2], _ = strconv.Atoi(velElem[2])
		stone.Position = position
		stone.Velocity = velocity
		hailStones = append(hailStones, stone)
	}

	position := positionThrowObliterate(hailStones)

	return position
}
