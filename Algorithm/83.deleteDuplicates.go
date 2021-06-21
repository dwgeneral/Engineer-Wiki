/*
 * 快慢指针, O(n), O(1)
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head.Next
	for ; fast != nil; fast = fast.Next {
		if slow.Val == fast.Val {
			slow.Next = fast.Next
		} else {
			slow = slow.Next
		}
	}
	return head
}