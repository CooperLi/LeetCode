package leetcode_0035_搜索插入位置

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。

示例 1:
输入: [1,3,5,6], 5
输出: 2

示例 2:
输入: [1,3,5,6], 2
输出: 1

示例 3:
输入: [1,3,5,6], 7
输出: 4

示例 4:
输入: [1,3,5,6], 0
输出: 0
*/

/*
思路：
二分法查找，如果num[mid]=target, return target
如果 num[mid] >= target, 在 left 和 mid - 1之间找
如果 num[mid] < target, 在 right 和 mid + 1 之间找
*/

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums) // 定义target在左闭右开的区间里，[left, right)  target
	for left < right {	// 因为left == right的时候，在[left, right)是无效的空间
		mid := left + (right - left) >> 1	// 固定写法，防止溢出
		if nums[mid] > target {
			right = mid	// target 在左区间，在[left, middle)中, 这里 right 不需要-1，因为本身 mid 这里就不包含
		} else if nums[mid] < target {
			left = mid + 1	// target 在右区间，在 [middle+1, right)中
		} else {
			right = mid	// 数组中找到目标值的情况，直接返回下标
		}
	}
	// 分别处理如下四种情况
	// 目标值在数组所有元素之前 [0,0)
	// 目标值等于数组中某一个元素 return middle
	// 目标值插入数组中的位置 [left, right) ，return right 即可
	// 目标值在数组所有元素之后的情况 [left, right)，return right 即可
	return right
}

