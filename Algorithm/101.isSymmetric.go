/*
 * DFS
 * 递归遍历这棵树, 从根节点出发, 判断左右子树是否相等, 如果相等, 再次下探一层判断
 * 递归终止条件为 节点为 nil
 */
 func isSymmetric(root *TreeNode) bool {
    return check(root, root)
}

func check(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil || right == nil {
        return false
    }
    return left.Val == right.Val && check(left.Left, right.Right) && check(left.Right, right.Left)
}

/*
 * BFS
 * 层序遍历这棵树, 对每一层的节点按照对称顺序放入队列, 依次进行比较
 * 注意当判断到两节点都为nil时, 应跳过该节点继续遍历其余节点, 而不是直接返回 false
 */
func isSymmetric(root *TreeNode) bool {
    queue := []*TreeNode{root, root}
    for len(queue) > 0 {
        left, right := queue[0], queue[1]
        queue = queue[2:]    
        if left == nil && right == nil {
            continue
        }
        if left == nil || right == nil {
            return false
        }
        if left.Val != right.Val {
            return false
        }
        queue = append(queue, left.Left)
        queue = append(queue, right.Right)
        queue = append(queue, left.Right)
        queue = append(queue, right.Left)
    }
    return true
}