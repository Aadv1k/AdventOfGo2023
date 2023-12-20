package day11

import (
	"fmt"
	"strings"

	"github.com/aadv1k/AdventOfGo2023/utils"
)

type GalaxyImage [][]byte

func printGalaxy(g GalaxyImage) {
	fmt.Println()
	for _, row := range g {
		for _, pixel := range row {
			fmt.Printf("%c ", pixel)
		}
		fmt.Println()
	}
}

func DoExpansion(g *GalaxyImage) {
	var emptyColumns []int

	for i := 0; i < len((*g)[0]); i++ {
		emptyColumns = append(emptyColumns, i)
	}

	for i := range *g {
		shouldExpand := true
		currentRow := 0

		for j := range (*g)[i] {
			currentRow = j

			if *g[i][j] != '.' {
				shouldExpand = false
			}
		}

		if shouldExpand {
			// just trust this code :)
			i, _ = utils.Find(emptyColumns, currentRow)
			if i >= 0 {
				emptyColumns = append(emptyColumns[:i+1], emptyColumns[i:]...)
			}

			// NOTE: this is hella sus
			newRow := make([]byte, len((*g)[0]))
			*g = append((*g)[:i+1], append(GalaxyImage{newRow}, (*g)[i+1:]...)...)
		}
	}

	// NOTE: ???
	for _, col := range emptyColumns {
		for _, row := range *g {
			row = append(row[:col+1], row[col+1:]...)
		}
	}

}

func Part01(input string) {
	lines := strings.Split(input, "\n")

	var galaxyImg GalaxyImage

	for i := range lines {
		var row []byte
		shouldExpand := true

		for j := range lines[i] {
			if lines[i][j] != '.' {
				shouldExpand = false
			}

			row = append(row, lines[i][j])
		}

		if shouldExpand {
			galaxyImg = append(galaxyImg, row, row)
			continue
		}

		galaxyImg = append(galaxyImg, row)
	}

	DoExpansion(&galaxyImg)

	printGalaxy(galaxyImg)
}
