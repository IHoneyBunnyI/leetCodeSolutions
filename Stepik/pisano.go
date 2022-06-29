package main

import "fmt"

// 0 1 1 2 3 5

func Pisano(m int64) int64 {
	var curr int64 = 1
	var prev int64 = 0
	if m == 1 {
		return 1
	}

	for i := 0; i < int(m*m); i++ {
		prev, curr = curr, (curr+prev)%m
		if curr == 1 && prev == 0 {
			return int64(i + 1)
		}
	}
	return 0
}

func main() {
	var n int64
	var m int64
	fmt.Scan(&n)
	fmt.Scan(&m)
	if n < 0 || m < 1 {
		return
	}
	if n == 0 || n == 1 {
		fmt.Println(n)
		return
	}
	pisano := Pisano(m)
	//print(pisano)

	var curr int64 = 1
	var prev int64 = 0
	n = n % pisano
	if n == 0 {
		fmt.Println(0)
		return
	}
	for i := 0; i < int(n-1); i++ {
		prev, curr = curr, prev+curr%m
	}
	fmt.Println(curr % m)
}
