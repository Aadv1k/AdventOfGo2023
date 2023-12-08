package day05

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aadv1k/AdventOfGo2023/utils"
)

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

	var scores []int

	_, seeds, _ := strings.Cut(mapContent[0], ": ")
	seedsAsWords := strings.Fields(seeds)

	// NOTE: This is an absolutelly filthy, disgusting, nasty and grisly way to handle this logic especially with u8+ numbers (pretty much our input)
	// 	     any other solution? Get faster processor lol. On a serious note, we can do a couple of things to speed this up.
	for i := 0; i < len(seedsAsWords)-1; i += 2 {
		baseSeed, iterAmount := utils.ParseInt(seedsAsWords[i]), utils.ParseInt(seedsAsWords[i+1])

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
