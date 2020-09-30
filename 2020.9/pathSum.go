package main

//func main() {
//	tree1 := TreeNode{
//		Val:   2,
//		Left:  nil,
//		Right: nil,
//	}
//	tree4 := TreeNode{
//		Val:   2,
//		Left:  nil,
//		Right: nil,
//	}
//	tree2 := TreeNode{
//		Val:   2,
//		Left:  &tree4,
//		Right: &tree1,
//	}
//	tree3 := TreeNode{
//		Val:   1,
//		Left:  &tree2,
//		Right: nil,
//	}
//	fmt.Print(pathSum(&tree3, 5))
//}

func pathSum(root *TreeNode, sum int) [][]int {
	//tree := root
	res := make([][]int, 0, 10)
	flag := make([]int, 0, 10)
	for a := range res {
		res[a] = make([]int, 0, 10)
	}
	res = path(root, res, flag, 0, sum)

	return res
}

func path(tree *TreeNode, res [][]int, flag []int, cur, sum int) [][]int {
	//ok := false
	if tree == nil {
		return res
	}

	cur += tree.Val
	flag = append(flag, tree.Val)
	if tree.Right == nil && tree.Left == nil {
		if cur == sum {
			res = append(res[:], nil)
			res[len(res)-1] = append(res[len(res)-1], flag...)
			return res

		}
	}

	res = path(tree.Left, res, flag, cur, sum)
	//if ok {
	//	return res, ok
	//}
	res = path(tree.Right, res, flag, cur, sum)
	//if ok {
	//	return res, ok
	//}
	return res
}
