package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isIntChar(c byte) bool {
	return c <= '9' && c >= '0'
}

func findNextInt(start int, column int, lines []string) int {
	for i := start; i < len(lines[column]); i++ {
		if isIntChar(lines[column][i]) {
			return i
		}
	}
	return len(lines[column])
}

func hasNeighbours(start int, end int, column int, lines []string) bool {
	directions := [][2]int{
		{1, -1}, {1, 0}, {1, 1},
		{0, -1}, {0, 1},
		{-1, -1}, {-1, 0}, {-1, 1},
	}

	for i := start; i < end; i++ {
		for _, dir := range directions {
			row, col := column+dir[0], i+dir[1]
			if row >= 0 && row < len(lines) && col >= 0 && col < len(lines[row]) {
				target := lines[row][col]
				if target == '*' {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	fptr, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fptr.Close()

	scanner := bufio.NewScanner(fptr)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	totalSum := 0

	for i := 1; i < len(lines)-1; i++ {
		numString := ""
		start, end := 0, 0
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] == '.' {
				end = j

				if hasNeighbours(start, end, i, lines) {
					converted, _ := strconv.Atoi(numString)
					totalSum += converted
				}

				//j = findNextInt(end, i, lines)
			}

			fmt.Printf("%c\n", lines[i][j]);
			numString += string(lines[i][j])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
