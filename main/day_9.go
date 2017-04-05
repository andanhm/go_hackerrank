package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var input string = scanner.Text()

	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	var factorial int = fact(int(number))
	fmt.Printf("%d",factorial)
}
