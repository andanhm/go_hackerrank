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
	//var arr = make([]int, int(number))
	scanner.Scan()
	var arrInput string = scanner.Text()
	var arr []string = strings.Split(arrInput, " ")
	fmt.Println(arr)
	var output string = ""
	for i:=int(number)-1;i>=0;i-- {
		output = output + arr[i] +" "
	}
	fmt.Println(output)
}