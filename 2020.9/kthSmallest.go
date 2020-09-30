package main

func kthSmallest(root *TreeNode, k int) int {
	arr := []int{}
	arr = dfsKthSmallest(root, arr)
	return arr[k-1]
}

func dfsKthSmallest(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}

	arr = dfsKthSmallest(root.Left, arr)
	arr = append(arr, root.Val)
	arr = dfsKthSmallest(root.Right, arr)

	return arr
}
