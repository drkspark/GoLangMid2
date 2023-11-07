package main

import (
	"fmt"
	"strings"
)

func main() {
	// For Characters
	fmt.Println(strings.ContainsAny("Germany", "G"))
	fmt.Println(strings.ContainsAny("Germany", "g"))

	// For substring
	fmt.Println(strings.Contains("India", "In"))
	fmt.Println(strings.Contains("Germany", "ger"))
	fmt.Println(strings.Contains("Japan", "Jap"))

	// Returns number of non-overlapping substrings
	fmt.Println(strings.Count("malayalam", "m"))
	fmt.Println(strings.Count("malayalam", "o"))

	// Check for equality case-insensitive
	fmt.Println(strings.EqualFold("Cat", "cAt"))
	fmt.Println(strings.EqualFold("India", "Indiana"))
}
