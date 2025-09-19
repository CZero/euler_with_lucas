package main

import "fmt"

func main() {
	limit := 1000
	valid := rangeUntil(limit)
	fmt.Printf("Sum of numbers divisable until %d = %d\n", limit, sumAll(valid))
}

func rangeUntil(number int) (valid []int) {
	for i := 0; i < number; i++ {
		if isMultipleOf3Or5(i) {
			valid = append(valid, i)
		}
	}
	return
}

func isMultipleOf3Or5(number int) bool {
	switch {
	case number < 3:
		return false
	case number%3 == 0:
		return true
	case number%5 == 0:
		return true
	default:
		return false
	}
}

func sumAll(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}
