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
	src         int
	dest        int
	rangeLength int
}

func ConvertStringToMapItem(s string) MapItem {
	parts := strings.Fields(s)

	dest, err := strconv.Atoi(parts[0])
	checkError(err, "Error converting dest to int")

	source, err := strconv.Atoi(parts[1])
	checkError(err, "Error converting src to int")

	length, err := strconv.Atoi(parts[2])
	checkError(err, "Error converting length to int")

	return MapItem{
		src:         source,
		dest:        dest,
		rangeLength: length,
	}
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func main() {
	content, err := os.ReadFile("sample.txt")
	checkError(err, "Error reading file")

	var conversionMaps [][]MapItem

	re := regexp.MustCompile(`\r?\n\r?\n`)
	mapContent := re.Split(string(content), -1)

	for _, item := range mapContent[1:] {
		mapStrs := strings.FieldsFunc(item, func(r rune) bool { return r == '\n' })

		var conversionMap []MapItem
		for _, mapStr := range mapStrs[1:] {
			if len(mapStr) == 0 {
				continue
			}
			conversionMap = append(conversionMap, ConvertStringToMapItem(mapStr))
		}
		conversionMaps = append(conversionMaps, conversionMap)
	}

	var scores []int

	_, seeds, _ := strings.Cut(mapContent[0], ": ")
	seedsAsWords := strings.Fields(seeds)

	// NOTE: This is an absolutelly filthy, disgusting, nasty and grisly way to handle this logic especially with u8+ numbers (pretty much our input)
	// 	     any other solution? Get faster processor lol. On a serious note, we can do a couple of things to speed this up.
	for i := 0; i < len(seedsAsWords)-1; i += 2 {
		baseSeed, iterAmount := parseInt(seedsAsWords[i], "Unable to convert seed to an integer"), parseInt(seedsAsWords[i+1], "Unable to convert seed to an integer")

		for j := 0; j < iterAmount; j++ {
			seed := baseSeed + j

			for _, conversionMap := range conversionMaps {
				seed = GetDestFromMap(conversionMap, seed)
			}

			scores = append(scores, seed)

			fmt.Printf("Seed '%d' location: %d\n", baseSeed+j, seed)
		}
	}

	fmt.Printf("The minimum of the above is %d\n", min(scores...))
}

func parseInt(s, errMsg string) int {
	i, err := strconv.Atoi(s)
	checkError(err, errMsg)
	return i
}

func GetDestFromMap(cMap []MapItem, seed int) int {
	for _, mapItem := range cMap {
		if seed >= mapItem.src && seed <= mapItem.src+mapItem.rangeLength {
			return mapItem.dest + (seed - mapItem.src)
		}
	}

	return seed
}

func min(values ...int) int {
	if len(values) == 0 {
		log.Fatal("min: empty slice")
	}

	minValue := values[0]
	for _, value := range values[1:] {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}
