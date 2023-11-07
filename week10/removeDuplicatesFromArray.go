package main

import "fmt"

func removeDuplicates(s []string) []string {
	bucket := make(map[string]bool)
	var result []string
	// _ : index
	for _, str := range s {
		ok := bucket[str]
		if !ok {
			bucket[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {
	array := []string{"hai", "my", "csd", "batch", "section", "a", "2020", "and", "hai", "csm"}
	fmt.Println("The given array of string is: ", array)
	fmt.Println()

	// Calling the Function
	result := removeDuplicates(array)
	fmt.Println("The array obtained after removing the duplicate entries is: ", result)
}
