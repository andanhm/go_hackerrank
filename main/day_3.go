package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"log"
	"math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var input string = scanner.Text()

	number, err := strconv.ParseFloat(input, 64); if err != nil {
		log.Fatal(err)
	}
	if math.Mod(number, 2.0) == 1 {
		fmt.Println("Weird")
	} else {
		if (number >= 2 && number <= 5) {
			fmt.Println("Not Weird")
		} else if (number >= 6&&number <= 20) {
			fmt.Println("Weird")
		}
		if (number > 20) {
			fmt.Println("Not Weird")
		}
	}
}
