package day05

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/aadv1k/AdventOfGo2023/utils"
)

type MapItem struct {
	src         int
	dest        int
	rangeLength int
}

func ConvertStringToMapItem(s string) MapItem {
	parts := strings.Fields(s)

	dest, _ := strconv.Atoi(parts[0])
	source, _ := strconv.Atoi(parts[1])
	length, _ := strconv.Atoi(parts[2])

	return MapItem{
		src:         source,
		dest:        dest,
		rangeLength: length,
	}
}

func Part01(input string) {
	var conversionMaps [][]MapItem

	re := regexp.MustCompile(`\r?\n\r?\n`)
	mapContent := re.Split(input, -1)

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
	for _, seedStr := range strings.Fields(seeds) {
		seed := utils.ParseInt(seedStr)

		for _, conversionMap := range conversionMaps {
			seed = GetDestFromMap(conversionMap, seed)
		}

		scores = append(scores, seed)

		//fmt.Printf("Seed '%s' location: %d\n", seedStr, seed)
	}

	fmt.Printf("The minimum of the above is %d\n", min(scores...))
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
