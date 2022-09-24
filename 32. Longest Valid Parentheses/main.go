package main

import (
	"fmt"
)

type Stack struct {
	data []int
}

func (s *Stack) Push(index int) {
	s.data = append(s.data, index)
}

func (s Stack) Empty() bool {
	return (len(s.data) == 0)
}

func (s *Stack) Pop() {
	if s.Empty() {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s Stack) Top() int {
	if len(s.data) != 0 {
		return s.data[len(s.data)-1]
	}
	return 0
}

func newStack() Stack {
	return Stack{data: make([]int, 0)}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValid(str string) int {
	var res int
	stack := newStack()
	stack.Push(-1)
	for index, val := range str {
		if val == '(' {
			stack.Push(index)
		} else {
			stack.Pop()
			if stack.Empty() {
				stack.Push(index)
			} else {
				res = max(res, index-stack.Top())
			}
		}
	}
	return res
}

func main() {
	var res int
	var str string
	str = "()"
	res = longestValid(str)
	fmt.Println(str, ":", res)
	//str = "((((((("
	//res = longestValid(str)
	//fmt.Println(str, ":", res)
	//str = ")))))))"
	//res = longestValid(str)
	//fmt.Println(str, ":", res)
	//str = "))()))"
	//res = longestValid(str)
	//fmt.Println(str, ":", res)
	//str = "())((())())"
	//res = longestValid(str)
	//fmt.Println(str, ":", res)
	//str = ")()())"
	//res = longestValid(str)
	//fmt.Println(str, ":", res)
	//str = ""
	//res = longestValid(str)
	//fmt.Println(str, ":", res)
	//str = ")()())"
	//res = longestValid(str)
	//fmt.Println(str, ":", res)

	str = ")()(()"
	res = longestValid(str)
	fmt.Println(str, ":", res)
	str = ")(((()"
	res = longestValid(str)
	fmt.Println(str, ":", res)
}
