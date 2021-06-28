package main

/*
 * 遍历网格, 如果遇到 = 1 的点, 进入 DFS, 将所有联通的1改为0, 同时记录+1, 然后回到主循环, 继续遍历,
 * 直到结束
 */
func numIslands(grid [][]byte) (count int) {
	m := len(grid)
	if m == 0 {
		return
	}
	n := len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				count++
			}
		}
	}
	return
}
