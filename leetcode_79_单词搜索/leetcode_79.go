package leetcode_79_单词搜索

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:
boards =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 words = "ABCCED", 返回 true.
给定 words = "SEE", 返回 true.
给定 words = "ABCB", 返回 false.
*/

//           x-1,y
//            ↑
//  x,y-1  ← x,y →  x,y+1
//            ↓
//           x+1,y
var direction = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
var boards [][]byte
var visited [][]bool
var words string
var n, m int

func exist(board [][]byte, word string) bool {
	n = len(board)
	if n == 0 {
		return false
	}
	m = len(board[0])
	words = word
	boards = board
	visited = buildTwoDimArray(n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func backtrack(i, j, curPos int) bool {
	// 是否判断到字符最后一个，退出条件
	if curPos == len(words)-1 {
		return boards[i][j] == words[curPos]
	}
	// 进入条件
	if boards[i][j] == words[curPos] {
		// 置为访问过
		visited[i][j] = true
		// 四个方向
		for k := 0; k < 4; k++ {
			x := i + direction[k][0]
			y := j + direction[k][1]
			// 先判断在不在范围，再看看是不是访问过了，否则顺序反过来，就会越界
			if inArea(x, y) && !visited[x][y] {
				if backtrack(x, y, curPos+1) {
					return true
				}
			}
		}
		// 恢复，回溯必备
		visited[i][j] = false
	}
	return false
}

func inArea(x, y int) bool {
	return x >= 0 && y >= 0 && x < n && y < m
}

func buildTwoDimArray(n, m int) [][]bool {
	a := make([][]bool, n)
	for i := range a {
		a[i] = make([]bool, m)
	}
	return a
}
