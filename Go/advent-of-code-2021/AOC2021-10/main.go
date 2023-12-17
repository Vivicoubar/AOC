package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Ouvrir le fichier en lecture
	file, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var errorsChars []string = make([]string, 0)
	var middleScores []int = make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var errorChar, seq = findErrorChar(line)
		if errorChar != "" {
			errorsChars = append(errorsChars, errorChar)
		} else {
			middleScores = append(middleScores, process(seq))
		}
	}
	fmt.Println(errorsChars)
	var sum = 0
	for i := 0; i < len(errorsChars); i++ {
		if errorsChars[i] == "}" {
			sum += 1197
		} else if errorsChars[i] == "]" {
			sum += 57
		} else if errorsChars[i] == ")" {
			sum += 3
		} else if errorsChars[i] == ">" {
			sum += 25137
		}
	}

	//sort middleScores
	for i := 0; i < len(middleScores); i++ {
		for j := i + 1; j < len(middleScores); j++ {
			if middleScores[i] > middleScores[j] {
				middleScores[i], middleScores[j] = middleScores[j], middleScores[i]
			}
		}
	}
	fmt.Println(middleScores[len(middleScores)/2])
	fmt.Println(sum)
}

func process(seq []string) int {
	var score = 0
	for i := 0; i < len(seq); i++ {
		score *= 5
		if seq[i] == ")" {
			score += 1
		}
		if seq[i] == "]" {
			score += 2
		}
		if seq[i] == "}" {
			score += 3
		}
		if seq[i] == ">" {
			score += 4
		}
	}
	return score
}

func findErrorChar(line string) (string, []string) {
	var stack []string = make([]string, 0)
	for i := 0; i < len(line); i++ {
		if len(stack) == 0 {
			stack = append(stack, string(line[i]))
		} else {
			if string(line[i]) == ")" {
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
				} else {
					return string(line[i]), make([]string, 0)
				}
			} else if string(line[i]) == "]" {
				if stack[len(stack)-1] == "[" {
					stack = stack[:len(stack)-1]
				} else {
					return string(line[i]), make([]string, 0)
				}
			} else if string(line[i]) == "}" {
				if stack[len(stack)-1] == "{" {
					stack = stack[:len(stack)-1]
				} else {
					return string(line[i]), make([]string, 0)
				}
			} else if string(line[i]) == ">" {
				if stack[len(stack)-1] == "<" {
					stack = stack[:len(stack)-1]
				} else {
					return string(line[i]), make([]string, 0)
				}
			} else {
				stack = append(stack, string(line[i]))
			}
		}
	}
	var seq = make([]string, 0)
	// At this point the line is incorrect
	for i := len(stack) - 1; i > -1; i-- {
		if stack[i] == "(" {
			seq = append(seq, ")")
		} else if stack[i] == "[" {
			seq = append(seq, "]")
		} else if stack[i] == "{" {
			seq = append(seq, "}")
		} else if stack[i] == "<" {
			seq = append(seq, ">")
		}
	}
	return "", seq
}
