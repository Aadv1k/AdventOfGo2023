package day10

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/aadv1k/AdventOfGo2023/utils"
)

const RecursionLimit = 200_000

var directions = [][2]int{
	{0, 1},
	{-1, 0}, {1, 0},
	{0, -1},
}

var stepsTaken = 0

func printField(field Field, current Vec2) {
	clearScreen()

	diameter := 12

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
				fmt.Printf("\x1b[31m[%c]\x1b[0m", '🐭')
			} else {
				fmt.Printf(" %c ", p.ptype)
			}
		}

		// Print a newline after each line
		fmt.Println()
	}

	time.Sleep(1 * time.Millisecond)
}

func clearScreen() {
	cmd := exec.Command("clear")
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

func CompareVec2EQ(v1, v2 Vec2) bool {
	return v1.x == v2.x && v1.y == v2.y
}

func isWithinBounds(field Field, x, y int) bool {
	return y >= 0 && y < len(field) && x >= 0 && x < len(field[0])
}

func distanceInLoop(start, point, loopSize int) int {
	clockwiseDistance := (point - start + loopSize) % loopSize
	counterclockwiseDistance := (start - point + loopSize) % loopSize

	return int(math.Min(float64(clockwiseDistance), float64(counterclockwiseDistance)))
}

func farthestInLoop(loop []Vec2, origin Vec2) int {
	if len(loop) == 0 {
		return 0
	}

	maxSteps := distanceInLoop(origin.x, loop[0].x, len(loop))

	for _, point := range loop[1:] {
		steps := distanceInLoop(origin.x, point.x, len(loop))
		if steps > maxSteps {
			maxSteps = steps
		}
	}

	return maxSteps
}

func traverse(field Field, start Vec2) []Vec2 {
	var visitedPoints []Vec2

	currentPipe := field[start.y][start.x]
	currentX, currentY := start.x, start.y

	if currentPipe.visited {
		return visitedPoints
	}

	stepsTaken++

	field[currentY][currentX].visited = true

	switch currentPipe.ptype {
	case 'J':
		if currentX-1 >= 0 && !field[currentY][currentX-1].visited {
			currentX--
		} else {
			currentY--
		}
	case '7':
		if currentY+1 < len(field) && !field[currentY][currentX-1].visited {
			currentX--
			break
		}

		if !field[currentY+1][currentX].visited {
			currentY++
		}

	case 'L':
		if currentX+1 < len(field[0]) && !field[currentY][currentX+1].visited {
			currentX++
		} else {
			currentY--
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

	if stepsTaken >= RecursionLimit {
		log.Panicf("Execeeded recursion limit of %d, either increase or something went wrong", RecursionLimit)
	}

	currentPipe.visited = true

	var current = Vec2{
		x: currentX,
		y: currentY,
	}

	if !CompareVec2EQ(current, start) {
		visitedPoints = append(visitedPoints, traverse(field, current)...)
	}
	visitedPoints = append(visitedPoints, current) // adding it before the 'if statement' flips it

	return visitedPoints
}

func Part01(input string) {
	lines := strings.Split(input, "\n")

	var field Field
	var animalIndex Vec2

	for i, line := range lines {
		var block []Pipe

		for j := range line {
			visited := false
			if line[j] == 'S' {
				animalIndex.x = j
				animalIndex.y = i
				visited = true
			}

			block = append(block, Pipe{
				ptype:   line[j],
				visited: visited,
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

	var distances []int

	for _, direction := range directions {
		x, y := direction[0]+animalIndex.x, direction[1]+animalIndex.y

		if !isWithinBounds(field, x, y) {
			continue
		}

		adjacent := field[y][x]

		if adjacent.ptype == '.' {
			continue
		}

		visitedPoints := traverse(field, Vec2{
			y: y,
			x: x,
		})

		for _, point := range visitedPoints {
			field[point.y][point.x].visited = false
		}

		idx := farthestInLoop(visitedPoints, animalIndex)

		distances = append(distances, idx)
		stepsTaken = 0
	}

	log.Printf("The maximum distance is %d", utils.Sum(distances))
}
