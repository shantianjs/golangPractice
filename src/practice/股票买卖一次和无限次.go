package practice


func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profix := 0
	minPrice := prices[0]
	for _, num := range prices {
		if num < minPrice {
			minPrice = num
		}
		if num-minPrice > profix {
			profix = num - minPrice
		}
	}
	return profix
}

//根据状态机的递归做法，大量数据会超时
func maxProfitII(prices []int) int {
	//初始状态,cur是否持有股票
	var grow func(day, profit int, has int) int
	grow = func(day, profit, stat int) int {
		if day >= len(prices) {
			return profit
		}
		//当前持有股票
		if stat == 1 {
			//继续持有
			s1 := grow(day+1, profit, 1)
			//卖出
			s2 := grow(day+1, profit+prices[day], 0)
			return max(s1, s2)
		} else if stat == 0 {
			//当前无股票
			//买入
			s1 := grow(day+1, profit-prices[day], 1)
			//不买入
			s2 := grow(day+1, profit, 0)
			return max(s1, s2)
		}
		return -1
	}
	return grow(0, 0, 0)
}

//根据状态机的递推公式
func maxProfitIII(prices []int) int {
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
		//当天持有由昨天持有，和今天购买转化
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	}
	//最后一天卖出股票，不要留在手上
	return dp[len(dp)-1][0]
}
