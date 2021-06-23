/*
 * DFS: O(n) O(n)
 * 从根节点开始, 递归的对树进行遍历, 直到递归到叶子结点, 才开始左右翻转, 可以理解为对树的后序遍历, 当然前序遍历也是可以的
 * 遍历到nil结点为递归的终止条件
 */
 func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    invertTree(root.Left)
    invertTree(root.Right)
    root.Left, root.Right = root.Right, root.Left
    return root
}

/*
 * BFS O(n) O(n)
 * 层序遍历的思想, 依次将节点放入队列, 进行左右翻转
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        curr := queue[0] 
        queue = queue[1:]
        
        curr.Left, curr.Right = curr.Right, curr.Left
        if curr.Left != nil {
            queue = append(queue, curr.Left)
        }
        if curr.Right != nil {
            queue = append(queue, curr.Right)
        }
    }
    return root
}