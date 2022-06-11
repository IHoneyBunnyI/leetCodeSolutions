package main

import "fmt"

type Slice []int

func (s Slice) max() int {
	max := s[0]
	for _, val := range s {
		if val > max {
			max = val
		}
	}
	return max
}

func (s Slice) Print() {
	out := make([][]byte, len(s))
	max := s.max()
	for i := range out {
		out[i] = make([]byte, max)
		for j := range out[i] {
			out[i][j] = '.'
		}
	}
	for i := range out {
		for j := 0; j < s[i]; j++ {
			out[i][max-j-1] = '#'
		}
	}
	for j := 0; j < max; j++ {
		for i := range out {
			fmt.Print(string(out[i][j]))
		}
		fmt.Println()
	}
}
