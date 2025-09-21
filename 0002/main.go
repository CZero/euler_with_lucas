package main

import (
	"fmt"
)

func main() {
	fibnumbers := fibonacciBelow(4000000)
	fmt.Printf("Sum: %d\n", sumEvenNumbers(fibnumbers))
}

func sumEvenNumbers(fibnumbers []int) (sum int) {
	for _, number := range fibnumbers {
		switch number % 2 {
		case 0:
			fmt.Printf("%d + %d = %d\n", sum, number, sum+number)
			sum += number
		}
	}
	return sum
}

func fibonacciBelow(limit int) (fibnumbers []int) {
	fibnumbers = []int{1, 2}
	for {
		next := fibnumbers[len(fibnumbers)-2] + fibnumbers[len(fibnumbers)-1]
		if next >= limit {
			return fibnumbers
		}
		fibnumbers = append(fibnumbers, next)
	}
}
