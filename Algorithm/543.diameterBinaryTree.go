/*
 * 递归: 先遍历左树,下探一层, 深度+1, 再遍历右树, 下探一层, 深度+1, 之和即为左子树最大深度与右子树最大深度
 * 这样得到的结果是必须穿过根节点的最长直径, 也可能存在不穿过根节点的直径大于此值, 所以需要在遍历的时候设置一个全局变量来记录过程中的最大值, 即为结果
 * 终止条件: 节点为nil
 * base case: 第一层深度为1
 */
 var m int
 func diameterOfBinaryTree(root *TreeNode) int {
	 m = 0
	 if root == nil {
		 return 0
	 }
	 dfs(root)
	 return m
 }
 
 func dfs(root *TreeNode) int {
	 if root == nil {
		 return 0
	 }
	 l, r := dfs(root.Left), dfs(root.Right)
	 m = max(m, l+r)
	 return max(l, r) + 1
 }
 
 func max(x, y int) int {
	 if x > y {
		 return x
	 }
	 return y
 }