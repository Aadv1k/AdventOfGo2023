package day05

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aadv1k/AdventOfGo2023/utils"
)

type Range struct {
	low  int
	high int
}

func GetDestRangeFromMap(conversionMap []MapItem, target Range) Range {
	for _, cMap := range conversionMap {
		// Case 1: the range "A" starts ahead of "B"
		//       -- (A)
		// ----		 (B)
		// Case 2: the range "A" doesn't fully cover "B"
		// ------     (A)
		// ---------- (B)
		if cMap.src > target.low || target.high > cMap.src+cMap.rangeLength {
			continue
		}

		offset := cMap.dest - cMap.src

		return Range{
			low:  target.low + offset,
			high: target.high + offset,
		}
	}

	return target
}

func Part02(input string) {
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

	_, seeds, _ := strings.Cut(mapContent[0], ": ")
	seedsAsWords := strings.Fields(seeds)

	for i := 0; i < len(seedsAsWords)-1; i += 2 {
		startRange, rangeLength := utils.ParseInt(seedsAsWords[i]), utils.ParseInt(seedsAsWords[i+1])
		endRange := startRange + rangeLength

		baseRange := Range{
			low:  startRange,
			high: endRange,
		}

		for _, conversionMap := range conversionMaps {
			baseRange = GetDestRangeFromMap(conversionMap, baseRange)
		}

		fmt.Printf("the range is now %v\n", baseRange)

	}
}
