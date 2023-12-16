package day08

import (
	"log"
	"regexp"
	"strings"
)

type Instruction struct {
	left  string
	right string
}

const ElemToMatch = "ZZZ"

func Part01(input string) {
	re := regexp.MustCompile("[A-Z]{3}")

	lines := strings.Split(input, "\n")

	desertMap := make(map[string]Instruction)
	var first string

	// Ignore the first two lines since they are the instruction set and a blank line
	for i, line := range lines[2:] {

		if len(line) == 0 {
			continue
		}

		matches := re.FindAllString(line, -1)
		desertMap[matches[0]] = Instruction{
			left:  matches[1],
			right: matches[2],
		}

		if i == 0 {
			first = matches[0]
		}
	}

	_ = first
	current := "AAA"
	steps := 0

	for current != ElemToMatch {
		for i := range strings.TrimSpace(lines[0]) {
			direction := lines[0][i]
			if direction == 'R' {
				current = desertMap[current].right
			} else if direction == 'L' {
				current = desertMap[current].left
			}
			steps++

			if current == ElemToMatch {
				break
			}
		}
	}


	log.Printf("It took %d steps to reach %s", steps, ElemToMatch)
}
