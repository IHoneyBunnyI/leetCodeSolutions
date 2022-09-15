package main

import (
	"fmt"
	"sort"
)

func deleteInSlice(sl *[]int, index int) {
	*sl = append((*sl)[:index], (*sl)[index+1:]...)
}

func findTwice(sl []int, elem int) (bool, int) {
	if elem == 0 && sl[1] > elem {
		return false, -1
	}
	for i, val := range sl {
		//fmt.Print(val, elem)
		if elem*2 == val {
			return true, i
		}
	}
	return false, -1
}

func findOriginalArray(changed []int) []int {
	length := len(changed)
	if length%2 != 0 {
		return []int{}
	}

	resCounter := 0
	res := make([]int, length/2)
	sort.Slice(changed, func(i, j int) bool {
		return changed[i] < changed[j]
	})
	for len(changed) != 0 {
		if ok, index := findTwice(changed, changed[0]); ok {
			//fmt.Println(" ", changed[0], index)
			res[resCounter] = changed[0]
			resCounter++
			deleteInSlice(&changed, index)
			deleteInSlice(&changed, 0)
		} else {
			return []int{}
		}
	}
	return res
}

func main() {
	//fmt.Println(findOriginalArray([]int{1, 2, 2, 4}))
	//fmt.Println(findOriginalArray([]int{6, 3, 0, 1}))
	fmt.Println(findOriginalArray([]int{6, 3, 1, 2, 0, 0}))
	fmt.Println(findOriginalArray([]int{2, 1}))
	fmt.Println(findOriginalArray([]int{1, 3, 4, 2, 6, 8}))
	//fmt.Print(sl)
}
