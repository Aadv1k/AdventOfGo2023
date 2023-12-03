package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const BLUE_COUNT = 14
const RED_COUNT = 12
const GREEN_COUNT = 13

func isGameLineGood(str string) bool {
	compareAgainstRegex := func (re *regexp.Regexp, target int) bool {
		for _, x := range re.FindAllString(str, -1) {
			cubeCount, _ := strconv.Atoi(x[0:2])
			if target < cubeCount {
				return false
			}
		}
		return true;
	}

	return compareAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) blue"), BLUE_COUNT) &&
	compareAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) red"), RED_COUNT) &&
	compareAgainstRegex(regexp.MustCompile("([0-9]|[0-1][0-9]|20) green"), GREEN_COUNT)
}

func main() {
	fptr, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer fptr.Close()

	scanner := bufio.NewScanner(fptr)

	sumOfGID, goodGames, lineCount := 0, 0, 1;
	for scanner.Scan() {
		line := scanner.Text();
		if isGameLineGood(line) {
			sumOfGID += lineCount + 1;
			goodGames ++;
		}
		lineCount ++;
	}


	fmt.Printf("There are %d possible games and their sum-score is %d\n", goodGames, sumOfGID);

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
