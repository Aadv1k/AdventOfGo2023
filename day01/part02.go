package day01

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/aadv1k/AdventOfGo2023/utils"
)

func Part02(input string) {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		firstDigit := getFirstDigitOrSpelling(line)
		lastDigit := lastDigitOrSpelling(line)

		t := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
		calibrationValue, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal(err)
		}

		sum += calibrationValue
	}

	fmt.Printf("The sum of all calibration, including words values is %d\n", sum)
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var digitPattern = "one|two|three|four|five|six|seven|eight|nine"
var digitRe = regexp.MustCompile(digitPattern)

func getFirstDigitOrSpelling(targetString string) int {
	var buf string

	for i := 0; i < len(targetString); i++ {
		buf += string(targetString[i])

		foundDigit := digitRe.FindString(buf)

		if len(foundDigit) != 0 {
			return digits[foundDigit]
		}

		if utils.IsDigit(targetString[i]) {
			val, _ := strconv.Atoi(string(targetString[i]))
			return val
		}
	}
	return 0
}

func reverseString(str string) string {
	var result string
	for _, c := range str {
		result = string(c) + result
	}
	return result
}

func lastDigitOrSpelling(targetString string) int {
	var buf string

	for i := len(targetString) - 1; i >= 0; i-- {
		buf += string(targetString[i])

		foundDigit := digitRe.FindString(reverseString(buf))

		if len(foundDigit) != 0 {
			return digits[foundDigit]
		}

		if utils.IsDigit(targetString[i]) {
			val, _ := strconv.Atoi(string(targetString[i]))
			return val
		}
	}

	return 0
}
