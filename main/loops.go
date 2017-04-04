package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var input string = scanner.Text()

	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	i := 1
	for i <= 10 {
		fmt.Printf("%d x %d = %d\n", number, i, int(number)*i)
		i++
	}
}
