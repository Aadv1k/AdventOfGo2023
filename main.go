package main

import (
	"github.com/aadv1k/AdventOfGo2023/day01"
	"github.com/aadv1k/AdventOfGo2023/day02"
	"github.com/aadv1k/AdventOfGo2023/day03"
	"github.com/aadv1k/AdventOfGo2023/utils"
)

func main() {
	day01Input, _ := utils.ReadFileIntoString("data/day01/input.txt")

	day01.Part01(day01Input)
	day01.Part02(day01Input)

	day02Input, _ := utils.ReadFileIntoString("data/day02/input.txt")

	day02.Part01(day02Input)
	day02.Part02(day02Input)

	day03Input, _ := utils.ReadFileIntoString("data/day03/input.txt")

	day03.Part01(day03Input)
	day03.Part02(day03Input)

}
