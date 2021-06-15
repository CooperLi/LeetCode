package leetcode_695_岛屿的最大面积

/*
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

示例 1:
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

示例 2:
[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。

注意: 给定的矩阵grid 的长度和宽度都不超过 50。
*/

var visited [][]bool         // 是否已访问
var dx = [4]int{-1, 0, 1, 0} // 方向
var dy = [4]int{0, -1, 0, 1} // 方向
var n, m int                 // 地图大小

func check(grid [][]int, x, y int) bool {
	if x < 0 || x >= n || y < 0 || y >= m { // 错误的路径，边界
		return true
	}
	if grid[x][y] == 0 {
		return true
	}
	if visited[x][y] {
		return true
	}
	return false
}

func dfs(grid [][]int, x, y int) int {
	if check(grid, x, y) { // 遇到边界，返回
		return 0
	}
	visited[x][y] = true // 记录已经访问
	area := 1
	for i := 0; i < 4; i++ { // 遍历四个方向，并累计面积
		area += dfs(grid, x+dx[i], y+dy[i])
	}
	return area
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0 // 记录最大面积
	visited = [][]bool{}
	n, m = len(grid), len(grid[0])
	for range make([]int, n) { // 初始化空的已访问数组
		visited = append(visited, make([]bool, m))
	}
	for i := 0; i < n; i++ { // 遍历地图
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !visited[i][j] { // 如果没被访问过，而且是土地（1），那么对该地进行dfs
				area := dfs(grid, i, j)
				if area > maxArea { // 保存最大值
					maxArea = area
				}
			}
		}
	}
	return maxArea
}
