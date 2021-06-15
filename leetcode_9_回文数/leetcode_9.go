package leetcode_9_回文数

import (
	"strconv"
)

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:
输入: 121
输出: true

示例 2:
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

示例 3:
输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？
*/

func isPalindrome(x int) bool {
	// 负数肯定不是回文
	if x < 0 {
		return false
	}
	value := intToSlice(x)
	size := len(value)
	start := 0
	for start < size {
		// 判断最前面和最后面两个是否相等
		if value[start] != value[size-start-1] {
			return false
		}
		start++
	}
	return true
}

func intToSlice(x int) []int {
	if x < 0 {
		x = x * -1
	}
	str := strconv.Itoa(x)
	var res = make([]int, len(str))
	for i := 0; i < len(str); i++ {
		res[i], _ = strconv.Atoi(string(str[i]))
	}
	return res
}

/*
使用纯数字解法
反转一半的数字，然后和剩下的对比
*/

func isPalindrome2(x int) bool {
	// 负数肯定不是
	// 对10取余数，如果结果为0，则表明最后一位是0，一定不是，因为第一位肯定不为0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	res := 0
	// 反转一半的数字，在判断是否相等
	// 当x小于保存的数时，表示已经到了中间
	// eg： x=12321 res=0 --> x=1232 res=1 --> x=123 res=12 --> x=12 res=123
	// 此时x小于等于res，则跳出循环去判断：偶数直接判断是否相等，奇数要先整除10在判断是否相等（因为多一位）
	for x > res {
		tmp := x % 10      // tmp总是等于x的最后一位
		x /= 10            // x循环，剔除掉最后一位
		res = res*10 + tmp // res变成x的相反数
	}
	return x == res || x == res/10
}
