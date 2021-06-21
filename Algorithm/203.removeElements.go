/*
 * 遍历节点, 如果满足条件就移除该节点, 考虑到头节点也可能被移除, 所以添加哨兵节点 dummy 来定位新的头节点
 * O(n) O(1)
 */
 func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return head
    }
    dummy := &ListNode{Val: 0, Next: head}
    prev := dummy
    for head != nil {
        if head.Val == val {
            prev.Next = head.Next
        } else {
            prev = head
        }
        head = head.Next
    }
    return dummy.Next
}

/*
 * 还可以考虑使用递归, 因为每一个节点做的事都是一样的, 都是要和val比较, 相等则返回下一个节点, 终止条件是节点为nil
 * O(n) O(n)
 */
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
       return nil 
    }
    head.Next = removeElements(head.Next, val)
    if head.Val == val {
        return head.Next
    }
    return head
}