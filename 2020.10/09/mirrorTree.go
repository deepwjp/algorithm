package main

//剑指 Offer 27. 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Right, root.Left = root.Left, root.Right
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
