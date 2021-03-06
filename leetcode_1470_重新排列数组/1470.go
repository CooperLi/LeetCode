package leetcode_1470_重新排列数组

/*
给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。
请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组。

示例 1：
输入：nums = [2,5,1,3,4,7], n = 3
输出：[2,3,5,4,1,7]
解释：由于 x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 ，所以答案为 [2,3,5,4,1,7]

示例 2：
输入：nums = [1,2,3,4,4,3,2,1], n = 4
输出：[1,4,2,3,3,2,4,1]

示例 3：
输入：nums = [1,1,2,2], n = 2
输出：[1,2,1,2]

提示：
1 <= n <= 500
nums.length == 2n
1 <= nums[i] <= 10^3
*/

/*
思考：
1. x1 下标为 i，y1 下标为 i+n
2. 新数组的偶数下标位置放x，奇数下标位置放y
*/
func shuffle(nums []int, n int) []int {
	if len(nums) == 0 || n == 1 {
		return nums
	}
	var result = make([]int, len(nums))
	for i := 0; i < n; i++ {
		// 偶数位
		result[2*i] = nums[i]
		// 奇数位
		result[2*i+1] = nums[i+n]
	}
	return result
}
