package day09

import (
	"fmt"
	"strings"
	"github.com/aadv1k/AdventOfGo2023/utils"
)

func Part02(input string) {
	lines := strings.Split(input, "\n")
	totalSum := 0

	for _, line := range lines {
		elems := strings.Fields(line)
		var elemsInt series

		for i := range elems {
			elemsInt = append(elemsInt, utils.ParseInt(elems[i]))
		}

		totalSum += RecursivelyFindNextProgressionReverse(elemsInt)
	}

	fmt.Printf("The sum of the extrapolated values in reverse is %d\n", totalSum)
}

func RecursivelyFindNextProgressionReverse(elems series) int {
	var diffs []int

	for i := 0; i < len(elems)-1; i++ {
		diffs = append(diffs, elems[i+1]-elems[i])
	}

	if SumIsZero(diffs) {
		return elems[0] + diffs[0]
	}

	lastDiff := RecursivelyFindNextProgressionReverse(diffs)

	return elems[0] + lastDiff
}
