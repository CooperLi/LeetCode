package leetcode_1512

/*
给你一个整数数组 nums 。
如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。

返回好数对的数目。

示例 1：
输入：nums = [1,2,3,1,1,3]
输出：4
解释：有 4 组好数对，分别是 (0,3), (0,4), (3,4), (2,5) ，下标从 0 开始

示例 2：
输入：nums = [1,1,1,1]
输出：6
解释：数组中的每组数字都是好数对

示例 3：
输入：nums = [1,2,3]
输出：0

提示：
1 <= nums.length <= 100
1 <= nums[i] <= 100
*/

/*
思路：
无脑题，两个指针，第二个比第一个快 1
相等结果就+1，最后遍历完输出结果
如果要求输出坐标数据，那么就创建一个二维数组，每次遇到相等的坐标就记录进去
 */

func numIdenticalPairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return len(result)
}
