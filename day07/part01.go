// Credit goes here: I ended up implementing the logic of this
// https://github.com/AleUP170/AdventOfCode/blob/main/Day07/01.py

package day07

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
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

type Hand struct {
	cards string
	bid   int
}

type RankedHand struct {
	hand Hand
	rank int
}

var matches = [][]Hand{
	{}, {}, {}, {}, {}, {}, {},
}

func findMatches(hands []Hand) {
	for _, hand := range hands {
		count := make(map[rune]int)

		for _, card := range hand.cards {
			count[card]++
		}

		value := 1
		for _, c := range count {
			value *= c
		}

		switch value {
		case 1:
			matches[6] = append(matches[6], hand)
		case 2:
			matches[5] = append(matches[5], hand)
		case 3:
			matches[3] = append(matches[3], hand)
		case 4:
			if len(count) == 2 {
				matches[1] = append(matches[1], hand)
			} else {
				matches[4] = append(matches[4], hand)
			}
		case 5:
			matches[0] = append(matches[0], hand)
		case 6:
			matches[2] = append(matches[2], hand)
		default:
			fmt.Println("oh no")
		}
	}
}

func convertAndOrderMatches() []RankedHand {
	var convertedMatches []RankedHand

	for _, category := range matches {
		var temp []RankedHand

		for _, hand := range category {
			// Substituting T J Q K A with A B C D E for ordering later
			cards := regexp.MustCompile(`A`).ReplaceAllString(hand.cards, "E")
			cards = regexp.MustCompile(`T`).ReplaceAllString(cards, "A")
			cards = regexp.MustCompile(`J`).ReplaceAllString(cards, "B")
			cards = regexp.MustCompile(`Q`).ReplaceAllString(cards, "C")
			cards = regexp.MustCompile(`K`).ReplaceAllString(cards, "D")

			num, err := strconv.ParseInt(cards, 16, 0)
			if err != nil {
				fmt.Println("Error converting cards to number:", err)
				continue
			}

			temp = append(temp, RankedHand{hand: hand, rank: int(num)})
		}

		sort.Slice(temp, func(i, j int) bool {
			return temp[i].rank > temp[j].rank
		})

		convertedMatches = append(convertedMatches, temp...)
	}

	return convertedMatches
}

func Part01(input string) {
	lines := input
	hands := make([]Hand, 0)

	re := regexp.MustCompile(`[\dAKQJT]+`)
	bidRe := regexp.MustCompile(` [\d]+`)

	for _, line := range regexp.MustCompile(`\n`).Split(lines, -1) {
		if len(line) == 0 {
			continue
		}

		cards := re.FindString(line)
		bid, _ := strconv.Atoi(bidRe.FindString(line)[1:])

		hands = append(hands, Hand{cards: cards, bid: bid})
	}

	findMatches(hands)

	convertedMatches := convertAndOrderMatches()

	total := 0
	for i := 0; i < len(convertedMatches); i++ {
		total += convertedMatches[i].hand.bid * (len(convertedMatches) - i)
	}

	fmt.Println(total)
}
