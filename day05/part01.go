package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MapItem struct {
	source      int
	dest        int
	rangeLength int
}

func ConvertStringToMapItem(s string) MapItem {
	parts := strings.Fields(s)

	dest, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Error converting source to int: %v", err)
	}

	source, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Error converting dest to int: %v", err)
	}

	length, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatalf("Error converting length to int: %v", err)
	}

	return MapItem{
		source:      source,
		dest:        dest,
		rangeLength: length,
	}
}

func main() {
	content, err := os.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	var conversionMaps [][]MapItem

	re := regexp.MustCompile(`\r?\n\r?\n`)
	mapContent := re.Split(string(content), -1)

	for _, item := range mapContent[1:] { // Ignore the seed string.
		mapStrs := strings.Split(item, "\n")

		var conversionMap []MapItem
		for _, mapStr := range mapStrs[1:] { // Ignore the title; eg "seed-to-soil map".
			if len(mapStr) == 0 {
				continue
			}
			conversionMap = append(conversionMap, ConvertStringToMapItem(mapStr))
		}
		conversionMaps = append(conversionMaps, conversionMap)
	}

	_, seeds, _ := strings.Cut(mapContent[0], ": ")
	for _, seedStr := range strings.Split(seeds, " ")[0:1] {
		seed, err := strconv.Atoi(seedStr)

		if err != nil {
			log.Fatalf("Unable to convert '%s' to an integer\n", seedStr)
		}

		fmt.Printf("Initial Seed value: %d\n", seed)
		for _, conversionMap := range conversionMaps {
			newSeed := GetDestFromMap(conversionMap, seed)
			seed = newSeed
			fmt.Printf("-> value: %d\n", seed)
		}

		fmt.Printf("Final seed value: %d\n", seed)
	}
}

func GetDestFromMap(cMap []MapItem, seed int) int {
	for _, mapItem := range cMap {
		if mapItem.source <= seed {
			maybeSeed := (seed - mapItem.source) + mapItem.dest
			return maybeSeed
		}
	}
	return seed
}
