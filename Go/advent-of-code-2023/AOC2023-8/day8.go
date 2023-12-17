package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	value string
	left  string
	right string
}

func main() {
	fmt.Println("Part 1 EX1: ", P1("ex.txt"))
	fmt.Println("Part 1 EX2: ", P1("ex2.txt"))
	fmt.Println("Part 1 : ", P1("input1.txt"))
	fmt.Println("Part 2 EX: ", P2("ex3.txt"))
	fmt.Println("Part 2 : ", P2("input1.txt"))
}

func P1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count = 0
	var path = ""
	var beginNodes = make([]string, 0)
	var mapNodes = make(map[string]*Node)
	for scanner.Scan() {
		line := scanner.Text()
		if count == 0 {
			path = line
			count++
			continue
		}
		if count > 1 {
			var elements = strings.Split(line, " = ")
			var value = elements[0]
			var left = strings.Split(elements[1], ", ")[0][1:]
			var right = strings.Split(elements[1], ", ")[1][:len(strings.Split(elements[1], ", ")[1])-1]
			var inMapNode, exists = mapNodes[value]
			if value[2] == 'A' {
				beginNodes = append(beginNodes, value)
			}
			if !exists {
				var node = Node{value, left, right}
				mapNodes[value] = &node
			} else {
				var leftNode, existsLeft = mapNodes[left]
				if existsLeft {
					inMapNode.left = leftNode.value
				} else {
					var node = Node{left, "", ""}
					mapNodes[left] = &node
				}
				var rightNode, existsRight = mapNodes[right]
				if existsRight {
					inMapNode.right = rightNode.value
				} else {
					var node = Node{right, "", ""}
					mapNodes[right] = &node
				}
			}
		}
		count++
	}
	var mod = len(path)
	var pointer = 0
	var stepCounter = 0
	var curNode = mapNodes["AAA"]
	for curNode.value != "ZZZ" {
		var next = path[pointer]
		if next == 'L' {
			curNode = mapNodes[curNode.left]
		} else {
			curNode = mapNodes[curNode.right]
		}
		pointer = (pointer + 1) % mod
		stepCounter++
	}
	return stepCounter
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
	var count = 0
	var path = ""
	var beginNodes = make([]string, 0)
	var mapNodes = make(map[string]*Node)
	for scanner.Scan() {
		line := scanner.Text()
		if count == 0 {
			path = line
			count++
			continue
		}
		if count > 1 {
			var elements = strings.Split(line, " = ")
			var value = elements[0]
			var left = strings.Split(elements[1], ", ")[0][1:]
			var right = strings.Split(elements[1], ", ")[1][:len(strings.Split(elements[1], ", ")[1])-1]
			var inMapNode, exists = mapNodes[value]
			if value[2] == 'A' {
				beginNodes = append(beginNodes, value)
			}
			if !exists {
				var node = Node{value, left, right}
				mapNodes[value] = &node
			} else {
				var leftNode, existsLeft = mapNodes[left]
				if existsLeft {
					inMapNode.left = leftNode.value
				} else {
					var node = Node{left, "", ""}
					mapNodes[left] = &node
				}
				var rightNode, existsRight = mapNodes[right]
				if existsRight {
					inMapNode.right = rightNode.value
				} else {
					var node = Node{right, "", ""}
					mapNodes[right] = &node
				}
			}
		}
		count++
	}
	var mod = len(path)
	var pointer = 0
	var stepCounter = 0
	var curCounts = make([]int, 0)
	for _, beginNode := range beginNodes {
		var curNode = mapNodes[beginNode]
		pointer = 0
		var curCount = 0
		for curNode.value[2] != 'Z' {
			var next = path[pointer]
			if next == 'L' {
				curNode = mapNodes[curNode.left]
			} else {
				curNode = mapNodes[curNode.right]
			}
			pointer = (pointer + 1) % mod
			curCount++
		}
		stepCounter++
		curCounts = append(curCounts, curCount)
	}
	//Find least common multiple
	var lcm = curCounts[0]
	for i := 1; i < len(curCounts); i++ {
		lcm = lcm * curCounts[i] / gcd(lcm, curCounts[i])
	}
	return lcm
}

func gcd(lcm int, i int) int {
	for i != 0 {
		lcm, i = i, lcm%i
	}
	return lcm
}
