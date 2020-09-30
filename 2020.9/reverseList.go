package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead *ListNode
	newHead = new(ListNode)
	newHead.Next = nil
	newHead.Val = head.Val
	head = head.Next
	for head != nil {
		temp := ListNode{
			Val:  head.Val,
			Next: newHead,
		}
		newHead = &temp
		head = head.Next
	}
	return newHead
}
