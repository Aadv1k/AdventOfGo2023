package main

import (
	"github.com/aadv1k/AdventOfGo2023/day01"
	"github.com/aadv1k/AdventOfGo2023/utils"
)

func main() {
	sample, _ := utils.ReadFileIntoString("data/day01/sample.txt")

	day01.Part01(sample)
	day01.Part02(sample)
}
