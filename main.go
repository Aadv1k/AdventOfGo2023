package main

import (
	"fmt"
	"time"

	"github.com/aadv1k/AdventOfGo2023/day01"
	"github.com/aadv1k/AdventOfGo2023/day02"
	"github.com/aadv1k/AdventOfGo2023/day03"
	"github.com/aadv1k/AdventOfGo2023/day04"
	"github.com/aadv1k/AdventOfGo2023/day05"
	"github.com/aadv1k/AdventOfGo2023/utils"
)

func main() {
	runDay("day01", day01.Part01, day01.Part02)
	runDay("day02", day02.Part01, day02.Part02)
	runDay("day03", day03.Part01, day03.Part02)
	runDay("day04", day04.Part01, day04.Part02)
	runDay("day05", day05.Part01, nil)
}

func runDay(day string, part01, part02 func(string)) {
	input, err := utils.ReadFileIntoString("data/" + day + "/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("===== %s =====\n", day)

	if part01 != nil {
		start := time.Now()
		part01(input)
		elapsed := time.Since(start)
		fmt.Printf("Part01 took %s\n", elapsed)
	}

	if part02 != nil {
		start := time.Now()
		part02(input)
		elapsed := time.Since(start)
		fmt.Printf("Part02 took %s\n", elapsed)
	}

	fmt.Println()
}
