package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Date struct {
	Day   int
	Month int
	Year  int
}

type EmpInfo struct {
	Empno        int
	DateOfJoin   Date
	LocationCode string
}

func printEmpByLocation(employees []EmpInfo, locationCode string) {
	fmt.Printf("Employees at location %s:\n", locationCode)
	for _, emp := range employees {
		if emp.LocationCode == locationCode {
			fmt.Printf("Empno: %05d, Date of Join: %02d/%02d/%02d\n", emp.Empno, emp.DateOfJoin.Day, emp.DateOfJoin.Month, emp.DateOfJoin.Year)
		}
	}
}

func printEmpByYearOfJoin(employees []EmpInfo, year int) {
	fmt.Printf("Employees who joined in %d:\n", year)
	for _, emp := range employees {
		if emp.DateOfJoin.Year == year {
			fmt.Printf("Empno: %05d, Date of Join: %02d/%02d/%02d, Location: %s\n", emp.Empno, emp.DateOfJoin.Day, emp.DateOfJoin.Month, emp.DateOfJoin.Year, emp.LocationCode)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create an array of 10 EmpInfo structures with random data
	var employees [10]EmpInfo
	for i := 0; i < 10; i++ {
		employees[i] = EmpInfo{
			Empno:        rand.Intn(100000),
			DateOfJoin:   Date{Day: rand.Intn(31) + 1, Month: rand.Intn(12) + 1, Year: rand.Intn(100)},
			LocationCode: fmt.Sprintf("%c%c%c%c", rand.Intn(26)+'A', rand.Intn(26)+'A', rand.Intn(26)+'A', rand.Intn(26)+'A'),
		}
	}

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: go run main.go -I locationcode OR -y year")
		return
	}

	switch args[0] {
	case "-I":
		locationCode := args[1]
		printEmpByLocation(employees[:], locationCode)
	case "-y":
		year, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid year argument:", args[1])
			return
		}
		printEmpByYearOfJoin(employees[:], year)
	default:
		fmt.Println("Invalid option:", args[0])
	}
}
