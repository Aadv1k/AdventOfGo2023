package day07

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var valueDict = map[rune]int{'J': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'Q': 11, 'K': 12, 'A': 13}

func Part02(input string) {
	lines := strings.Split(input, "\n")

	hands := make([]Hand, 0)

	re := regexp.MustCompile(`[\dAKQJT]+`)
	bidRe := regexp.MustCompile(` [\d]+`)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		cards := re.FindString(line)
		bid, _ := strconv.Atoi(bidRe.FindString(line)[1:])

		hands = append(hands, Hand{cards: cards, bid: bid})
	}

	matches := [][]Hand{{}, {}, {}, {}, {}, {}, {}}

	for _, hand := range hands {
		count := make(map[rune]int)

		for _, i := range hand.cards {
			if _, ok := count[i]; ok {
				count[i] += 1
			} else {
				count[i] = 1
			}
		}

		if count['J'] > 0 {
			highV := 0
			highKey := 'J'
			for y := range count {
				if y != 'J' {
					if count[y] > highV {
						highKey = y
						highV = count[y]
					} else if count[y] == highV && valueDict[y] > valueDict[highKey] {
						highKey = y
					}
				}
			}
			if highKey != 'J' {
				count[highKey] += count['J']
				delete(count, 'J')
			}
		}

		value := 1
		for _, i := range count {
			value *= i
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

	convertedMatches := make([][]int, 0)

	for _, x := range matches {
		temp := make([][]int, 0)
		for _, i := range x {
			y := strings.ReplaceAll(i.cards, "A", "E")
			y = strings.ReplaceAll(y, "T", "A")
			y = strings.ReplaceAll(y, "J", "1")
			y = strings.ReplaceAll(y, "Q", "C")
			y = strings.ReplaceAll(y, "K", "D")
			val, _ := strconv.ParseInt(y, 16, 0)
			temp = append(temp, []int{int(val), i.bid})
		}
		sort.Slice(temp, func(i, j int) bool {
			return temp[i][0] > temp[j][0]
		})
		for _, i := range temp {
			convertedMatches = append(convertedMatches, i)
		}
	}

	total := 0
	for x := 0; x < len(convertedMatches); x++ {
		total += convertedMatches[x][1] * (len(convertedMatches) - x)
	}

	fmt.Println(total)
}
