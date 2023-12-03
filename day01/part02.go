package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"log"
	"regexp"
)

func isIntChar(c int) bool {
	return c <= 57 && c >= 48;
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var digitPattern string = "one|two|three|four|five|six|seven|eight|nine";
var digitRe = regexp.MustCompile(digitPattern);

func getFirstDigitOrSpelling(targetString string) int {
	buf := "";

	for i := 0; i < len(targetString); i++ {
		buf += string(targetString[i]);

		foundDigit := digitRe.FindString(buf);

		if len(foundDigit) != 0 {
			return digits[foundDigit];
		}

		if isIntChar(int(targetString[i])) {
			val, _ := strconv.Atoi(string(targetString[i]));
			return val;
		}
	}

	fmt.Printf("NOT FOUND");
	return 0;
}


func reverseString(str string) string {
	result := "";
	for _, c := range str {
		result = string(c) + result;
	}
	return result; 
}

func getLastDigitOrSpelling(targetString string) int {
	buf := "";

	for i := len(targetString) - 1; i >= 0; i-- {
		buf += string(targetString[i]);

		foundDigit := digitRe.FindString(reverseString(buf));

		if len(foundDigit) != 0 {
			return digits[foundDigit];
		}

		if isIntChar(int(targetString[i])) {
			val, _ := strconv.Atoi(string(targetString[i]));
			return val;
		}
	}

	fmt.Printf("NOT FOUND: %s\n", targetString);
	return 0;
}

func main() {
	fptr, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err);
	}

	defer fptr.Close();

	scanner := bufio.NewScanner(fptr);
	sum := 0;

	for scanner.Scan() {
		var line string = scanner.Text();
		
		t, err := strconv.Atoi(strconv.Itoa(getFirstDigitOrSpelling(line)) + strconv.Itoa(getLastDigitOrSpelling(line)));

		 if err != nil {
		 	log.Fatal(err);
		 }

		sum += t;
	}

	fmt.Printf("The sum of all calibration values is %d\n", sum);

  if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
