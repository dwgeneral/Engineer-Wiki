/*
 * 直观的做法是将此链表翻转成新链表, 逐一比对两个链表即可, 但这样空间复杂度为O(n)
 * 考虑可以从链表头部到中间节点进行翻转, 然后前后逐一比对, 这样就不用新建链表了
 * O(n) O(1)
 */
 func isPalindrome(head *ListNode) bool {
    slow, fast, newHead := head, head, head
    for fast != nil {
        if fast.Next == nil {
           slow = slow.Next 
           break 
        }
        fast = fast.Next.Next
        prev := slow
        slow = slow.Next
        prev.Next = newHead
        newHead = prev
    }
    for slow != nil {
        if slow.Val != newHead.Val {
            return false
        }
        slow = slow.Next
        newHead = newHead.Next
    }
    return true
}