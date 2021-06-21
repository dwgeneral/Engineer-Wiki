func reverseList(head *ListNode) (prev *ListNode) {
	for head != nil {
		nxt := head.Next
		head.Next = prev
		prev = head
		head = nxt
	}
	return
}