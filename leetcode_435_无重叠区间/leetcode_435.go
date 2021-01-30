package leetcode_435_无重叠区间

import (
	"sort"
)

/*
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:
可以认为区间的终点总是大于它的起点。
区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

示例 1:
输入: [ [1,2], [2,3], [3,4], [1,3] ]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。

示例 2:
输入: [ [1,2], [1,2], [1,2] ]
输出: 2
解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。

示例 3:
输入: [ [1,2], [2,3] ]
输出: 0
解释: 你不需要移除任何区间，因为它们已经是无重叠的了。
*/

/*
我认为区间调度的核心：
1. 按照end或者start的值从小到大将区间排序
2. 按照规则遍历区间，进行处理
*/

// 【灵魂】计算不重叠区间
func intervalSchedule(intvs [][]int) int {
	if len(intvs) == 0 {
		return 0
	}

	// 按照end升序排列
	sort.Slice(intvs, func(i, j int) bool {
		return intvs[i][1] < intvs[j][1]
	})

	// 至少有一个区间不相交
	count := 1
	// 排序过后第一个区间就是x
	xEnd := intvs[0][1]

	for _, v := range intvs {
		start := v[0]
		// 开始的数字大于等于上一个结束的，所以说明是不重叠的
		if start >= xEnd {
			count++
			// 更新xEnd为下一个无重叠的end
			xEnd = v[1]
		}
	}

	return count
}

func eraseOverlapIntervals(intervals [][]int) int {
	// 已经找到了不重叠区间的个数，所以总的减去这个就得到了要剔除的个数
	return len(intervals) - intervalSchedule(intervals)
}
