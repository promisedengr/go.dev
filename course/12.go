package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	wordGroups := make(map[int][]string)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter words (separated by spaces):")

	// Read a line of input
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into words
		words := strings.Fields(line)

		// Process each word
		for _, word := range words {
			word = strings.TrimSpace(word)
			wordLen := len(word)

			// Only consider words with at least 4 characters
			if wordLen >= 4 {
				wordGroups[wordLen] = append(wordGroups[wordLen], word)
			}
		}

		// Continue reading input lines
		fmt.Println("Enter more words or press Ctrl+D (Ctrl+Z on Windows) to finish:")
	}

	// Sort and print words by length
	var sortedLengths []int
	for wordLen := range wordGroups {
		sortedLengths = append(sortedLengths, wordLen)
	}
	sort.Ints(sortedLengths)

	for _, wordLen := range sortedLengths {
		words := wordGroups[wordLen]
		sort.Strings(words)

		fmt.Printf("%d: %d-letter words in alphabetical order\n", wordLen, wordLen)
		for _, word := range words {
			fmt.Println(word)
		}
	}
}
