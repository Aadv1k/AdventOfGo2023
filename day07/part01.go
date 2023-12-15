package day07

import (
	"log"
	"strings"
	"github.com/aadv1k/AdventOfGo2023/utils"
)

const (
	HighCard = iota + 1
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

var Strengths = []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

func indexOf(element byte, array []byte) int {
	for i := 0; i < len(array); i++ {
		if array[i] == element {
			return i
		}
	}
	return -1
}

func CompareHandsGT(hand1, hand2 string) bool {
	for i := range hand1 {
		if indexOf(hand1[i], Strengths) >= indexOf(hand2[i], Strengths) {
			return true
		}
	}
	return false
}

func GetHandType(hand string) int {
	letters := make(map[rune]int)
	frequency := make([]int, 6)

	for i := range hand {
		letters[rune(hand[i])]++
	}

	for i := range letters {
		frequency[letters[i]]++
	}

	if frequency[5] == 1 { // the number 5 appeared once
		return FiveKind
	} else if frequency[4] == 1 && frequency[1] == 1 { // the number 4 appeared once, and 1 once
		return FourKind
	} else if frequency[2] == 1 && frequency[3] == 1 {
		return FullHouse
	} else if frequency[3] == 1 && frequency[1] == 2 {
		return ThreeKind
	} else if frequency[2] == 2 {
		return TwoPair
	} else {
		return HighCard
	}
}

func Part01(input string) {
	lines := strings.Split(input, "\n")

	var hands []string
	winnings := make(map[string]int)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		word := strings.Fields(line)

		hands = append(hands, word[0])
		winnings[word[0]] = utils.ParseInt(word[1])

		rank := GetHandType(word[0])
		
		log.Printf("%s -> %d", word[0], rank)
	}

	_ = hands
}
