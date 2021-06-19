/*
 * 最直接想到的办法是用 hashmap 记录一下走过的节点, 当出时现第一个走过的节点时, 即为结果; 
 * 但空间复杂度为 O(n); 不符合要求
 * 还可以使用快慢指针来解决, fast=2*slow, 当slow与fast第一次相遇后,让slow重新回到起点, 
 * slow在环外, fast在环内, 二者再同频走, 则再次相遇时即为入环的第一个节点
 */
 func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }    
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next 
        fast = fast.Next.Next
        if slow == fast {
            slow = head
            for slow != fast {
                slow = slow.Next
                fast = fast.Next
            }
            return slow
        } 
    }
    return nil
}