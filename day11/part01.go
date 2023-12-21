package day11

import (
	"fmt"
	"strings"

	"github.com/aadv1k/AdventOfGo2023/utils"
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
	var ret [][]byte

	emptyColumns := make([]int, len(*g))
	for i := range *g {
		emptyColumns[i] = i
	}

	for i, elem := range *g {
		shouldExpand := true
		for j := range elem {
			if (*g)[i][j] != '.' {
				shouldExpand = false

				found, _ := utils.Find[int](emptyColumns, j)
				if found != -1 {
					emptyColumns = append(emptyColumns[:found], emptyColumns[found+1:]...)
				}
			}
		}

		if shouldExpand {
			var toInsert = []byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'}
			ret = append(ret, toInsert, toInsert)
		}

		ret = append(ret, (*g)[i])
	}

	for _, col := range emptyColumns {
		for i := range ret {
			var newRow []byte
			newRow = append(newRow, ret[i][:col]...)
			newRow = append(newRow, []byte{'.', '.'}...)
			newRow = append(newRow, ret[i][col:]...)

			ret[i] = newRow
		}
	}

	*g = ret

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
