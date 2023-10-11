package main

import (
	"fmt"
	"math"
)

// Function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	maxPrimes := 120
	primeSlices := make([][]int, maxPrimes)

	primeCount := 0
	currentSliceSize := 1

	for primeCount < maxPrimes {
		currentSlice := make([]int, currentSliceSize)
		for i := 2; ; i++ {
			if isPrime(i) {
				currentSlice[primeCount%currentSliceSize] = i
				primeCount++
				if primeCount%currentSliceSize == 0 {
					break
				}
			}
		}
		primeSlices[primeCount-1] = currentSlice
		currentSliceSize++
	}

	// Print the contents of each slice
	for i, slice := range primeSlices {
		fmt.Printf("Slice-%d: %v\n", i+1, slice)
	}
}
