package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)
// We can also do with interface which is also a better solution

type PhoneBook struct {
	Name   string
	Number int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var input string = scanner.Text()

	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	var phoneBook map[string]string
	phoneBook = make(map[string]string)
	for i := 0; i < int(number); i++ {
		scanner.Scan()
		var input []string = strings.Split(scanner.Text(), " ")
		if (len(input) == 2) {
			var name, phoneNo string = input[0], input[1];
			phoneBook[name] = phoneNo
		}
	}
	for i := 0; i < int(number); i++ {
		scanner.Scan()
		var name string = scanner.Text()
		if name != "" {
			if phoneNumber, ok := phoneBook[name]; ok {
				fmt.Printf("%s=%s\n", name, phoneNumber)
			} else {
				fmt.Println("Not found")
			}
		}
	}
}