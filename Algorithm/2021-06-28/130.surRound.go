package main

/*
 * 遍历矩阵, 如果 = 'o' 就 DFS, 四个方向判断, 如果到达了边界, 就什么也不做, 退回到主循环
 * 如果没有到达边界, 就把所有的 'o' 改成 x
 * 还有一个问题, 如何区分边界O及相邻O. 这里可以在程序开始时, 先遍历下四个边, 如果有O就DFS标记为R
 * 最后在遍历整个矩阵时, 把所有的 R 写回为 O, 所有的 O 改为 X
 */
func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}
	for i := 1; i < n-1; i++ {
		dfs(board, 0, i)
		dfs(board, m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'R' {
				board[i][j] = 'O'
			}
		}
	}
	return
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'R'
	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)
}
