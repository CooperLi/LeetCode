package leetcode_1389

/*
给你两个整数数组 nums 和 index。你需要按照以下规则创建目标数组：

目标数组 target 最初为空。
按从左到右的顺序依次读取 nums[i] 和 index[i]，在 target 数组中的下标 index[i] 处插入值 nums[i] 。
重复上一步，直到在 nums 和 index 中都没有要读取的元素。
请你返回目标数组。

题目保证数字插入位置总是存在。

示例 1：
输入：nums = [0,1,2,3,4], index = [0,1,2,2,1]
输出：[0,4,1,3,2]
解释：
nums       index     target
0            0        [0]
1            1        [0,1]
2            2        [0,1,2]
3            2        [0,1,3,2]
4            1        [0,4,1,3,2]

示例 2：
输入：nums = [1,2,3,4,0], index = [0,1,2,3,0]
输出：[0,1,2,3,4]
解释：
nums       index     target
1            0        [1]
2            1        [1,2]
3            2        [1,2,3]
4            3        [1,2,3,4]
0            0        [0,1,2,3,4]

示例 3：
输入：nums = [1], index = [0]
输出：[1]

提示：
1 <= nums.length, index.length <= 100
nums.length == index.length
0 <= nums[i] <= 100
0 <= index[i] <= i
*/

/*
思考:
暴力法：考数组的插入，insert a value in a slice at a given index. 方法：
copy(arr[i+1:],arr[i:])
arr[i] = value
也就是把 index 之后的元素整体往后移动一个单位，然后把 index 这里的元素插入到 index 位置
*/
func createTargetArray(nums []int, index []int) []int {
	var target = make([]int, len(nums))
	for k, v := range index {
		copy(target[v+1:], target[v:])
		target[v] = nums[k]
	}
	// return target[:len(nums)]
	return target
}
