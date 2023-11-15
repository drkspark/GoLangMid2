package main

import "fmt"

func main() {
	var row int
	fmt.Print("Enter number of rows: ")
	fmt.Scanln(&row)
	spaces := row - 1

	for i := 0; i < row; i++ {
		for j := 0; j < spaces; j++ {
			fmt.Printf("   ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Printf(" * ")
		}
		fmt.Println()
		spaces--
	}
}
