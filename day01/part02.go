package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

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

func isDigit(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func getFirstDigitOrSpelling(targetString string) int {
	var buf string

	for i := 0; i < len(targetString); i++ {
		buf += string(targetString[i])

		foundDigit := digitRe.FindString(buf)

		if len(foundDigit) != 0 {
			return digits[foundDigit]
		}

		if isDigit(targetString[i]) {
			val, _ := strconv.Atoi(string(targetString[i]))
			return val
		}
	}

	fmt.Printf("NOT FOUND")
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

		if isDigit(targetString[i]) {
			val, _ := strconv.Atoi(string(targetString[i]))
			return val
		}
	}

	fmt.Printf("NOT FOUND: %s\n", targetString)
	return 0
}

func main() {
	fptr, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fptr.Close()

	scanner := bufio.NewScanner(fptr)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		firstDigit := getFirstDigitOrSpelling(line)
		lastDigit := lastDigitOrSpelling(line)

		t := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
		calibrationValue, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal(err)
		}

		sum += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part02: The sum of all calibration, including words values is %d\n", sum)
}
