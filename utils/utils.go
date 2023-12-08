package utils

import (
	"os"
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

	return string(content), nil
}
