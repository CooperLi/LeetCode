package leetcode_1_两数之和

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
/*
保存映射： 值 -> index
遍历映射，如果找到 target-值，说明存在这两个数字使得他们的和=target，返回这两个index
*/
func twoSum(nums []int, target int) []int {
	maps := make(map[int]int, len(nums))
	for i, b := range nums {
		if j, ok := maps[target-b]; ok {
			return []int{j, i}
		}
		maps[b] = i
	}
	return nil
}
