package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	sum := nums[0]
	res := sum
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		res = max(res, sum)
	}
	return res
}

func main() {
	fmt.Println(maxSubArray([]int{0, 1, 2, 3}))
}
