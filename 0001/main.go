package main

// type Divisable map[int]bool

// func (s Divisable) has(num int) bool {
// 	_, ok := s[num]
// 	return ok
// }

func main() {
	rightnumbers := Divisable{}
}

func IsMultipleOf3Or5(number int) bool {
	switch {
	case number < 3:
		return false
	case number%3 == 0:
		return true
	case number%5 == 0:
		return true
	}
	return true
}
