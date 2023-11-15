package main

import (
	"fmt"
	"math"
)

func main() {
	var num [10]float64
	var sum, mean, sd float64
	sum = 0
	fmt.Println("Enter 10 elements: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter %d element: ", i+1)
		fmt.Scan(&num[i])
		sum += num[i]
	}
	mean = sum / 10

	for j := 0; j < 10; j++ {
		sd += math.Pow(num[j]-mean, 2)
	}

	sd = math.Sqrt(sd / 10)
	fmt.Println("Standard deviation of given elements is: ", sd)
}
