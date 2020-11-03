package 剑指offer

/*
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。



示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL


限制：

0 <= 节点个数 <= 5000 */
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
