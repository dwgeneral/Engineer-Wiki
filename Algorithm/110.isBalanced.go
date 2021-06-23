/*
 * DFS: 遍历到左子树的叶子结点, 和右子树的叶子结点, 遍历的同时记录高度, 然后比较绝对值是否大于 1
 * 不仅仅需要判断根节点两边的左右子树是否平衡, 还要递归的判断内部的各个子树是否平衡
 */
 func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(node *TreeNode) int {
    if node == nil {
        return 0
    }
    return max(depth(node.Left), depth(node.Right)) + 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -1 * x
    }
    return x
}