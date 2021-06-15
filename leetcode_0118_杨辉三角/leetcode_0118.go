package leetcode_0118_杨辉三角

import (
	"fmt"
)

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行
在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:
输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	if numRows == 0 {
		return res
	}
	// 第一行的一个1
	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}
	// 处理剩余行
	for i := 1; i < numRows; i++ {
		res = append(res, genNext(res[i-1]))
	}
	return res
}

func genNext(p []int) []int {
	res := make([]int, 1, len(p)+1)
	// 先扩增一位, 第0位即为0
	res = append(res, p...)
	for i := 0; i < len(res)-1; i++ {
		// 每一位更新成: 现在的值 + 下一位的值
		res[i] = res[i] + res[i+1]
		fmt.Println(res)
	}
	fmt.Println("=========", res)
	return res
}
