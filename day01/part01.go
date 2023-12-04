package main
	
import (
	"fmt"
	"log"
	"strconv"
	"bufio"
	"os"
	"unicode"
)

func isDigit(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func SumOfLine(line string) (int, error) {
	first, last := 0, 0

	for i := 0; i < len(line); i++ {
		ch := line[i]
		if isDigit(ch) {
			first = int(ch)

			for j := len(line) - 1; j >= i; j-- {
				ch := line[j]
				if isDigit(ch) {
					last = int(ch)
					break
				}
			}
			break
		}
	}

	convertedDigit, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
	if err != nil {
		return 0, fmt.Errorf("unable to convert string to integer: %w", err)
	}

	return convertedDigit, nil
}

func main() {
	fptr, err := os.Open("input.txt");
	if err != nil {
		log.Fatal(err);
	}

	defer fptr.Close();

	fscanner := bufio.NewScanner(fptr);

	sum := 0;
	for fscanner.Scan() {
		var line string = fscanner.Text();

		calibrationValue, err := SumOfLine(line);

		if err != nil {
			log.Fatal(err);
		}

		sum += calibrationValue;
	}

	fmt.Printf("Part01: The sum of all calibration values is %d\n", sum);


  if err := fscanner.Err(); err != nil {
		log.Fatal(err)
	}
}
