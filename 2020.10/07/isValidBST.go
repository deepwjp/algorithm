package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) (res bool) {
	sort := []int{}
	dfsIsValidBST(root, &sort)
	for i := 0; i < len(sort)-1; i++ {
		if sort[i] > sort[i+1] {
			return false
		}
	}
	return true
}

func dfsIsValidBST(root *TreeNode, sort *[]int) {
	if root == nil {
		return
	}

	dfsIsValidBST(root.Left, sort)
	*sort = append(*sort, root.Val)
	dfsIsValidBST(root.Right, sort)
}
