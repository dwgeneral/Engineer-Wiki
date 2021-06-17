/*
 * 这道题可以使用递归和迭代两种方法来实现, 考虑到递归存在递归调用栈的开销,
 * 所以考虑使用迭代方式来做, 时间复杂度为: O(n) 空间复杂度: O(1)
 * 1 -> 2 -> 4
 * 1 -> 3 -> 4
 * 1 -> 1 -> 2 -> 3 -> 4 -> 4
 */
// 迭代
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return dummy.Next
}
