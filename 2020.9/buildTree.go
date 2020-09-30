package main

/**中序遍历 inorder = []
后序遍历 postorder = []*/
//func main() {
//	inorder := []int{9, 3, 15, 20, 7}
//	postorder := []int{9, 15, 7, 20, 3}
//	fmt.Print(buildTree(inorder, postorder))
//}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var root *TreeNode
	root = new(TreeNode)
	if len(inorder) == 0 {
		return nil
	}
	// 后续遍历找根
	root.Val = postorder[len(postorder)-1]
	// 找到耿姐点在中序遍历中的位置，对中序遍历数组进行划分
	indexOf := Contains(inorder, postorder[len(postorder)-1])
	root.Right = buildTree(inorder[indexOf+1:], postorder[indexOf:len(postorder)-1])
	root.Left = buildTree(inorder[:indexOf], postorder[:indexOf])
	return root
}

func Contains(slice []int, s int) int {
	for index, value := range slice {
		if value == s {
			return index
		}
	}
	return -1
}
