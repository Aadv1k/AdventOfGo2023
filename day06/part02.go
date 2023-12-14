package day06

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aadv1k/AdventOfGo2023/utils"
)

func Part02(input string) {
	lines := strings.Split(input, "\n")

	re := regexp.MustCompile(`\d+`)

	raceTime := utils.ParseInt(strings.Join(re.FindAllString(lines[0], -1), ""))
	raceRecord := utils.ParseInt(strings.Join(re.FindAllString(lines[1], -1), ""))

	waysToWin := 0

	for pressedFor := 0; pressedFor <= raceTime; pressedFor++ {
		currentTime := (raceTime - pressedFor) * pressedFor
		if currentTime > raceRecord {
			waysToWin++
		}
	}

	fmt.Printf("The total ways to win for the single giant race is %d\n", waysToWin)
}
