package main

import "fmt"

func main() {
	var first, second string
	fmt.Print("Enter first string: ")
	fmt.Scan(&first)

	fmt.Print("Enter second string: ")
	fmt.Scan(&second)

	fmt.Println(first + second)
}
