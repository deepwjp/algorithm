package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func main() {
//	arr := []int{-10, -3, 0, 5, 9}
//	fmt.Println(sortedArrayToBST(arr))
//}

func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode
	root = solceSortedArrayToBST(root, nums)
	return root
}

func solceSortedArrayToBST(root *TreeNode, nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root = new(TreeNode)
	if len(nums) == 1 {
		root.Val = nums[0]
		root.Right = nil
		root.Left = nil
		return root
	}
	root.Val = nums[len(nums)/2]
	root.Left = solceSortedArrayToBST(root.Left, nums[:len(nums)/2])
	root.Right = solceSortedArrayToBST(root.Right, nums[len(nums)/2+1:])
	return root
}
