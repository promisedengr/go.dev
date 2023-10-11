package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseComplex(input string) (float64, float64, error) {
	re := regexp.MustCompile(`([-+]?\d*\.?\d+)([-+]?\d*\.?\d+)i`)
	match := re.FindStringSubmatch(input)

	if len(match) != 3 {
		return 0, 0, fmt.Errorf("Invalid complex number format: %s", input)
	}

	realPart, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid real part: %s", match[1])
	}

	imaginaryPart, err := strconv.ParseFloat(match[2], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid imaginary part: %s", match[2])
	}

	return realPart, imaginaryPart, nil
}

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("Usage: go run main.go complexnum1 binaryop complexnum2")
		return
	}

	complexNum1Real, complexNum1Imaginary, err := parseComplex(args[0])
	if err != nil {
		fmt.Println("Error parsing complex number 1:", err)
		return
	}

	operation := args[1]

	complexNum2Real, complexNum2Imaginary, err := parseComplex(args[2])
	if err != nil {
		fmt.Println("Error parsing complex number 2:", err)
		return
	}

	switch operation {
	case "+":
		resultReal := complexNum1Real + complexNum2Real
		resultImaginary := complexNum1Imaginary + complexNum2Imaginary
		fmt.Printf("Result: %v%vi\n", resultReal, resultImaginary)
	case "-":
		resultReal := complexNum1Real - complexNum2Real
		resultImaginary := complexNum1Imaginary - complexNum2Imaginary
		fmt.Printf("Result: %v%vi\n", resultReal, resultImaginary)
	case "*":
		resultReal := complexNum1Real*complexNum2Real - complexNum1Imaginary*complexNum2Imaginary
		resultImaginary := complexNum1Real*complexNum2Imaginary + complexNum1Imaginary*complexNum2Real
		fmt.Printf("Result: %v%vi\n", resultReal, resultImaginary)
	default:
		fmt.Println("Invalid operation:", operation)
	}
}
