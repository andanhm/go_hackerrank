package main

import (
	"fmt"
	"os"
	"bufio"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("Hello, World.\n"+input)
}