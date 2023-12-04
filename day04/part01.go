package main
	
import (
	"fmt"
	"log"
	"bufio"
	"os"
	"regexp"
	"math"
	"strings"
)

func getPointsForLine(line string) int {
	_, cardPoints, _ := strings.Cut(line, ": ");
	givenPoints, winningPoints, _ := strings.Cut(cardPoints, " | ");

	powCount := 0;

	re := regexp.MustCompile("([0-9][0-9]|[0-9])");

	for _, givenPoint := range re.FindAllString(givenPoints, -1) {
		for _, winningPoint := range re.FindAllString(winningPoints, -1) {
			if strings.TrimSpace(givenPoint) == strings.TrimSpace(winningPoint) {
				powCount ++;
			}
		}
	}


	if powCount <= 2 {
		return powCount;
	}
	return int(math.Pow(2, float64(powCount - 1)));
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
		sum += int(getPointsForLine(line));
	}

	fmt.Printf("Total points of the cards %d\n", sum);


  if err := fscanner.Err(); err != nil {
		log.Fatal(err)
	}
}
