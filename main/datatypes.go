package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	var integerInput string
	var doubleInput string
	var stringInput string
	scanner.Scan()
	integerInput = scanner.Text()
	scanner.Scan()
	doubleInput = scanner.Text()
	scanner.Scan()
	stringInput = scanner.Text()

	integer, _ := strconv.ParseUint(integerInput, 10, 0)
	double,_:=strconv.ParseFloat(doubleInput,10)
	fmt.Println(i + integer)
	fmt.Printf("%.1f\n",d + double)
	fmt.Println(s + stringInput)
}