package main

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
}

func main() {
	sonOfLeft := TreeNode{nil, nil}
	daughterOfLeft := TreeNode{nil, nil}
	left := TreeNode{&daughterOfLeft, &sonOfLeft}
	sonOfRight := TreeNode{nil, nil}
	daughterOfRight := TreeNode{nil, nil}
	root := TreeNode{&left, &TreeNode{&sonOfRight, &daughterOfRight}}
	fmt.Print(getDiameter(&root))
}

func getDiameter(root *TreeNode) int {
	max := 0
	calc(root, &max)
	return max
}

func calc(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := calc(root.left, max)
	right := calc(root.right, max)

	if left+right > *max {
		*max = left + right
	}
	if left < right {
		return right + 1
	}
	return left + 1
}
