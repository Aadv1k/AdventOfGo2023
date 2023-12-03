package main
	
import (
	"fmt"
	"log"
	"strconv"
	"bufio"
	"os"
)

func isIntChar(c int) bool {
	return c <= 57 && c >= 48;
}

func getSumOfLine(line string) string {
	first, last := 0, 0

	for i := 0; i < len(line); i++ {
		ch := int(line[i]);
		if isIntChar(ch) {
			first = ch;

			for j := len(line) - 1; j >= i; j-- {
				ch := int(line[j]);
				if isIntChar(ch) {
					last = ch;
					break;
				}
			}
			break;
		}
	}
	return string(first) + string(last);
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
		i, _ := strconv.Atoi(getSumOfLine(line));
		sum += i;
	}

	fmt.Printf("The sum of all calibration values is %d\n", sum);


  if err := fscanner.Err(); err != nil {
		log.Fatal(err)
	}
}
