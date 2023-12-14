package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aadv1k/AdventOfGo2023/day01"
	"github.com/aadv1k/AdventOfGo2023/day02"
	"github.com/aadv1k/AdventOfGo2023/day03"
	"github.com/aadv1k/AdventOfGo2023/day04"
	"github.com/aadv1k/AdventOfGo2023/day05"
	"github.com/aadv1k/AdventOfGo2023/day06"
	"github.com/aadv1k/AdventOfGo2023/utils"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("You need to provide the day!\n")
	}

	switch os.Args[1] {
	case "day01":
		runDay("day01", day01.Part01, day01.Part02)
	case "day02":
		runDay("day02", day02.Part01, day02.Part02)
	case "day03":
		runDay("day03", day03.Part01, day03.Part02)
	case "day04":
		runDay("day04", day04.Part01, day04.Part02)
	case "day05":
		runDay("day05", day05.Part01, day05.Part02)
	case "day06":
		runDay("day06", day06.Part01, nil)
	default:
		log.Fatalf("Unknown day: %s\n", os.Args[1])
	}
}

func runDay(day string, part01, part02 func(string)) {
	input, err := utils.ReadFileIntoString("data/" + day + "/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("===== %s =====\n", day)

	if part01 != nil {
		start := time.Now()
		fmt.Printf("Part01: ")
		part01(input)
		elapsed := time.Since(start)
		fmt.Printf("\t(took %s)\n", elapsed)
	}

	if part02 != nil {
		start := time.Now()
		fmt.Printf("Part02: ")
		part02(input)
		elapsed := time.Since(start)
		fmt.Printf("\t(took %s)\n", elapsed)
	}

	fmt.Println()
}
