package leetcode_0322_零钱兑换

func coinChange(coins []int, amount int) int {
    // 这是由于 dp[amount] 最大不可能超过 amount（最小面值为 1 元），所以 amount + 1 就是一个无意义的数了。
	var dp = make([]int, amount+1, amount+1)
	dp[0] = 0
	for i:=0; i<len(dp);i++ {
	    cost := 1<<63-1
	    // to solve sub problem
	    for _, coin := range coins {
            if i-coin >= 0 {
                cost = min(cost, dp[i-coin]+1)
            }
        }
        dp[i] = cost
    }
	if dp[amount] > amount + 1 {
	    return -1
    }
	return dp[amount]
}

func min(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}
