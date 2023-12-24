package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Module struct {
	label      string
	moduleType string
	output     []string
	data       map[string]string
}

type ModAndSig struct {
	input       Module
	destination Module
	signal      string
}

func main() {
	fmt.Println("P1: Ex1", P1("AOC2023-20/ex.txt"))
	fmt.Println("P1: Ex2", P1("AOC2023-20/ex2.txt"))
	fmt.Println("P1:", P1("AOC2023-20/input1.txt"))
	fmt.Println("P2:", P2("AOC2023-20/input1.txt"))
}

func P1(input string) int {
	return pushButton(getModules(input), 1000, nil)
}

func P2(input string) int {
	var modules = getModules(input)
	//Looking at the input, we can see that rs module is the only one to call the rx module
	//We also see that it's a Conjunction module
	//We can then deduce that the rx module is called only when all the inputs of the conjunction are low
	// -> PPCM (Again :( )
	var data = modules["rs"].data["input"]
	var toCheck = make([]Module, 0)
	for _, moduleLabel := range strings.Split(data, "-") {
		var module = modules[strings.Split(moduleLabel, ":")[0]]
		toCheck = append(toCheck, module)
	}
	return pushButton(modules, 1e10, toCheck)
}

func getModules(input string) map[string]Module {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data = make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	var modules map[string]Module = make(map[string]Module)
	for _, line := range data {
		var moduleElements = strings.Split(line, " -> ")
		var moduleLabel = moduleElements[0][1:len(moduleElements[0])]
		var moduleType = string(moduleElements[0][0])
		if moduleElements[0] == "broadcaster" {
			moduleLabel = "broadcaster"
			moduleType = "broadcaster"
		}
		var moduleOutput = strings.Split(moduleElements[1], ", ")
		var moduleData = make(map[string]string)
		var module = Module{moduleLabel, moduleType, moduleOutput, moduleData}
		modules[moduleLabel] = module
	}
	var button = Module{"button", "button", []string{"broadcast"}, make(map[string]string)}
	modules["button"] = button
	//We read another time the file to update the data of the conjunctions modules
	for _, module := range modules {
		if module.moduleType == "&" {
			var label = module.label
			for _, targetModule := range modules {
				for _, output := range targetModule.output {
					if output == label {
						var storedInput = module.data["input"]
						var storedModules = strings.Split(storedInput, "-")
						var found = false
						for i, storedModule := range storedModules {
							var storedModuleLabel = strings.Split(storedModule, ":")[0]
							if storedModuleLabel == targetModule.label {
								found = true
								storedModules[i] = targetModule.label + ":low"
								break
							}
						}
						if !found {
							if storedModules[0] == "" {
								storedModules = make([]string, 0)
							}
							storedModules = append(storedModules, targetModule.label+":low")
						}
						var newInput = strings.Join(storedModules, "-")
						module.data["input"] = newInput
					}
				}
			}
		}
	}
	return modules
}

func pushButton(modules map[string]Module, count int, toCheck []Module) int {
	var ppcmData = make(map[string]int)
	var foundData = 0
	var periodMap = make(map[string]int)
	var appearMap = make(map[string]int)
	var numberHigh = 0
	var numberLow = 0
	var moduleStack []ModAndSig = make([]ModAndSig, 0)
	for i := 0; i < count; i++ {
		moduleStack = append(moduleStack, ModAndSig{modules["button"], modules["broadcaster"], "low"})
		for len(moduleStack) > 0 {
			var input = moduleStack[0].input
			var module = moduleStack[0].destination
			var signal = moduleStack[0].signal

			//FOR PART2
			if toCheck != nil {
				var found = false
				var moduleToBeChecked Module
				for _, moduleToCheck := range toCheck {
					if moduleToCheck.label == module.label {
						found = true
						moduleToBeChecked = moduleToCheck
						break
					}
				}
				if found && signal == "low" {
					var _, appearOk = appearMap[moduleToBeChecked.label]
					var period, periodOk = periodMap[moduleToBeChecked.label]
					if appearOk && periodOk {
						ppcmData[moduleToBeChecked.label] = i - period
						foundData++
					}
					periodMap[moduleToBeChecked.label] = i
					appearMap[moduleToBeChecked.label] += 1
					if foundData == len(toCheck) {
						var values = make([]int, 0)
						for _, value := range ppcmData {
							values = append(values, value)
						}
						var ppcm = ppcmCalc(values)
						return ppcm
					}
				}

			}

			if signal == "high" {
				numberHigh++
			} else {
				numberLow++
			}
			moduleStack = moduleStack[1:]
			if module.moduleType == "%" {
				//FLIP FLOP
				var state, ok = module.data["state"]
				if !ok {
					module.data["state"] = "off"
					state = "off"
				}
				if signal != "high" {
					if state == "off" {
						module.data["state"] = "on"
						for _, output := range module.output {
							var newSignal = "high"
							//var log = module.label + "-" + newSignal + " -> " + output
							//fmt.Println(log)
							moduleStack = append(moduleStack, ModAndSig{module, modules[output], newSignal})
						}
						modules[module.label] = module
					} else {
						module.data["state"] = "off"
						for _, output := range module.output {
							var newSignal = "low"
							//var log = module.label + "-" + newSignal + " -> " + output
							//fmt.Println(log)
							moduleStack = append(moduleStack, ModAndSig{module, modules[output], newSignal})
						}
						modules[module.label] = module
					}
				}

			} else if module.moduleType == "&" {
				//Conjunction
				//Update input storage
				var storedInput, ok = module.data["input"]
				if !ok {
					module.data["input"] = ""
					storedInput = ""
				} else {
					storedInput = module.data["input"]
				}
				var storedModules = strings.Split(storedInput, "-")
				var found = false
				for i, storedModule := range storedModules {
					var storedModuleLabel = strings.Split(storedModule, ":")[0]
					if storedModuleLabel == input.label {
						found = true
						storedModules[i] = input.label + ":" + signal
						break
					}
				}
				if !found {
					if storedModules[0] == "" {
						storedModules = make([]string, 0)
					}
					storedModules = append(storedModules, input.label+":"+signal)
				}
				var newInput = strings.Join(storedModules, "-")
				module.data["input"] = newInput
				//Check if all inputs are high
				var allHigh = true
				for _, storedModule := range storedModules {
					if strings.Split(storedModule, ":")[1] == "low" {
						allHigh = false
						break
					}
				}
				if allHigh {
					for _, output := range module.output {
						var newSignal = "low"
						//var log = module.label + "-" + newSignal + " -> " + output
						//fmt.Println(log)
						moduleStack = append(moduleStack, ModAndSig{module, modules[output], newSignal})
					}
					modules[module.label] = module
				} else {
					for _, output := range module.output {
						var newSignal = "high"
						//var log = module.label + "-" + newSignal + " -> " + output
						//fmt.Println(log)
						moduleStack = append(moduleStack, ModAndSig{module, modules[output], newSignal})
					}
					modules[module.label] = module
				}
			} else if module.label == "broadcaster" {
				for _, output := range module.output {
					//var log = module.label + "-" + newSignal + " -> " + output
					//fmt.Println(log)
					moduleStack = append(moduleStack, ModAndSig{module, modules[output], signal})
				}
			}
		}
	}
	return (numberHigh * numberLow)
}

func ppcmCalc(values []int) int {
	var ppcm = 1
	for _, value := range values {
		ppcm = ppcm * value / gcd(ppcm, value)
	}
	return ppcm

}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
