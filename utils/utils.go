package utils

import (
	"os"
	"strconv"
	"log"
	"strings"
	"unicode"
)

func IsDigit(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func ReadFileIntoString(filepath string) (string, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func SplitLines(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n' || r == '\r'
	})
}


func Min(values ...int) int {
	if len(values) == 0 {
		log.Fatal("min: empty slice")
	}

	minValue := values[0]
	for _, value := range values[1:] {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}
