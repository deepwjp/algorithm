package main

//剑指 Offer 22. 链表中倒数第k个节点

//保持两指针间距为k
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var fast = head
	for k > 0 && fast != nil {
		fast = fast.Next
		k--
	}
	var slow = head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
