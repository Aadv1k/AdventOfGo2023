package day11

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/aadv1k/AdventOfGo2023/utils"
)

func printGalaxy(g [][]byte) {
	fmt.Println()
	for _, row := range g {
		for _, pixel := range row {
			fmt.Printf("%c ", pixel)
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

				found, _ := utils.Find(emptyColumns, j)
				if found != -1 {
					emptyColumns = append(emptyColumns[:found], emptyColumns[found+1:]...)
				}
			}
		}

		if shouldExpand {
			ret = append(ret, []byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'})
			ret = append(ret, []byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'})
		}

		ret = append(ret, (*g)[i])
	}

	for _, col := range emptyColumns {
		for i := range ret {
			if col >= 0 && col <= len(ret[i]) {
				ret[i] = append(ret[i][:col], append([]byte{'.'}, ret[i][col:]...)...)
			}
		}
	}

	*g = ret
}

// x, y
func manhattanDistance(p1, p2 []int) int {
	return int(math.Abs(float64(p2[0]-p1[0])) + math.Abs(float64(p2[1]-p1[1])))
}

// NOTE: this solution is a bit dodgy, this is due to the way we compute the distances, the function
// manhattanDistance can be greately improved
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

	var points [][]int

	for i := range galaxyImg {
		for j := range galaxyImg[i] {
			if galaxyImg[i][j] == '#' {
				points = append(points, []int{i, j})
			}
		}
	}

	var shortestPaths []int
	for i := range points {
		current := points[i]

		for j := range points[i+1:] {
			shortestPaths = append(shortestPaths, manhattanDistance(current, points[j]))
		}
	}

	log.Print(utils.Sum(shortestPaths))
}
