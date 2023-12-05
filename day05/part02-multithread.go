package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
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

var (
	conversionMaps [][]MapItem
	mu             sync.Mutex
)

func main() {
	content, err := os.ReadFile("input.txt")
	checkError(err, "Error reading file")

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

	_, seeds, _ := strings.Cut(mapContent[0], ": ")
	seedsAsWords := strings.Fields(seeds)

	firstHalf := seedsAsWords[0 : len(seedsAsWords)/2]
	secondHalf := seedsAsWords[len(seedsAsWords)/2:]

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		result := ProcessSeeds(firstHalf)
		fmt.Println(result)
	}()

	go func() {
		defer wg.Done()
		result := ProcessSeeds(secondHalf)
		fmt.Println(result)
	}()

	wg.Wait()
}

func ProcessSeeds(seeds []string) int {
	var scores []int

	for i := 0; i < len(seeds)-1; i += 2 {
		baseSeed, iterAmount := parseInt(seeds[i], "Unable to convert seed to an integer"), parseInt(seeds[i+1], "Unable to convert seed to an integer")

		for j := 0; j < iterAmount; j++ {
			seed := baseSeed + j

			for _, conversionMap := range conversionMaps {
				seed = GetDestFromMap(conversionMap, seed)
			}

			mu.Lock()
			scores = append(scores, seed)
			mu.Unlock()

			fmt.Printf("Seed '%d' location: %d\n", baseSeed+j, seed)
		}
	}

	mu.Lock()
	defer mu.Unlock()

	minValue := min(scores...)
	fmt.Printf("The minimum of the above is %d\n", minValue)
	return minValue
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
