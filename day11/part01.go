package day11

import (
	"fmt"
	"strings"
)

func printGalaxy(g [][]byte) {
	fmt.Println()
	for _, row := range g {
		for _, pixel := range row {
			fmt.Printf("%c", pixel)
		}
		fmt.Println()
	}
}

func DoExpansion(g *[][]byte) {
	for i := range *g {
		shouldExpand := true

		for j := range (*g)[i] {
			if (*g)[i][j] != '.' {
				shouldExpand = false
			}
		}

		if shouldExpand {
			var toInsert = []byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'}

			*g = append((*g)[:i+1], (*g)[i:]...)
			(*g)[i] = toInsert
		}
	}

}

func Part01(input string) {
	lines := strings.Split(input, "\n")

	var galaxyImg [][]byte

	for i := range lines {
		var row []byte

		for j := range lines[i] {
			row = append(row, lines[i][j])
		}

		galaxyImg = append(galaxyImg, row)
	}

	DoExpansion(&galaxyImg)

	printGalaxy(galaxyImg)
}
