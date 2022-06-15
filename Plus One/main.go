package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			return digits
		}
	}
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9, 9, 0}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}
