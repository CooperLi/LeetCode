package leetcode_70_爬楼梯

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

示例 1：
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶

示例 2：
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
*/

/*
如果一个问题的最优解包含了其中子问题的最优解，那么称其具有最优子结构的性质。
什么意思？青蛙在面对 n 个台阶时的解决方案数是 f(n)，那么我们知道 f(n)=f(n−1)+f(n−2)。
其中的 f(n−1) 与 f(n−2) 就是两个子问题的最优解，此时我们可以理解成一个问题的最优解包含了其子问题的最优解，那么这个时候这种问题具有了最优子结构性质。

状态转移方程：
dp[i] = dp[i - 1] + dp[i - 2]
i 代表当前问题的规模，即所需要跳过的台阶数。
dp[i] 代表的是跳过 i 个台阶的方案数量
*/
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
