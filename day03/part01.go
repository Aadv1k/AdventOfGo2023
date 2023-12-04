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

	hasAdjacentSymbol := func(colStart int, colEnd int, row int) bool { 
		for i := colStart; i < colEnd; i++ {
			for _, dir := range directions { 
				newCol := i + dir[0]
				newRow := row + dir[1]

				if newCol >= 0 && newCol < len(lines[0]) && newRow >= 0 && newRow < len(lines) {
					if re.MatchString(string(lines[newRow][newCol])) {
						return true
					}
				}
			}
		}
		return false
	}

	totalSum := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if isIntChar(lines[i][j]) {
				endX := j

				for endX < len(lines[0]) && lines[i][endX] != '.' && !re.MatchString(string(lines[i][endX])) {
					endX++
				}

				if hasAdjacentSymbol(j, endX, i) {
					numeral, _ := strconv.Atoi(lines[i][j:endX]);
					fmt.Printf("Good part: %d\n", numeral);
					totalSum += numeral
				}

				j = endX 
			}
		}
	}

	fmt.Printf("The total sum of the good engine parts is %d", totalSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
