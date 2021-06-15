package leetcode_1213_三个有序数组的交集

import (
	"sort"
)

/*
给出三个均为 严格递增排列 的整数数组 arr1，arr2 和 arr3。
返回一个由 仅 在这三个数组中 同时出现 的整数所构成的有序数组。

示例：
输入: arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9], arr3 = [1,3,4,5,8]
输出: [1,5]
解释: 只有 1 和 5 同时在这三个数组中出现.

提示：
1 <= arr1.length, arr2.length, arr3.length <= 1000
1 <= arr1[i], arr2[i], arr3[i] <= 2000
*/

/*
思路：构造一个map，遍历三个数组，看map里面的值哪个是3，就是谁
*/

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	// 保存映射关系：
	// key  : arr里的每个数字
	// value: 出现的次数
	var keyMap = make(map[int]int, 0)
	var res []int
	// 遍历三个数组，每遍历一个数字，该数字对应的出现次数+1
	for _, v := range arr1 {
		keyMap[v]++
	}
	for _, v := range arr2 {
		keyMap[v]++
	}
	for _, v := range arr3 {
		keyMap[v]++
	}
	// 找到所有值是3的，也就是三个数组中都有的数字
	for i, v := range keyMap {
		if v == 3 {
			// 添加key到结果里
			res = append(res, i)
		}
	}
	// 因为要求有序，所以排下序
	sort.Ints(res)
	return res
}
