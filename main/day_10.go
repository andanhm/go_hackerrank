package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var input string = scanner.Text()

	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	var numberInBinary string =  strconv.FormatInt(number, 2)
	var maxConsecutiveOneNum int = 0;
	var consecutiveOneNum int = 0;
	for i := 0; i < len(numberInBinary); i++ {
		if (string([]rune(numberInBinary)[i]) == "1") {
			consecutiveOneNum++;
			maxConsecutiveOneNum = int(math.Max(float64(maxConsecutiveOneNum), float64(consecutiveOneNum)));
		} else {
			consecutiveOneNum = 0;
		}
	}
	fmt.Printf("%d",maxConsecutiveOneNum);
}
