package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type transfo struct {
	destination int
	source      int
	length      int
}

type seedInt struct {
	start int
	end   int
}

func main() {
	fmt.Println(P1("AOC2023-5/ex.txt"))
	fmt.Println(P1("AOC2023-5/input1.txt"))
	fmt.Println(P2("AOC2023-5/ex.txt"))
	fmt.Println(P2("AOC2023-5/input1.txt"))
}

func P1(input string) int {
	// Ouvrir le fichier en lecture
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var originalSeeds = make([]int, 0)
	var transfoFromSoil = make([]transfo, 0)
	var transfoFromFerti = make([]transfo, 0)
	var transfoFromWater = make([]transfo, 0)
	var transfoFromLight = make([]transfo, 0)
	var transfoFromTempe = make([]transfo, 0)
	var transfoFromHumid = make([]transfo, 0)
	var transfoFromSeed = make([]transfo, 0)
	var currentTransfo []transfo
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var token = line[:5]
		if token == "seeds" {
			var seed = strings.Split(line[7:], " ")
			for _, s := range seed {
				var intSeed, _ = strconv.Atoi(s)
				originalSeeds = append(originalSeeds, intSeed)
			}
		} else if token == "seed-" {
			currentTransfo = transfoFromSeed
			continue
		} else if token == "soil-" {
			transfoFromSeed = currentTransfo
			currentTransfo = transfoFromSoil
			continue
		} else if token == "ferti" {
			transfoFromSoil = currentTransfo
			currentTransfo = transfoFromFerti
			continue
		} else if token == "water" {
			transfoFromFerti = currentTransfo
			currentTransfo = transfoFromWater
			continue
		} else if token == "light" {
			transfoFromWater = currentTransfo
			currentTransfo = transfoFromLight
			continue
		} else if token == "tempe" {
			transfoFromLight = currentTransfo
			currentTransfo = transfoFromTempe
			continue
		} else if token == "humid" {
			transfoFromTempe = currentTransfo
			currentTransfo = transfoFromHumid
			continue
		}
		var element = strings.Split(line, " ")
		var dest, _ = strconv.Atoi(element[0])
		var source, _ = strconv.Atoi(element[1])
		var numRange, _ = strconv.Atoi(element[2])
		var newTransfo = transfo{dest, source, numRange}
		currentTransfo = append(currentTransfo, newTransfo)
	}
	transfoFromHumid = currentTransfo
	var maps = make([][]transfo, 0)
	maps = append(maps, transfoFromSeed)
	maps = append(maps, transfoFromSoil)
	maps = append(maps, transfoFromFerti)
	maps = append(maps, transfoFromWater)
	maps = append(maps, transfoFromLight)
	maps = append(maps, transfoFromTempe)
	maps = append(maps, transfoFromHumid)
	var p1 = 10e9
	for _, seed := range originalSeeds {
		for _, m := range maps {
			for _, transf := range m {
				var offset = transf.destination - transf.source
				if seed >= transf.source && seed <= transf.source+transf.length {
					seed = seed + offset
					break
				}

			}
		}
		if float64(seed) < p1 {
			p1 = float64(seed)
		}
	}
	return int(p1)
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
	var originalSeeds = make([]seedInt, 0)
	var transfoFromSoil = make([]transfo, 0)
	var transfoFromFerti = make([]transfo, 0)
	var transfoFromWater = make([]transfo, 0)
	var transfoFromLight = make([]transfo, 0)
	var transfoFromTempe = make([]transfo, 0)
	var transfoFromHumid = make([]transfo, 0)
	var transfoFromSeed = make([]transfo, 0)
	var currentTransfo []transfo
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var token = line[:5]
		if token == "seeds" {
			var seed = strings.Split(line[7:], " ")
			for i := 0; i < len(seed) && i+1 < len(seed); i += 2 {
				var begin, _ = strconv.Atoi(seed[i])
				var rangeNum, _ = strconv.Atoi(seed[i+1])
				var newInt = seedInt{begin, begin + rangeNum}
				originalSeeds = append(originalSeeds, newInt)
			}
		} else if token == "seed-" {
			currentTransfo = transfoFromSeed
			continue
		} else if token == "soil-" {
			transfoFromSeed = currentTransfo
			currentTransfo = transfoFromSoil
			continue
		} else if token == "ferti" {
			transfoFromSoil = currentTransfo
			currentTransfo = transfoFromFerti
			continue
		} else if token == "water" {
			transfoFromFerti = currentTransfo
			currentTransfo = transfoFromWater
			continue
		} else if token == "light" {
			transfoFromWater = currentTransfo
			currentTransfo = transfoFromLight
			continue
		} else if token == "tempe" {
			transfoFromLight = currentTransfo
			currentTransfo = transfoFromTempe
			continue
		} else if token == "humid" {
			transfoFromTempe = currentTransfo
			currentTransfo = transfoFromHumid
			continue
		}
		var element = strings.Split(line, " ")
		var dest, _ = strconv.Atoi(element[0])
		var source, _ = strconv.Atoi(element[1])
		var numRange, _ = strconv.Atoi(element[2])
		var newTransfo = transfo{dest, source, numRange}
		currentTransfo = append(currentTransfo, newTransfo)
	}

	transfoFromHumid = currentTransfo
	var maps = make([][]transfo, 0)
	maps = append(maps, transfoFromSeed)
	maps = append(maps, transfoFromSoil)
	maps = append(maps, transfoFromFerti)
	maps = append(maps, transfoFromWater)
	maps = append(maps, transfoFromLight)
	maps = append(maps, transfoFromTempe)
	maps = append(maps, transfoFromHumid)
	var seed = originalSeeds
	for i := 0; i < len(maps); i++ {
		var newSeed = make([]seedInt, 0)
		for len(seed) > 0 {
			var begin, end = seed[0].start, seed[0].end
			seed = seed[1:]
			var mapped = false
			for _, transf := range maps[i] {
				var overlapStart = max(begin, transf.source)
				var overlapEnd = min(end, transf.source+transf.length)
				if overlapStart < overlapEnd {
					mapped = true
					newSeed = append(newSeed, seedInt{overlapStart + transf.destination - transf.source, overlapEnd + transf.destination - transf.source})
					if overlapStart > begin {
						seed = append(seed, seedInt{begin, overlapStart})
					}
					if overlapEnd < end {
						seed = append(seed, seedInt{overlapEnd, end})
					}
					break
				}
			}
			if !mapped {
				newSeed = append(newSeed, seedInt{begin, end})
			}
		}
		seed = newSeed
	}
	var minInt = seed[0]
	for _, s := range seed {
		if s.start < minInt.start {
			minInt = s
		}
	}
	var p2 = minInt.start
	return p2
}
