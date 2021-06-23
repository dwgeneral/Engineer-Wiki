/*
 * DFS: 遍历二叉树, 如果遍历到叶子结点, 就将路径加入结果集, 返回.
 * 如果是非叶子结点, 就拼接path, 然后继续向下递归
 */
 func binaryTreePaths(root *TreeNode) (res []string) {
    dfs(root, "", &res)
    return
}

func dfs(node *TreeNode, path string, res *[]string) {
    if node == nil {
        return
    } 
    if node.Left == nil && node.Right == nil {
        path += strconv.Itoa(node.Val)
        *res = append(*res, path)
        return
    }
    path += strconv.Itoa(node.Val) + "->"
    dfs(node.Left, path, res)
    dfs(node.Right, path, res)
}