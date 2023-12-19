package day10

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var directions = [][2]int{
	{0, 1},
	{-1, 0}, {1, 0},
	{0, -1},
}

var stepsTaken = 0

func printField(field Field, current Vec2) {
	clearScreen()

	diameter := 4

	startX := current.x - diameter
	if startX < 0 {
		startX = 0
	}

	startY := current.y - diameter
	if startY < 0 {
		startY = 0
	}

	endX := current.x + diameter + 1
	if endX > len(field[0]) {
		endX = len(field[0])
	}

	endY := current.y + diameter + 1
	if endY > len(field) {
		endY = len(field)
	}

	for i := startY; i < endY; i++ {
		for j := startX; j < endX; j++ {
			p := field[i][j]

			if i == current.y && j == current.x {
				fmt.Printf("\x1b[31m[%c]\x1b[0m", p.ptype)
			} else {
				fmt.Printf(" %c ", p.ptype)
			}
		}
		fmt.Println()
	}

	time.Sleep(250 * time.Millisecond)
}

func clearScreen() {
	cmd := exec.Command("clear") // for Unix-like systems
	cmd.Stdout = os.Stdout
	cmd.Run()
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

func isWithinBounds(field Field, x, y int) bool {
	return y >= 0 && y < len(field) && x >= 0 && x < len(field[0])
}

func traverse(field Field, start Vec2) Vec2 {
	currentPipe := field[start.y][start.x]
	currentX, currentY := start.x, start.y

	stepsTaken++

	printField(field, start)

	switch currentPipe.ptype {
	case 'J':
		if currentX-1 >= 0 && !field[currentY][currentX-1].visited {
			currentX--
		} else {
			currentY--
		}
	case 'L':
		if currentX+1 < len(field[0]) && !field[currentY][currentX+1].visited {
			currentX++
		} else {
			currentY--
		}
	case '7':
		if currentY-1 >= 0 && !field[currentY-1][currentX].visited {
			currentY--
		} else {
			currentX--
		}
	case 'F':
		if currentY+1 < len(field) && !field[currentY+1][currentX].visited {
			currentY++
		} else {
			currentX++
		}
	case '|':
		if currentY+1 < len(field) && !field[currentY+1][currentX].visited {
			currentY++
		} else {
			currentY--
		}
	case '-':
		if currentX+1 < len(field[0]) && !field[currentY][currentX+1].visited {
			currentX++
		} else {
			currentX--
		}
	case '.':
	case 'S':
		break
	default:
		log.Panic("Something has gone haywire")
	}

	currentPipe.visited = true

	var current = Vec2{
		x: currentX,
		y: currentY,
	}

	if isWithinBounds(field, currentX, currentY) && CompareVec2GT(start, current) {
		current = traverse(field, current)
	}

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
				ptype:   line[j],
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

				if !isWithinBounds(field, x, y) {
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

	log.Printf("final: %d %d %d", final.x, final.y, stepsTaken)
}
