package day04

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func Part01(input string) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		sum += int(getPointsForLine(line))
	}

	fmt.Printf("Total points of the cards %d\n", sum)
}

func getPointsForLine(line string) int {
	_, cardPoints, _ := strings.Cut(line, ": ")
	givenPoints, winningPoints, _ := strings.Cut(cardPoints, " | ")

	powCount := 0

	re := regexp.MustCompile("([0-9][0-9]|[0-9])")

	for _, givenPoint := range re.FindAllString(givenPoints, -1) {
		for _, winningPoint := range re.FindAllString(winningPoints, -1) {
			if strings.TrimSpace(givenPoint) == strings.TrimSpace(winningPoint) {
				powCount++
			}
		}
	}

	if powCount <= 2 {
		return powCount
	}
	return int(math.Pow(2, float64(powCount-1)))
}
