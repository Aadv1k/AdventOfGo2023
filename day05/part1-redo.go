package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

type LookupTable map[int]int

func ConvertMapToLookupTable(xToYMap string) LookupTable {
	tbl := make(LookupTable)

	for _, mapStr := range strings.Split(xToYMap, "\n")[1:] {
		if len(mapStr) == 0 {
			continue
		}
		nums := strings.Fields(mapStr)

		destinationStartStr, sourceStartStr, rangeLengthStr := nums[0], nums[1], nums[2]

		destinationStart, err := strconv.Atoi(destinationStartStr)
		checkError(err, "Error converting destination start to integer")

		sourceStart, err := strconv.Atoi(sourceStartStr)
		checkError(err, "Error converting source start to integer")

		rangeLength, err := strconv.Atoi(rangeLengthStr)
		checkError(err, "Error converting range length to integer")

		for i := 0; i < rangeLength; i++ {
			tbl[sourceStart] = destinationStart
			sourceStart++
			destinationStart++
		}

	}

	return tbl
}

func main() {
	content, err := os.ReadFile("input.txt")
	checkError(err, "Error reading file")

	var lookupTables []LookupTable

	re := regexp.MustCompile(`\r?\n\r?\n`)
	mapContent := re.Split(string(content), -1)

	for _, str := range mapContent[1:] { // ignore the "seeds" line
		lookupTables = append(lookupTables, ConvertMapToLookupTable(str))
	}

	_, seeds, _ := strings.Cut(mapContent[0], ": ")
	for _, seedStr := range strings.Fields(seeds) {
		seed, _ := strconv.Atoi(seedStr)

		for _, table := range lookupTables {
			found, exists := table[seed]
			if !exists {
				continue
			}
			seed = found
		}
		fmt.Printf("Seed '%s' location: %d\n", seedStr, seed)
	}
}
