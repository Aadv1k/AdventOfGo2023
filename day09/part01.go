package day09

import (
	"fmt"
	"strings"
	"github.com/aadv1k/AdventOfGo2023/utils"
)

type series []int

func Part01(input string) {
	lines := strings.Split(input, "\n")
	totalSum := 0

	for _, line := range lines {
		elems := strings.Fields(line)
		var elemsInt series

		for i := range elems {
			elemsInt = append(elemsInt, utils.ParseInt(elems[i]))
		}

		totalSum += RecursivelyFindNextProgression(elemsInt)
	}


	fmt.Printf("The sum of the extrapolated values is %d\n", totalSum)
}

func SumIsZero(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum == 0
}


func RecursivelyFindNextProgression(elems series) int {
	var diffs []int

	for i := 0; i < len(elems)-1; i++ {
		diffs = append(diffs, elems[i+1]-elems[i])
	}

	if SumIsZero(diffs) {
		return elems[len(elems)-1] + diffs[0]
	}

	lastDiff := RecursivelyFindNextProgression(diffs)

	return elems[len(elems)-1] + lastDiff
}
