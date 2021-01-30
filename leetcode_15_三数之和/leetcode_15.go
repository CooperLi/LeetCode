package leetcode_15_三数之和

import (
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
/*
k i j的顺序
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	res := make([][]int, 0)
	for k := range nums[:size-2] {
		if nums[k] > 0 {
			// 因为 j > i > k
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			// 跳过重复元素，当k-1和k一样，就跳过k
			continue
		}
		i, j := k+1, size-1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			switch {
			case sum < 0:
				i++
				// 去重
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			case sum > 0:
				j--
				// 去重
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			default:
				res = append(res, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
	}
	return res
}
