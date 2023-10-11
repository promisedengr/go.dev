package main

import (
	"bufio"
	"fmt"
	"os"
)

// splitCopy copies odd-numbered lines to destFile1 and even-numbered lines to destFile2.
func splitCopy(srcFile, destFile1, destFile2 *os.File) error {
	scanner := bufio.NewScanner(srcFile)
	writer1 := bufio.NewWriter(destFile1)
	writer2 := bufio.NewWriter(destFile2)

	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()

		if lineNumber%2 == 1 {
			_, err := writer1.WriteString(line + "\n")
			if err != nil {
				return err
			}
			writer1.Flush()
		} else {
			_, err := writer2.WriteString(line + "\n")
			if err != nil {
				return err
			}
			writer2.Flush()
		}

		lineNumber++
	}

	return nil
}

func main() {
	// Check if the correct number of command-line arguments are provided
	if len(os.Args) != 4 {
		fmt.Println("Usage: main sourceFile destFile1 destFile2")
		// current command: go run 11.go handle.txt handle1.txt handle2.txt
		return
	}

	sourceFile := os.Args[1]
	destFile1 := os.Args[2]
	destFile2 := os.Args[3]

	// Open the source file
	src, err := os.Open(sourceFile)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer src.Close()

	// Open the destination files for writing
	dest1, err := os.Create(destFile1)
	if err != nil {
		fmt.Println("Error creating destFile1:", err)
		return
	}
	defer dest1.Close()

	dest2, err := os.Create(destFile2)
	if err != nil {
		fmt.Println("Error creating destFile2:", err)
		return
	}
	defer dest2.Close()

	// Call splitCopy to copy odd and even numbered lines
	err = splitCopy(src, dest1, dest2)
	if err != nil {
		fmt.Println("Error copying lines:", err)
	}
}
