package main

import (
	"fmt"
	"os"
	"strconv"
)

// num2hex converts an unsigned integer to its hexa equivalent as a string.
// If upperCase is true, it returns the hexa value in uppercase.
func num2hex(n uint64, upperCase bool) string {
	var format string
	if upperCase {
		format = "%X"
	} else {
		format = "%x"
	}
	return fmt.Sprintf(format, n)
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <number> <-u>")
		return
	}

	number, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		fmt.Println("Invalid number:", args[0])
		return
	}

	option := args[1]

	if option == "-u" {
		hexaValue := num2hex(number, true)
		fmt.Println(hexaValue)
	} else {
		hexaValue := num2hex(number, false)
		fmt.Println(hexaValue)
		fmt.Println("Invalid option:", option)
	}
}
