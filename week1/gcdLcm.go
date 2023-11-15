package main

import "fmt"

func gcd(n1 int, n2 int) int {
	if n1 == 0 {
		return n2
	}
	return gcd(n2%n1, n1)
}

func lcm(n1 int, n2 int) int {
	g := gcd(n1, n2)
	return (n1 * n2) / g
}

func main() {
	var n1, n2, operation int
	fmt.Println("Enter 2 positive numbers:")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)

	fmt.Print("Enter 1 for LCM and 2 for GCD: ")
	fmt.Scanln(&operation)

	if operation == 1 {
		val := lcm(n1, n2)
		fmt.Printf("The LCM of %d and %d is %d", n1, n2, val)
	} else {
		val := gcd(n1, n2)
		fmt.Printf("The GCD of %d and %d is %d", n1, n2, val)
	}

}
