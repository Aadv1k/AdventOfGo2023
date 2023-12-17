package day08

import (
	"fmt"
	"regexp"
	"strings"
)

func Part02(input string) {
	// Getting directions, turning L to 0 and R to 1
	directions := strings.ReplaceAll(input[:len(input)-1], "L", "0")
	directions = strings.ReplaceAll(directions, "R", "1")

	lines := strings.Split(input, "\n")[1:]
	// Getting key: [value1, value2] from each line
	mapDict := make(map[string][]string)
	startA := []string{}
	for _, x := range lines {
		key := regexp.MustCompile("[A-Z]+").FindString(x)
		mappings := regexp.MustCompile(`\(.+\)`).FindString(x)
		mappings = regexp.MustCompile("[A-Z]+").FindAllString(mappings, -1)
		if len(mappings) > 1 {
			mapDict[key] = mappings[1:]
		} else {
			mapDict[key] = nil
		}
		// Get keys that end with A
		if matched, _ := regexp.MatchString("..A", key); matched {
			startA = append(startA, key)
		}
	}

	fmt.Println(startA)

	// Iterating through the map with the directions
	steps := 0
	currentKeys := startA
	mod := len(directions)
	stepsList := make([]int, len(currentKeys))
	// Bruteforcing would take too long and it can't really be parallelized
	// Thankfully the input is set up so that the length of the first loop
	// (from xxA to xxZ) is the same and subsequent loops of xxZ to xxZ
	// This isn't obvious unless you print a couple loops (which is kinda bad from a design perspective imo)
	for {
		for i, k := range currentKeys {
			if mappings, ok := mapDict[k]; ok && len(mappings) > 0 {
				currKey := mappings[int(directions[steps%mod]-'0')]
				currentKeys[i] = currKey
				if matched, _ := regexp.MatchString("..Z", currKey); matched {
					stepsList[i] = steps + 1
				}
			}
		}
		steps++
		if !containsNil(stepsList) {
			break
		}
	}

	// Since we know the length of each loop we can just find the least common multiplier
	total := lcm(stepsList...)
	fmt.Println(total)
}

func containsNil(arr []int) bool {
	for _, v := range arr {
		if v == 0 {
			return true
		}
	}
	return false
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(arr ...int) int {
	result := arr[0]
	for _, value := range arr[1:] {
		result = result * value / gcd(result, value)
	}
	return result
}

