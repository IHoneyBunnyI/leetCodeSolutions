package main

func maxArea(height []int) int {
	right := len(height) - 1
	left := 0
	max := 0
	tmp := 0
	if height[left] < height[right] {
		max = height[left] * (right - left)
	} else {
		max = height[right] * (right - left)
	}
	for left != right {
		if height[left] < height[right] {
			tmp = height[left] * (right - left)
		} else {
			tmp = height[right] * (right - left)
		}
		if tmp > max {
			max = tmp
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	print(maxArea(a), "\n")
	a = []int{1, 1}
	print(maxArea(a), "\n")
	a = []int{1, 4, 9, 4, 5}
	print(maxArea(a), "\n")
}
