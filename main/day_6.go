package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var input string = scanner.Text()
	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < int(number); i++ {
		scanner.Scan()
		var input string = scanner.Text()
		output1, output2 := "", ""
		var char []string = strings.Split(input, "")
		for j := 0; j < len(input); j++ {
			if (j % 2 == 0) {
				output1 = output1 + char[j]
			} else {
				output2 = output2 + char[j]
			}
		}
		fmt.Println(output1 + " " + output2)
	}
}