package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type config struct {
	pos   int
	group int
	len   int
}

var configMap = make(map[config]int)

func main() {
	fmt.Println(P1("AOC2023-12/ex.txt"))
	fmt.Println(P1("AOC2023-12/input1.txt"))
	fmt.Println(P2("AOC2023-12/ex.txt"))
	fmt.Println(P2("AOC2023-12/input1.txt"))
}

func P1(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	var lines = make([]string, 0)
	var groups = make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//Split each line into a slice of strings with " " as separator
		var curLine = strings.Split(line, " ")[0]
		lines = append(lines, curLine)
		var curGroup = strings.Split(strings.Split(line, " ")[1], ",")
		var curGroupNum = make([]int, 0)
		for i := 0; i < len(curGroup); i++ {
			var num, _ = strconv.Atoi(curGroup[i])
			curGroupNum = append(curGroupNum, num)
		}
		groups = append(groups, curGroupNum)
	}
	var sum = 0
	for i := 0; i < len(lines); i++ {
		configMap = make(map[config]int)
		sum += countPossible(lines[i], groups[i], 0, 0, 0)
	}
	return sum
}

func P2(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	var lines = make([]string, 0)
	var groups = make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//Split each line into a slice of strings with " " as separator
		var curLine = strings.Split(line, " ")[0]
		lines = append(lines, curLine)
		var curGroup = strings.Split(line, " ")[1]
		var group = curGroup
		for i := 0; i < 4; i++ {
			curGroup = curGroup + "," + group
		}
		var newCurGroup = strings.Split(curGroup, ",")
		var curGroupNum = make([]int, 0)
		for i := 0; i < len(newCurGroup); i++ {
			var num, _ = strconv.Atoi(newCurGroup[i])
			curGroupNum = append(curGroupNum, num)
		}
		groups = append(groups, curGroupNum)
	}
	for i := 0; i < len(lines); i++ {
		var curLine = lines[i]
		for j := 0; j < 4; j++ {
			lines[i] = lines[i] + "?" + curLine
		}
	}
	var sum = 0
	for i := 0; i < len(lines); i++ {
		configMap = make(map[config]int)
		sum += countPossible(lines[i], groups[i], 0, 0, 0)
	}
	return sum
}

// Recursive function to count the number of possible combinations
func countPossible(points string, groups []int, pos int, group int, curLen int) int {
	var possibleComb = 0
	//Trying an optimisation with configMap
	var curConfig = config{pos, group, curLen}
	//If we already have the configuration in the map, we return it directly so we don't have to calculate it again
	if val, ok := configMap[curConfig]; ok {
		return val
	}
	//If we are at the end of the string, this is our close condition
	if pos == len(points) {
		if group == len(groups) && curLen == 0 {
			//If all groups have been found (we'll build them with the correct length) and we are not in a group, we count it
			return 1
		} else if group == len(groups)-1 && curLen == groups[group] {
			//If the last one is finishing, and we have the right length, we count it
			return 1
		} else {
			//The group is not finished, therefore not a valid configuration, we do not count it
			return 0
		}
	}
	// We will try to put a '.' at the current position
	if string(points[pos]) == "?" || string(points[pos]) == string('.') {
		//If we are not in a group, we can continue
		if curLen == 0 {
			possibleComb += countPossible(points, groups, pos+1, group, curLen)
			// If we are at the end of the group, and it as the correct length, we restart a new group
		} else if group < len(groups) && groups[group] == curLen {
			possibleComb += countPossible(points, groups, pos+1, group+1, 0)
			// If we are in a group, we continue the group
		}
	}
	// We will try to put a '#' at the current position
	if string(points[pos]) == "?" || string(points[pos]) == string('#') {
		//Continue the group we are in
		possibleComb += countPossible(points, groups, pos+1, group, curLen+1)
	}

	//Save the configuration in the map
	configMap[curConfig] = possibleComb
	return possibleComb
}
