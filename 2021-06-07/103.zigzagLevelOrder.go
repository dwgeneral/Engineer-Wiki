/*
 * 先通过BFS算法进行二叉树的层序遍历, 然后针对每一层的值进行奇偶性判断来控制访问方向
 * 时间复杂度为 O(n); n 为节点个数; 空间复杂度为: O(n)
 */
func zigzagLevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		level := make([]int, 0)
		q := len(queue)
		for j := 0; j < q; j++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if i%2 != 0 {
			for j := 0; j < q/2; j++ {
				level[j], level[q-j-1] = level[q-j-1], level[j]
			}
		}

		res = append(res, level)
	}
	return
}
