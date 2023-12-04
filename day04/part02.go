package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Card struct {
	winningPoints []string
	givenPoints   []string
	totalCount    int
}

func getPointsForCard(card Card) int {
	powCount := 0

	for _, givenPoint := range card.givenPoints {
		for _, winningPoint := range card.winningPoints {
			if strings.TrimSpace(givenPoint) == strings.TrimSpace(winningPoint) {
				powCount++
			}
		}
	}

	return powCount
}

func lexLineIntoCard(line string) Card {
	_, cardDataStr, _ := strings.Cut(line, ": ")
	cardData := strings.Split(cardDataStr, " | ")

	re := regexp.MustCompile("([0-9][0-9]|[0-9])")

	winningPoints := re.FindAllString(cardData[0], -1)
	givenPoints := re.FindAllString(cardData[1], -1)

	return Card{
		winningPoints: winningPoints,
		givenPoints:   givenPoints,
		totalCount:    1,
	}
}

func main() {
	var cards []Card

	fptr, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fptr.Close()

	fscanner := bufio.NewScanner(fptr)

	for fscanner.Scan() {
		line := fscanner.Text()
		cards = append(cards, lexLineIntoCard(line))
	}

	for i, card := range cards {
		points := getPointsForCard(card);

		fmt.Printf("Point is %d\n", points);
		for j := 1; j <= points; j++ {
			fmt.Printf("-> Modifying card %d at card %d\n", i+j+1, i+1);
			cards[i+j].totalCount += 1 * cards[i].totalCount;
		}
	}

	totalCards := 0
	for _, card := range cards {
		totalCards += card.totalCount
	}

	for i, card := range cards {
		fmt.Printf("Card %d -> count is %d\n", i+1, card.totalCount)
	}

	fmt.Printf("There are %d total cards in the stack!\n", totalCards)

	if err := fscanner.Err(); err != nil {
		log.Fatal(err)
	}
}
