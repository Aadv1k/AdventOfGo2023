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
	if len(os.Args) <= 2 {
		log.Fatalf("You need to provide both the day and the path to input!\n\t%s <day> <path/to/sample.txt>\n", os.Args[0])
	}

	day := os.Args[1]
	filePath := os.Args[2]

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File not found: %s\n", filePath)
	}

	input, err := utils.ReadFileIntoString(filePath)
	if err != nil {
		log.Fatalf("Error reading input file: %v\n", err)
	}

	switch day {
	case "day01":
		runDay(day, day01.Part01, day01.Part02, input)
	case "day02":
		runDay(day, day02.Part01, day02.Part02, input)
	case "day03":
		runDay(day, day03.Part01, day03.Part02, input)
	case "day04":
		runDay(day, day04.Part01, day04.Part02, input)
	case "day05":
		runDay(day, day05.Part01, day05.Part02, input)
	case "day06":
		runDay(day, day06.Part01, day06.Part02, input)
	default:
		log.Fatalf("Unknown day: %s\n", day)
	}
}

func runDay(day string, part01, part02 func(string), input string) {
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
