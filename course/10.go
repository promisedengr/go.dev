package main

import (
	"fmt"
	"sort"
)

// country_cap represents a country with its population.
type country_cap struct {
	country    string
	population int
}

// countryCapSlice is a slice of country_cap.
type countryCapSlice []country_cap

// Implement the sort.Interface for sorting by country name in ascending order.
func (s countryCapSlice) Len() int           { return len(s) }
func (s countryCapSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s countryCapSlice) Less(i, j int) bool { return s[i].country < s[j].country }

// Implement the sort.Interface for sorting by population in descending order.
func (s countryCapSlice) LenByPopulation() int           { return len(s) }
func (s countryCapSlice) SwapByPopulation(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s countryCapSlice) LessByPopulation(i, j int) bool { return s[i].population > s[j].population }

func main() {
	// Initialize an array of 10 elements with data.
	countries := countryCapSlice{
		{"USA", 331995364},
		{"China", 144421610},
		{"India", 138000438},
		{"Brazil", 209288278},
		{"Russia", 146599183},
		{"Pakistan", 225199937},
		{"Nigeria", 206139587},
		{"Bangladesh", 166303498},
		{"Mexico", 125959205},
		{"Japan", 125507472},
	}

	// Sort the array by country name in ascending order and print the elements.
	sort.Sort(countries)
	fmt.Println("Sorted by country name (ascending):")
	for _, c := range countries {
		fmt.Printf("Country: %s, Population: %d\n", c.country, c.population)
	}

	// Sort the array by population in descending order and print the elements.
	sort.Sort(sort.Reverse(countries))
	fmt.Println("\nSorted by population (descending):")
	for _, c := range countries {
		fmt.Printf("Country: %s, Population: %d\n", c.country, c.population)
	}
}
