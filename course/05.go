package main

import (
	"fmt"
	"strings"
)

func removeConsecutiveDuplicates(input, removalChars string) string {
	result := []rune{}
	lastLetter := input[len(input)-1:]
	updatedInput := input[:len(input)-1]

	for i := 0; i < len(updatedInput); {
		char := rune(updatedInput[i])

		// Check if the character is in the removal string
		if !strings.ContainsRune(removalChars, char) {
			result = append(result, char)
			i++
		} else {
			// Skip characters in the removal string
			for i < len(updatedInput) && rune(updatedInput[i]) == char {
				i++
			}
		}
	}
	// append the last letter
	return string(result) + lastLetter
}

func main() {
	fmt.Println(removeConsecutiveDuplicates("gooddeed", "od"))  // should return "geed"
	fmt.Println(removeConsecutiveDuplicates("aabbccdd", "abc")) // should return "dd"
	fmt.Println(removeConsecutiveDuplicates("abcdefg", "xyz"))  // should return "abcdefg"
}
