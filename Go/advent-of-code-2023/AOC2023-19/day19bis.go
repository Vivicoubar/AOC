package main

import (
	"strconv"
)

type ScrapInterval struct {
	low  int
	high int
}

type State struct {
	label              string
	parametersInterval map[string]ScrapInterval
}

func P2(input string) int {
	var workFlowMap, _ = inputToWorkFlowMap(input)
	return dynamicProgramming("in", map[string]ScrapInterval{"x": ScrapInterval{1, 4000}, "m": ScrapInterval{1, 4000}, "a": ScrapInterval{1, 4000}, "s": ScrapInterval{1, 4000}}, make(map[string]int), workFlowMap)
}

// We use dynamic programming to solve this problem, which is a recursive function.
func dynamicProgramming(label string, parametersInterval map[string]ScrapInterval, memory map[string]int, workflows map[string]WorkFlow) int {
	if label == "A" {
		//If we are in the accept state, we return the number of possible scraps.
		//Which can be calculated by multiplying the number of possible values for each parameter.
		var out = 1
		for _, interval := range parametersInterval {
			out *= interval.high - interval.low + 1
		}
		return out
	} else if label == "R" {
		//If we are in the reject state, we return 0, since there is no possible scraps.
		return 0
	}
	state := State{label, parametersInterval}
	if value, ok := memory[encodeState(state)]; ok {
		//If we already calculated the number of possible scraps for this state, we return it.
		return value
	}
	var newConstraints = parametersInterval
	var workFlow = workflows[label]
	var out = 0
	for _, rule := range workFlow.rules {
		//We apply each rule to the current state
		//And we calculate the number of possible scraps for each new state.
		var newLabel = rule.trueLabel
		out += dynamicProgramming(newLabel, applyRule(newConstraints, rule), memory, workflows)
		//We also apply the inverted rule to the current state to calculate the number of possible scraps for each new state.
		//We need to do this to take into account the case where the rule is not applied.
		newConstraints = applyRule(newConstraints, invertRule(rule))
	}
	//We also need to take into account the case where no rule is applied.
	out += dynamicProgramming(workFlow.labelElse, newConstraints, memory, workflows)
	memory[encodeState(state)] = out
	return out
}

func invertRule(rule Rule) Rule {
	var greater = rule.greater
	if greater {
		return Rule{rule.checkedParameter, false, rule.checkedValue + 1, rule.trueLabel}
	} else {
		return Rule{rule.checkedParameter, true, rule.checkedValue - 1, rule.trueLabel}
	}
}

func encodeState(state State) string {
	var out = state.label + "/"
	for parameter, interval := range state.parametersInterval {
		out += parameter + "_" + strconv.Itoa(interval.low) + "_" + strconv.Itoa(interval.high) + "-"
	}
	return out[:len(out)-1]
}

func applyRule(parametersInterval map[string]ScrapInterval, rule Rule) map[string]ScrapInterval {
	var newInterval = make(map[string]ScrapInterval)
	var checkedParameter = rule.checkedParameter
	var checkedValue = rule.checkedValue
	var greater = rule.greater
	var low = parametersInterval[checkedParameter].low
	var high = parametersInterval[checkedParameter].high
	if greater {
		if checkedValue+1 >= low {
			//We need to add 1 to the checked value because the rule is applied only if the parameter is strictly greater than the checked value.
			low = checkedValue + 1
		}
	} else {
		if checkedValue-1 <= high {
			//We need to remove 1 to the checked value because the rule is applied only if the parameter is strictly lower than the checked value.
			high = checkedValue - 1
		}
	}
	for parameter, intervals := range parametersInterval {
		if parameter == checkedParameter {
			newInterval[parameter] = ScrapInterval{low, high}
		} else {
			newInterval[parameter] = intervals
		}
	}
	return newInterval
}
