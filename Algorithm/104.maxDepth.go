/*
 * DFS: 递归遍历左右子数各自的深度, 取较大者, 累计即为树的最大深度
 * O(n) O(n)
 */
 func maxDepth(root *TreeNode) (depth int) {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

/*
 * BFS: 层序遍历, 每一层有节点就深度+1, 直到遍历完所有节点
 * O(n) O(n)
 */
func maxDepth(root *TreeNode) (depth int) {
    if root == nil {
        return 0
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        level := len(queue)
        for level > 0 {
            curr := queue[0]
            queue = queue[1:]
            if curr.Left != nil {
                queue = append(queue, curr.Left)
            }
            if curr.Right != nil {
                queue = append(queue, curr.Right)
            }
            level--
        }
        depth++
    }
    return depth
}