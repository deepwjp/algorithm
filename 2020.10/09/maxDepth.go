package main

//剑指 Offer 55 - I. 二叉树的深度 && 104
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := solves(root, 0)
	return depth

}

func solves(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth += 1
	tempLeft := solves(root.Left, depth)

	tempRight := solves(root.Right, depth)
	if tempLeft > tempRight {
		return tempLeft
	}
	return tempRight
}
