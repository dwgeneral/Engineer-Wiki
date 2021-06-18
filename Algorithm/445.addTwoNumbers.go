/*
 * 此题直观的做法可以是将两链表进行反转,然后遍历取和后再反转回来
 * 如果不修改输入链表的话, 可以利用栈的先进后出特性, 循环取和
 * 时间复杂度: O(max(n, m)) 空间复杂度:O(n+m)
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) (newHead *ListNode) {
    s1, s2, carry := pushList(l1), pushList(l2), 0
    for len(*s1) > 0 || len(*s2) > 0 {
        carry += popStack(s1) + popStack(s2)
        newHead, carry = &ListNode{Val: carry % 10, Next: newHead}, carry/10
    }
    if carry != 0 {
        newHead = &ListNode{Val: carry, Next: newHead}
    }
    return
}

func pushList(head *ListNode) *[]int {
    stack := new([]int)
    for node := head; node != nil; node = node.Next {
        *stack = append(*stack, node.Val)
    }
    return stack
}

func popStack(stack *[]int) (pop int) {
    if len(*stack) > 0 {
        pop = (*stack)[len(*stack)-1]
        *stack = (*stack)[:len(*stack)-1]
    }
    return
}