package leetcode_1137_第N个泰波那契数

/*
泰波那契序列 Tn 定义如下：
T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

示例 1：
输入：n = 4
输出：4
解释：
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4

示例 2：
输入：n = 25
输出：1389537
*/

/*
dp数组存取之前求过的结果，当下次要的时候直接取
*/
func tribonacci(n int) int {
	var dp = make([]int, 38)
	if dp[n] != 0 {
		return dp[n]
	}
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		res := tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
		dp[n] = res
		return res
	}
}
