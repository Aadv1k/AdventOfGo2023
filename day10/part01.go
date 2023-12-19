package day10

import (
	"fmt"
	"log"
	"strings"
	"time"
)

var directions = [][2]int{
	{0, 1},
	{-1, 0}, {1, 0},
	{0, -1},
}

func printField(field Field, current Vec2) {
	fmt.Print("\033[H\033[2J")

	for i := range field {
		for j := range field[i] {
			p := field[i][j]

			if i == current.y && j == current.x {
				// Highlight the current pipe with ANSI escape codes for red text
				fmt.Printf("\x1b[31m[%c]\x1b[0m", p.ptype)
			} else {
				fmt.Printf(" %c ", p.ptype)
			}
		}
		fmt.Println()
	}

	time.Sleep(750 * time.Millisecond)
}

type Pipe struct {
	visited bool
	ptype   byte
}

type Field [][]Pipe

type Vec2 struct {
	x int
	y int
}

func CompareVec2GT(v1, v2 Vec2) bool {
	return v1.x > v2.x || (v1.x == v2.x && v1.y > v2.y)
}

func traverse(field Field, start Vec2) Vec2 {
	curPipe := field[start.y][start.x]

	if curPipe.ptype == 'S' || curPipe.ptype == '.' {
		return start
	}

	currentX, currentY := start.x, start.y

	switch curPipe.ptype {
	case 'J':
		if !field[start.y][start.x-1].visited {
			// assume we will go to top
			currentY--
		} else {
			// we will go to left
			currentX--
		}
	case '-':
		if field[start.y][start.x-1].visited {
			currentX++
		} else {
			currentX--
		}
	case 'F':
		if !field[start.y][start.x+1].visited {
			// assume we will go down
			currentY++
		} else {
			// assume we will go right
			currentX++
		}
	case 'L':
		if field[start.y-1][start.x].visited {
			// assume we will go to right
			currentX++
		} else {
			// we will go to top
			currentY--
		}
	case '7':
		if field[start.y-1][start.x].visited {
			// we'll go left
			currentX--
		} else {
			// we'll go down
			currentY--
		}
	case '|':
		if !field[start.y-1][start.x].visited {
			currentY++
		} else {
			currentY--
		}
	default:
		log.Panicf("Expected a valid pipe, got: %c", curPipe.ptype)
	}

	current := Vec2{x: currentX, y: currentY}
	if CompareVec2GT(start, current) {
		traverse(field, current)
	}

	printField(field, current)
	return current
}

func Part01(input string) {
	lines := strings.Split(input, "\n")

	var field Field
	var animalIndex [2]int

	for i, line := range lines {
		var block []Pipe

		for j := range line {

			if line[j] == 'S' {
				animalIndex[0] = j
				animalIndex[1] = i
			}

			block = append(block, Pipe{
				ptype:   line[j], // fixed the index to use j instead of i
				visited: false,
			})
		}

		field = append(field, block)
	}

	// gofmt:ignore
	var directions = [][2]int{
		{0, 1},
		{-1, 0}, {1, 0},
		{0, -1},
	}
	// gofmt:ignore

	var final Vec2

	for i := animalIndex[1]; i < len(field)-1; i++ {
		for j := animalIndex[0]; j < len(field[0])-1; j++ {

			for _, direction := range directions {
				x, y := direction[0]+j, direction[1]+i

				if y < 0 || y >= len(field) || x < 0 || y >= len(field[0]) {
					continue
				}

				adjacent := field[y][x]

				if adjacent.ptype == '.' {
					continue
				}

				final = traverse(field, Vec2{
					y: y,
					x: x,
				})
			}
		}
	}

	log.Printf("final: %d %d", final.x, final.y)
}
