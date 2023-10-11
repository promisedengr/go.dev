package main

import "fmt"

var words = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

func num2words(num int) string {
	if num == 0 {
		return "zero"
	}

	result := ""
	base := 1 // Base for powers of 10

	for num > 0 {
		digit := num % 10
		result = words[digit] + result // Prepend word representation
		num /= 10

		if num > 0 {
			result = " " + result // Add space if more digits follow
		}

		base *= 10
	}

	return result
}

func main() {
	fmt.Println(num2words(123))  // should return "one two three"
	fmt.Println(num2words(4567)) // should return "four five six seven"
	fmt.Println(num2words(0))    // should return "zero"
}
