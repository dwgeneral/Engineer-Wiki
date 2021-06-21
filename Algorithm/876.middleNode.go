/*
 * 快慢指针 当快指针到达尾部时, 慢指针正好到达中间节点, 再基于奇偶判断一下
 * O(n) O(1)
 */
 func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    if fast == nil {
        return slow
    }
    return slow.Next
}