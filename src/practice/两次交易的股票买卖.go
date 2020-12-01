package practice

func maxProfitIIII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	//天数，第几笔交易，当前是否存在股票
	dp := make([][2][2]int, len(prices), len(prices))
	dp[0][0][0] = 0
	dp[0][0][1] = -prices[0]
	dp[0][1][0] = 0
	dp[0][1][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		//之前没有和当前卖出
		dp[i][0][0] = max(dp[i-1][0][0], dp[i-1][0][1] + prices[i])
		//之前买入和现在买入
		dp[i][0][1] = max(dp[i-1][0][1], - prices[i])
		//从第一次无股票买入
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - prices[i])
		//从第二次卖出
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])

	}
	return dp[len(dp)-1][1][0]
}
