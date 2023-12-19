package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	checkedParameter string
	greater          bool
	checkedValue     int
	trueLabel        string
}

type WorkFlow struct {
	rules     []Rule
	labelElse string
}

type Scrap struct {
	x int
	m int
	a int
	s int
}

func main() {
	fmt.Println("Partie 1 :", P1("AOC2023-19/ex.txt"))
	fmt.Println("Partie 1 :", P1("AOC2023-19/input1.txt"))
	fmt.Println("Partie 2 :", P2("AOC2023-19/ex.txt"))
	fmt.Println("Partie 2 :", P2("AOC2023-19/input1.txt"))

}

func inputToWorkFlowMap(input string) (map[string]WorkFlow, []Scrap) {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil, nil
	}
	defer file.Close()

	var workFlowMap = make(map[string]WorkFlow)
	var scraps = make([]Scrap, 4)
	var addingRules = true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			addingRules = false
			continue
		}
		line := scanner.Text()
		if addingRules {
			elements := strings.Split(line, "{")
			key := elements[0]
			workFlowElements := elements[1][:len(elements[1])-1]
			elements = strings.Split(workFlowElements, ",")
			var labelElse = elements[len(elements)-1]
			var workFlow = WorkFlow{[]Rule{}, labelElse}
			for _, element := range elements[:len(elements)-1] {
				rulesElement := strings.Split(element, ":")
				var trueLabel = rulesElement[len(rulesElement)-1]
				ruleCondition := rulesElement[0]
				var checkedParamater = ruleCondition[0]
				var greater = ruleCondition[1] == '>'
				var checkedValue, _ = strconv.Atoi(ruleCondition[2:])
				var rule = Rule{string(checkedParamater), greater, checkedValue, trueLabel}
				workFlow.rules = append(workFlow.rules, rule)
			}
			workFlowMap[key] = workFlow
		} else {
			line = line[1 : len(line)-1]
			elements := strings.Split(line, ",")
			var intElements = make([]int, 4)
			for i, element := range elements {
				var numElement, _ = strconv.Atoi(element[2:])
				intElements[i] = numElement
			}
			var scrap = Scrap{intElements[0], intElements[1], intElements[2], intElements[3]}
			scraps = append(scraps, scrap)
		}
	}
	return workFlowMap, scraps
}

func P1(input string) int {
	workFlowMap, scraps := inputToWorkFlowMap(input)
	var acceptedScraps = make([]Scrap, 0)
	var p1 = 0
	for _, scrap := range scraps {
		if scrap.x == 0 && scrap.m == 0 && scrap.a == 0 && scrap.s == 0 {
			continue
		}
		if isAcceptedScrap(scrap, workFlowMap) {
			acceptedScraps = append(acceptedScraps, scrap)
			p1 += sumParts(scrap)
		}
	}
	return p1
}

func isAcceptedScrap(scrap Scrap, flowMap map[string]WorkFlow) bool {
	var label = "in"
	for label != "A" && label != "R" {
		var workFlow = flowMap[label]
		var rules = workFlow.rules
		var hasTrueLabel = false
		for checkedRule := range rules {
			var rule = rules[checkedRule]
			var checkedValue = scrap.x
			switch rule.checkedParameter {
			case "m":
				checkedValue = scrap.m
			case "a":
				checkedValue = scrap.a
			case "s":
				checkedValue = scrap.s
			}
			if rule.greater {
				if checkedValue > rule.checkedValue {
					label = rule.trueLabel
					hasTrueLabel = true
					break
				}
			} else {
				if checkedValue < rule.checkedValue {
					label = rule.trueLabel
					hasTrueLabel = true
					break
				}
			}
		}
		if !hasTrueLabel {
			label = workFlow.labelElse
		}
	}
	return label == "A"
}

func sumParts(scrap Scrap) int {
	return scrap.x + scrap.m + scrap.a + scrap.s
}
