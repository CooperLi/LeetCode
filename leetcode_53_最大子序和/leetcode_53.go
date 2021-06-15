package leetcode_53_最大子序和

/*
给定一个整数数组nums，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:
输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/

func maxSubArray(nums []int) int {
	// 暂定最大和是第一个元素
	max := nums[0]
	// 遍历整个数组
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i; j < len(nums); j++ {
			if j != i {
				sum += nums[j]
			}
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
