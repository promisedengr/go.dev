package main

import "fmt"

func capitalizeWords(input string) string {
	words := []rune(input)
	capitalizeNext := true

	for i := range words {
		if words[i] == ' ' {
			capitalizeNext = true
		} else if capitalizeNext {
			words[i] = 'A' + (words[i] - 'a')
			capitalizeNext = false
		}
	}

	return string(words)
}

func main() {
	fmt.Println(capitalizeWords("one two three"))         // should return "One Two Three"
	fmt.Println(capitalizeWords("this is a test string")) // should return "This Is A Test String"
}
