package day02

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	BlueCount  = 14
	RedCount   = 12
	GreenCount = 13
)

func isGameLineGood(str string) bool {
	compareAgainstRegex := func(re *regexp.Regexp, target int) bool {
		for _, x := range re.FindAllString(str, -1) {
			cubeCount, _ := strconv.Atoi(x[0:2])
			if target < cubeCount {
				return false
			}
		}
		return true
	}

	return compareAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) blue"), BlueCount) &&
		compareAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) red"), RedCount) &&
		compareAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) green"), GreenCount)
}

func Part01(input string) {
	sumOfGID, goodGames, lineCount := 0, 0, 1
	for _, line := range strings.Split(input, "\r\n") {
		if isGameLineGood(line) {
			sumOfGID += lineCount + 1
			goodGames++
		}
		lineCount++
	}

	fmt.Printf("There are %d possible games and their sum-score is %d\n", goodGames, sumOfGID)
}
