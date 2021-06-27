package main

/*
 * 这道题考查的主要是链表元素的操作, 我们可以遍历两链表元素, 依次相加, 然后记录进位情况
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	for carry := 0; carry != 0 || l1 != nil || l2 != nil; curr = curr.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: carry % 10}
		carry /= 10
	}
	return dummy.Next
}
