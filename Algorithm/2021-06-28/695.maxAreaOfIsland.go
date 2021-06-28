package main

/*
 * 遍历网格, 如果 = 1 进入 DFS, 相邻为1累加面积, 计算完把1变成0, 防止重复计算, 最后返回最大值
 */
func maxAreaOfIsland(grid [][]int) (res int) {
	m := len(grid)
	if m == 0 {
		return
	}
	n := len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(i, j))
			}
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
