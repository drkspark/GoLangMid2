package main

import "fmt"

func main() {
	var rows int
	val := 1
	fmt.Print("Enter number of rows: ")
	fmt.Scan(&rows)

	for i := 0; i < rows; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf(" %d ", val)
			val++
		}
		fmt.Println()
	}
}
