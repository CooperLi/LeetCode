package leetcode_1099_小于K的两数之和

import (
	"math"
	"sort"
)

/*
给你一个整数数组 A 和一个整数 K，请在该数组中找出两个元素，使它们的和小于 K 但尽可能地接近 K，返回这两个元素的和。
如不存在这样的两个元素，请返回 -1。

示例 1：
输入：A = [34,23,1,24,75,33,54,8], K = 60
输出：58
解释：
34 和 24 相加得到 58，58 小于 60，满足题意。

示例 2：
输入：A = [10,20,30], K = 15
输出：-1
解释：
我们无法找到和小于 15 的两个元素。
*/

func twoSumLessThanK(A []int, K int) int {
	if A == nil || len(A) == 0 {
		return -1
	}
	// 先确保从小到大排序
	sort.Ints(A)
	i, j := 0, len(A)-1
	// max 先取最小值
	max := -1 << 63
	// 结束条件是i=j，两个指针遇到一起
	for i != j {
		// 目前两个指针的和
		tmp := A[i] + A[j]
		// 和大于等于K，右边指针左移一位，同时跳过此次循环
		if tmp >= K { // 注意，必须是大于等于，否则会有坑
			j--
			continue
		} else { // 和小于K，说明存在可能的最接近K，先保存起来
			// max = max(tmp, max)
			// 如果目前两指针之和大约max，则替换
			if tmp >= max {
				max = tmp
			}
			// 如果小于max，则不替换max，继续比较，也就是i++
			i++
		}
	}
	if max == math.MinInt64 {
		return -1
	}
	return max
}
