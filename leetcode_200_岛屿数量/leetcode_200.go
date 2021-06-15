package leetcode_200_岛屿数量

/*
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:
输入:
11110
11010
11000
00000
输出: 1

示例 2:
输入:
11000
11000
00100
00011
输出: 3
*/

/*
思路一：深度优先遍历DFS
目标是找到矩阵中 “岛屿的数量” ，上下左右相连的 1 都被认为是连续岛屿。
dfs方法： 设目前指针指向一个岛屿中的某一点 (i, j)，寻找包括此点的岛屿边界。
从 (i, j) 向此点的上下左右 (i+1,j),(i-1,j),(i,j+1),(i,j-1) 做深度搜索。
终止条件：
    (i, j) 越过矩阵边界;
    grid[i][j] == 0，代表此分支已越过岛屿边界。
    搜索岛屿的同时，执行 grid[i][j] = '0'，即将岛屿所有节点删除，以免之后重复搜索相同岛屿。
主循环：
    遍历整个矩阵，当遇到 grid[i][j] == '1' 时，从此点开始做深度优先搜索 dfs，岛屿数 count + 1 且在深度优先搜索中删除此岛屿。
最终返回岛屿数 count 即可。
 */
func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
        return
	}
	grid[i][j] = '0'
	dfs(grid, i+1,j)    // 上
	dfs(grid, i-1,j)    // 下
	dfs(grid, i, j-1)   // 左
	dfs(grid, i, j+1)   // 右
}

func numIslands(grid [][]byte) int {
    num := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                dfs(grid, i, j)
                num++
            }
        }
    }
    return num
}
