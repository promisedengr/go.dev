package main

import "fmt"

func maxPower(M, N int) int {
	K := 0
	power := 1
	for M >= power {
		power *= N
		K++
	}
	return K - 1
}

func main() {
	fmt.Println(maxPower(80000, 5)) // should return 7
	fmt.Println(maxPower(30000, 9)) // should return 4
}
