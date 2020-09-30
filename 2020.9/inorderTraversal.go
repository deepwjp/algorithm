package main

func inorderTraversal(root *TreeNode) []int {
	arr := []int{}
	arr = dfsInorderTraversal(root, arr)
	return arr
}

func dfsInorderTraversal(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = dfsInorderTraversal(root.Left, arr)
	arr = append(arr, root.Val)
	arr = dfsInorderTraversal(root.Right, arr)
	return arr
}
