package day02

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func getMinimumPowerForLine(str string) int {
	minAgainstRegex := func(re *regexp.Regexp) int {
		ret := 0
		for _, x := range re.FindAllString(str, -1) {
			lhs, _, _ := strings.Cut(x, " ")
			count, _ := strconv.Atoi(lhs)
			if ret <= count {
				ret = count
			}
		}
		return ret
	}
	return minAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) blue")) *
		minAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) red")) *
		minAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) green"))
}

func Part02(input string) {
	sumOfPowers := 0
	for _, line := range strings.Split(input, "\r\n") {
		sumOfPowers += getMinimumPowerForLine(line)
	}

	fmt.Printf("The sum of minimum powers for each game is %d\n", sumOfPowers)
}
