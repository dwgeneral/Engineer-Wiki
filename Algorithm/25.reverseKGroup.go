/*
 * 考虑用循环来界定k个数, 然后循环调用反转链表函数
 * 反转链表逻辑使用头插法进行位置变换
 * 时间复杂度: O(n); 空间复杂度 O(n)
 * 1 -> 2 -> 3 -> 4 -> 5; k = 2
 * 2 -> 1 -> 4 -> 3 -> 5
 */
func reverseKGroup(head *ListNode, k int) (newHead *ListNode) {
	curr := head
	for i := 0; i < k; i++ {
		if curr == nil {
			return head
		}
		curr = curr.Next
	}
	newHead = reverse(head, curr)
	head.Next = reverseKGroup(curr, k)
	return
}

func reverse(head *ListNode, tail *ListNode) *ListNode {
	prev := tail
	for head != tail {
		nxt := head.Next
		head.Next = prev
		prev = head
		head = nxt
	}
	return prev
}
