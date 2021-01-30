package leetcode_704_二分查找

/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

示例 2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
*/

/*
左=0
右=长度-1
循环 左小于等于右
    中间=（左+右）/2， 这里写成 (left+right) >>1 是为了防止溢出
    如果中间刚好等于目标，就返回中间
    如果小于，就取右半边，于是左变成mid往右移动一位：left = mid + 1
    如果大于，就取左半边，于是右变成mid往左移动一位：right = mid - 1
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
