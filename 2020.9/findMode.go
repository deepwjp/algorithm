package main

//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
////func main() {
////	tree1 := TreeNode{
////		Val:   2,
////		Left:  nil,
////		Right: nil,
////	}
////	tree2 := TreeNode{
////		Val:   2,
////		Left:  nil,
////		Right: &tree1,
////	}
////	tree3 := TreeNode{
////		Val:   1,
////		Left:  &tree2,
////		Right: nil,
////	}
////	fmt.Print(findMode(&tree3))
////}
//
//func findMode(root *TreeNode) []int {
//	hashMap := make(map[int]int)
//	_, tempValue := 0, 0
//	arr := []int{}
//	solveFindMode(root, hashMap)
//	for key, value := range hashMap {
//		if value > tempValue {
//			tempValue = value
//			arr = []int{key}
//		} else if value == tempValue {
//			arr = append(arr, key)
//		}
//	}
//	return arr
//}
//
//func solveFindMode(tree *TreeNode, hashMap map[int]int) {
//	if tree == nil {
//		return
//	}
//	hashMap[tree.Val]++
//	solveFindMode(tree.Left, hashMap)
//	solveFindMode(tree.Right, hashMap)
//	return
//}
