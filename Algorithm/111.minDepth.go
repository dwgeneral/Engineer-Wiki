/*
 * DFS
 * 首先可以想到使用深度优先搜索的方法，遍历整棵树，记录最小深度。
 * 对于每一个非叶子节点，我们只需要分别计算其左右子树的最小叶子节点深度。这样就将一个大问题转化为了小问题，可以递归地解决该问题。
 */
 func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }
    minD := math.MaxInt32
    if root.Left != nil {
        minD = min(minDepth(root.Left), minD)
    }
    if root.Right != nil {
        minD = min(minDepth(root.Right), minD)
    }
    return minD + 1
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}