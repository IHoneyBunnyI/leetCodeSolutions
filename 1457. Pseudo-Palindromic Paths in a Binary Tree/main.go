package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numElems(nodes []int, elem int) int {
	count := 0
	for _, val := range nodes {
		if val == elem {
			count++
		}
	}
	return count
}

func IsPseudoPalindrom(nodes []int) bool {
	Palindrom := 0
	fmt.Println(nodes)
	for i := 0; i < 10; i++ {
		if numElems(nodes, i)%2 == 1 {
			Palindrom++
			if Palindrom > 1 {
				return false
			}
		}
	}
	return true
}

func deepSearch(root *TreeNode, path int, count *int) {
	if root == nil {
		return
	}
	path = path ^ (1 << root.Val)
	//fmt.Println("PATH", *path)
	if root.Left == nil && root.Right == nil {
		//fmt.Println(root.Val, path)
		if (path & (path - 1)) == 0 {
			*count++
		}
	}
	deepSearch(root.Left, path, count)
	deepSearch(root.Right, path, count)
}

func pseudoPalindromicPaths(root *TreeNode) int {
	//nodes := make([]int, 0)
	res := 0
	path := 0
	deepSearch(root, path, &res)
	println(res)
	return 0
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 1}
	//root.Right = &TreeNode{Val: 1}
	//root.Left = &TreeNode{Val: 1}
	//root.Left.Left = &TreeNode{Val: 1}
	//root.Left.Right = &TreeNode{Val: 3}
	//root.Left.Right.Right = &TreeNode{Val: 1}
	pseudoPalindromicPaths(root)
}

// Встраивание Каналы, вг ВАИТ груп Деадлок
