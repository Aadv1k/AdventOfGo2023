package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func isIntChar(c byte) bool {
	return c <= '9' && c >= '0'
}

// This represents a number location with one adjacent symbol.
type HalfGear struct {
	y      int
	startX int
	endX   int
	value  int
}

func PrintHalfGear(gear HalfGear) {
	fmt.Printf(" Gear at Y=%d, StartX=%d, EndX=%d, Value=%d\n", gear.y, gear.startX, gear.endX, gear.value)
}

func main() {
	fptr, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer fptr.Close()

	scanner := bufio.NewScanner(fptr)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var directions = [][2]int{
		{-1, 1}, {0, 1}, {1, 1},
		{-1, 0}, {1, 0},
		{-1, -1}, {0, -1}, {1, -1},
	}

	re := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,<>\/?]`)

	getAdjacentSymbolLoc := func(colStart, colEnd, row int) ([2]int, bool) {
		for i := colStart; i < colEnd; i++ {
			for _, dir := range directions {
				newCol := i + dir[0]
				newRow := row + dir[1]

				if newCol >= 0 && newCol < len(lines[0]) && newRow >= 0 && newRow < len(lines) {
					if re.MatchString(string(lines[newRow][newCol])) {
						return [2]int{newRow, newCol}, true
					}
				}
			}
		}
		return [2]int{}, false
	}

	storedGears := make(map[string]int)

	gearRatioSum := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if isIntChar(lines[i][j]) {
				endX := j

				for endX < len(lines[0]) && lines[i][endX] != '.' && !re.MatchString(string(lines[i][endX])) {
					endX++
				}

				loc, found := getAdjacentSymbolLoc(j, endX, i)
				if !found {
					continue
				}

				key := fmt.Sprintf("%d-%d", loc[0], loc[1])
				value, exists := storedGears[key]

				numeral, _ := strconv.Atoi(lines[i][j:endX])

				j = endX

				if !exists {
					storedGears[key] = numeral
					continue
				}

				gearRatioSum += value * numeral

			}
		}
	}

	fmt.Printf("The total gear-ratio sum is %d\n", gearRatioSum);

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
