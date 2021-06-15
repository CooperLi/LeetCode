package leetcode_121_买卖股票的最佳时机

import (
	"math"
)

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
注意你不能在买入股票前卖出股票。

示例 1:
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

示例 2:
输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/

/*
遍历一次
保存更低的买入价格min和更高的profit
*/
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	profit := 0
	min := math.MaxInt64
	for i := 0; i < len(prices); i++ {
		// 有更低的价格，保存起来
		if prices[i] < min {
			min = prices[i]
		}
		// 赚的更多的情况，更新profit
		if prices[i]-min > profit {
			// 因为只允许买卖一次，所以是"=" 而不是 "+="
			profit = prices[i] - min
		}
	}
	return profit
}
