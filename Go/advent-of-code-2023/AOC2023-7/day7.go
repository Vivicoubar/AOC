package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	value    string
	bid      int
	typeHand int
}

type ByTypeAndPower []Hand
type ByTypeAndPower2 []Hand

func main() {
	// Ouvrir le fichier en lecture
	fmt.Println("Part 1 Ex: ", P1("AOC2023-7/ex.txt"))
	fmt.Println("Part 1 : ", P1("AOC2023-7/input1.txt"))
	fmt.Println("Part 2 Ex: ", P2("AOC2023-7/ex.txt"))
	fmt.Println("Part 2 : ", P2("AOC2023-7/input1.txt"))
}

func P1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var allHands = make([]Hand, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var handInfo = strings.Split(line, " ")
		var handBid, _ = strconv.Atoi(handInfo[1])
		var hand Hand = Hand{handInfo[0], handBid, 0}
		allHands = append(allHands, hand)
	}
	for i, hand := range allHands {
		var allCard = make([]int, 14)
		for _, char := range hand.value {
			var card = string(char)
			if card == "A" {
				card = "1"
			}
			if card == "J" {
				card = "11"
			}
			if card == "Q" {
				card = "12"
			}
			if card == "K" {
				card = "13"
			}
			if card == "T" {
				card = "10"
			}

			var idCard, _ = strconv.Atoi(card)
			allCard[idCard] += 1
		}
		var countFive = 0
		var countFour = 0
		var countThree = 0
		var countTwo = 0
		for _, nbCard := range allCard {
			if nbCard == 0 {
				continue
			}
			if nbCard == 5 {
				countFive += 1
			}
			if nbCard == 4 {
				countFour += 1
			}
			if nbCard == 3 {
				countThree += 1
			}
			if nbCard == 2 {
				countTwo += 1
			}
		}
		if countFive == 1 {
			var newHand = Hand{hand.value, hand.bid, 7}
			allHands[i] = newHand
		} else if countFour == 1 {
			var newHand = Hand{hand.value, hand.bid, 6}
			allHands[i] = newHand
		} else if countThree == 1 && countTwo == 1 {
			var newHand = Hand{hand.value, hand.bid, 5}
			allHands[i] = newHand
		} else if countThree == 1 {
			var newHand = Hand{hand.value, hand.bid, 4}
			allHands[i] = newHand
		} else if countTwo == 2 {
			var newHand = Hand{hand.value, hand.bid, 3}
			allHands[i] = newHand
		} else if countTwo == 1 {
			var newHand = Hand{hand.value, hand.bid, 2}
			allHands[i] = newHand
		} else {
			var newHand = Hand{hand.value, hand.bid, 1}
			allHands[i] = newHand
		}
	}
	sort.Sort(ByTypeAndPower(allHands))
	var p1 = 0
	for i, hand := range allHands {
		p1 += (i + 1) * hand.bid
	}
	return p1
}

func P2(input string) int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var allHands = make([]Hand, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var handInfo = strings.Split(line, " ")
		var handBid, _ = strconv.Atoi(handInfo[1])
		var hand = Hand{handInfo[0], handBid, 0}
		allHands = append(allHands, hand)
	}
	for i, hand := range allHands {
		var allCard = make(map[string]int)
		var nbJok = 0
		for _, char := range hand.value {
			var card = string(char)
			if card == "J" {
				nbJok += 1
				continue
			}
			allCard[card] += 1
		}
		var occur = make([]int, 0, len(allCard))
		for _, nbCard := range allCard {
			occur = append(occur, nbCard)
		}
		sort.Ints(occur)
		if nbJok == 5 {
			var newHand = Hand{hand.value, hand.bid, 7}
			allHands[i] = newHand
			continue
		}
		if occur[len(occur)-1]+nbJok == 5 {
			var newHand = Hand{hand.value, hand.bid, 7}
			allHands[i] = newHand
		} else if occur[len(occur)-1]+nbJok == 4 {
			var newHand = Hand{hand.value, hand.bid, 6}
			allHands[i] = newHand
		} else if occur[len(occur)-1]+nbJok == 3 && len(occur) > 1 && occur[len(occur)-2] == 2 {
			var newHand = Hand{hand.value, hand.bid, 5}
			allHands[i] = newHand
		} else if occur[len(occur)-1]+nbJok == 3 {
			var newHand = Hand{hand.value, hand.bid, 4}
			allHands[i] = newHand
		} else if occur[len(occur)-1]+nbJok == 2 && len(occur) > 1 && occur[len(occur)-2] == 2 {
			var newHand = Hand{hand.value, hand.bid, 3}
			allHands[i] = newHand
		} else if occur[len(occur)-1]+nbJok == 2 {
			var newHand = Hand{hand.value, hand.bid, 2}
			allHands[i] = newHand
		} else {
			var newHand = Hand{hand.value, hand.bid, 1}
			allHands[i] = newHand
		}
	}
	sort.Sort(ByTypeAndPower2(allHands))
	var p2 = 0
	for i, hand := range allHands {
		p2 += (i + 1) * hand.bid
	}
	return p2
}

func (a ByTypeAndPower) Len() int      { return len(a) }
func (a ByTypeAndPower) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTypeAndPower) Less(i, j int) bool {
	// First, compare by type
	if a[i].typeHand != a[j].typeHand {
		return a[i].typeHand < a[j].typeHand
	}
	// If types are equal, compare by card power
	for k := 0; k < 5; k++ {
		if compareCardPower1(strings.Split(a[i].value, "")[k]) < compareCardPower1(strings.Split(a[j].value, "")[k]) {
			return true
		} else if compareCardPower1(strings.Split(a[i].value, "")[k]) > compareCardPower1(strings.Split(a[j].value, "")[k]) {
			return false
		}
	}
	return false
}

func compareCardPower1(cardValue string) int {
	if cardValue > "0" && cardValue <= "9" {
		card, err := strconv.Atoi(cardValue)
		if err != nil {
			fmt.Println("Error converting card value to integer:", err)
		}
		return card
	} else {
		if cardValue == "T" {
			return 10
		} else if cardValue == "J" {
			return 11
		} else if cardValue == "Q" {
			return 12
		} else if cardValue == "K" {
			return 13
		} else if cardValue == "A" {
			return 14
		}
	}
	return 0
}

func (a ByTypeAndPower2) Len() int      { return len(a) }
func (a ByTypeAndPower2) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTypeAndPower2) Less(i, j int) bool {
	// First, compare by type
	if a[i].typeHand != a[j].typeHand {
		return a[i].typeHand < a[j].typeHand
	}
	// If types are equal, compare by card power
	for k := 0; k < 5; k++ {
		if compareCardPower2(strings.Split(a[i].value, "")[k]) < compareCardPower2(strings.Split(a[j].value, "")[k]) {
			return true
		} else if compareCardPower2(strings.Split(a[i].value, "")[k]) > compareCardPower2(strings.Split(a[j].value, "")[k]) {
			return false
		}
	}
	return false
}

func compareCardPower2(cardValue string) int {
	if cardValue > "0" && cardValue <= "9" {
		card, err := strconv.Atoi(cardValue)
		if err != nil {
			fmt.Println("Error converting card value to integer:", err)
		}
		return card
	} else {
		if cardValue == "T" {
			return 10
		} else if cardValue == "J" {
			return 0
		} else if cardValue == "Q" {
			return 12
		} else if cardValue == "K" {
			return 13
		} else if cardValue == "A" {
			return 14
		}
	}
	return 1
}
