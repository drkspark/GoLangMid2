package main

import "fmt"

func main() {
	var num, rem, temp int
	var reverse = 0
	fmt.Scan(&num)
	temp = num

	for {
		rem = num % 10
		reverse = reverse*10 + rem
		num /= 10
		if num == 0 {
			break
		}
	}

	if temp == reverse {
		fmt.Printf("The given number %d is a palindrome", temp)
	} else {
		fmt.Printf("The given number %d is not a palindrome", temp)
	}
}
