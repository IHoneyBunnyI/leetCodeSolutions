package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	start := len(s) - 1
	for i := start; i >= 0; i-- {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\r' || s[i] == '\v' || s[i] == '\t' {
			continue
		} else {
			start = i
			break
		}
	}
	for i := start; i >= 0; i-- {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\r' || s[i] == '\v' || s[i] == '\t' {
			return length
		}
		length++
	}
	return length
}

func main() {
	fmt.Println(lengthOfLastWord("h  "))
	fmt.Println(lengthOfLastWord("hll world"))
	fmt.Println(lengthOfLastWord("h a"))
	fmt.Println(lengthOfLastWord("h asad "))
}
