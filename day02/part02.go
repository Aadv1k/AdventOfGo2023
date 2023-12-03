package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"os"
	"regexp"
	"strconv"
)

func getMinimumPowerForLine(str string) int {
	minAgainstRegex := func (re *regexp.Regexp) int {
		ret := 0;
		for _, x := range re.FindAllString(str, -1) {
			lhs, _, _ := strings.Cut(x, " ");
			count, _ := strconv.Atoi(lhs);
			if ret <= count {
				ret = count;
			}
		}
		return ret;
	}
	return minAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) blue")) *
	minAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) red")) *
	minAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) green"))
}

func main() {
	fptr, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer fptr.Close()

	scanner := bufio.NewScanner(fptr)


	sumOfPowers := 0;
	for scanner.Scan() {
		line := scanner.Text();
		sumOfPowers += getMinimumPowerForLine(line);
	}

	fmt.Printf("The sum of minimum powers for each game is %d\n", sumOfPowers);

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
