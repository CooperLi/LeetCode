package leetcode_46_全排列

/*
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

/*
回溯，顾名思义，要往回走，为啥往回走，因为没路可走（走到底了）
那这里就有几个地方需要思考：

怎么走？即选择
怎么回来？即撤销选择
什么时候结束？即退出条件
想清楚这几个，我们就可以来做题了
套用著名的回溯公式(from labuladong)：


result = []
func backtrack(路径，选择列表) {
	if 满足结束条件 {
		result.add(路径)
	}
	return

	for 选择 in 选择列表 {
		做选择
		backtrack(路径，选择列表)
		撤销选择
	}
}
*/

var res [][]int

func permute(nums []int) [][]int {
	var track []int
	res = [][]int{}
	backtrack(nums, track)
	return res
}

func backtrack(nums []int, track []int) {
	if len(track) == len(nums) { // 递归出口，我们找到的和原长度相等
		tmp := make([]int, len(track))
		copy(tmp, track)       // 安全拷贝
		res = append(res, tmp) // 添加进最终结果
		return
	}
	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}
		track = append(track, nums[i]) // 先添加进我们的预选表中
		backtrack(nums, track)         // 调用回溯
		track = track[:len(track)-1]   // 恢复状态
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
