package practice

//含冷冻期一天
func maxProfitCool(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	//第n天，0无股票 1持有股票的利润
	dp := make([][2]int, len(prices), len(prices))
	//初始状态
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		//当天不持有股票状态由，昨天不持有和，昨天不持有今天卖掉转换
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		if i > 1 {
			//当天持有由昨天持有，和冷冻期后天购买转化
			dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
		} else {
			//昨天持有和昨天没有，今天买
			dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
		}
	}
	//最后一天卖出股票，不要留在手上
	return dp[len(dp)-1][0]
}
