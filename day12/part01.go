package day12

import (
	"log"
	"regexp"
	"strings"

	"github.com/aadv1k/AdventOfGo2023/utils"
)

func Part01(input string) {
	hotSprings := make(map[string][]int)

	for _, line := range strings.Split(input, "\n") {
		elems := strings.Fields(line)

		var springCount []int
		for _, letter := range strings.Split(elems[1], ",") {
			springCount = append(springCount, utils.ParseInt(letter))
		}

		hotSprings[elems[0]] = springCount
	}

	waysToArrange := 0
	for spring, inst := range hotSprings {
		groups := getContiguousSprings(spring)

		if len(inst) == len(groups) {
			waysToArrange++
			continue
		}

		if len(inst) > len(groups) {
			// first remove the obvious matches

			for i := range inst {
				if WasConsumed(&groups, inst[i]) {
					waysToArrange++
				}
			}

		}

		log.Printf("%s %v -> %v", spring, groups, inst)
	}
}

func WasConsumed(a *[]string, x int) bool {
	for i := range *a {
		current := (*a)[i]
		if len(current) == x {
			return true
		}
		currentRunes := []rune(current)

		for j := 0; j < x; j++ {
			if i+j >= len(currentRunes) {
				break // inner loop
			}
			currentRunes[i+j] = '#'
		}
		(*a)[i] = string(currentRunes)
	}

	return false
}

func getContiguousSprings(spring string) []string {
	re := regexp.MustCompile("[#|?]+")
	return re.FindAllString(spring, -1)
}

/* ?###???????? 3,2,1

 */
