package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxRight(height []int, left int) int {
	max := height[left]
	index := left + 1
	lenHeight := len(height)
	for i := index; i < lenHeight; i++ {
		if height[i] >= max {
			return i
		}
	}

	//index = left + 1
	max = height[left+1]
	for i := index; i < lenHeight; i++ {
		if height[i] >= max {
			max = height[i]
			index = i
		}
	}
	if index == left+1 {
		return -1
	}
	return index
}

func trap(height []int) int {
	res := 0
	left := 0
	right := -1
	for i, v := range height { //skip all 0 in start and find left
		if v != 0 {
			left = i
			break
		}
	}
	maxIndex := len(height) - 1
	for left != maxIndex {

		right = maxRight(height, left)
		//fmt.Print("l: ", left, " r: ", right, " r - l: ", right-left, "\n")
		if right != -1 {
			//fmt.Print("l: ", left, " r: ", right, " r - l: ", right-left, "\n")
			if right-left > 1 {
				res += min(height[left], height[right]) * (right - left - 1)
				for i := left + 1; i < right; i++ {
					res -= height[i]
				}
			}
			left = right
		} else {
			left++
		}
		//right = -1
	}
	return res
}

func main() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 0, 0, 0, 0}
	Slice(input).Print()
	fmt.Println("Res:= ", trap(input), "\n")
	input = []int{0, 0, 0, 0, 4, 0, 0, 3, 0, 0, 0}
	Slice(input).Print()
	fmt.Println("Res:= ", trap(input), "\n")
	input = []int{0, 0, 0, 9, 4, 5, 9, 0, 0}
	Slice(input).Print()
	fmt.Println("Res:= ", trap(input), "\n")

	input = []int{0, 0, 0, 4, 1, 5, 9, 0, 8}
	Slice(input).Print()
	fmt.Println("Res:= ", trap(input), "\n")
	//PrintSlice(input)
	//slice := [][]byte{{'.', '.', '#'}, {'.', '#', '#'}, {'.', '.', '#'}}
	//for i := range slice {
	//for j := range slice {
	//fmt.Print(string(slice[j][i]))
	//}
	//fmt.Println()
	//}
}
