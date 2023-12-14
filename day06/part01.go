package day06

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aadv1k/AdventOfGo2023/utils"
)

func Part01(input string) {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`\d+`)

	totalTimes := re.FindAllString(lines[0], -1)
	bestDistances := re.FindAllString(lines[1], -1)

	var totalWaysToWin []int

	// here 1 ms = 1 mm, so 1:1 ratio
	for i, time := range totalTimes {
		waysToWin := 0

		for pressedFor := 0; pressedFor <= utils.ParseInt(time); pressedFor++ {
			currentTime := (utils.ParseInt(time) - pressedFor) * pressedFor
			if currentTime > utils.ParseInt(bestDistances[i]) {
				waysToWin++
			}
		}

		totalWaysToWin = append(totalWaysToWin, waysToWin)
	}

	product := 1
	for _, elem := range totalWaysToWin {
		product *= elem
	}

	fmt.Printf("The product of total ways to win is %d\n", product)

}
