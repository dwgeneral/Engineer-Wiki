package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
 * 方法一: 前后指针
 * 可以考虑使用前后指针, 前指针先往前走n步, 后指针在头节点, 二者相差n个节点, 然后二者同时前进, 前指针.next到达结尾时, 后指针即为倒数第n个节点的前驱节点, 执行删除操作即可
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := head, head
	for i := 0; i < n; i++ {
		right = right.Next
	}
	if right == nil {
		head = head.Next
		return head
	}
	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return head
}

/*
 * 方法二: 计算链表长度N, 然后遍历链表, N-n处的位置, 就是要删除的节点
 */
/*
 * 方法三: 利用栈的先进后出来做
 */
