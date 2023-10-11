package main

import "fmt"

var wordToDigit = map[string]int{
	"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
	"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

var multiWordToDigit = map[string]int{
	"ten": 10, "eleven": 11, "twelve": 12, "thirteen": 13, "fourteen": 14,
	"fifteen": 15, "sixteen": 16, "seventeen": 17, "eighteen": 18, "nineteen": 19,
	"twenty": 20, "thirty": 30, "forty": 40, "fifty": 50, "sixty": 60,
	"seventy": 70, "eighty": 80, "ninety": 90,
}

func words2num(input string) int {
	result := 0
	wordStart := 0

	for i := 0; i <= len(input); i++ {
		if i == len(input) || input[i] == ' ' {
			word := input[wordStart:i]

			if digit, found := wordToDigit[word]; found {
				result = result*10 + digit
			} else if digit, found := multiWordToDigit[word]; found {
				result += digit
			} else {
				fmt.Printf("Error: Unrecognized word '%s'\n", word)
				return -1 // Return an error code or handle the error as needed
			}

			wordStart = i + 1
		}
	}

	return result
}

func main() {
	fmt.Println(words2num("one two three"))         // should return 123
	fmt.Println(words2num("five six seven"))        // should return 567
	fmt.Println(words2num("twenty two four seven")) // should return 2247
}
