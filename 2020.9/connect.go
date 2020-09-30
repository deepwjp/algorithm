package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	maps := make(map[int]*Node)
	maps = dfsConnect(root, 0, maps)
	return root
}
func dfsConnect(root *Node, depth int, maps map[int]*Node) map[int]*Node {
	if root == nil {
		return maps
	}
	dfsConnect(root.Right, depth+1, maps)
	dfsConnect(root.Left, depth+1, maps)
	if _, ok := maps[depth]; ok {
		root.Next = maps[depth]
	} else {
		root.Next = nil
	}
	maps[depth] = root
	return maps
}
