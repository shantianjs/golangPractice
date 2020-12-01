package practice

func coinChange(coins []int, amount int) int {
	record := make([]int, amount+1, amount+1)
	idx := 1
	//idx代表amount
	for idx <= amount {
		res := -1
		for _, coin := range coins {
			//开始amount比coin面值小，和无解的
			if coin > idx || record[idx - coin] == -1{
				continue
			}
			//初次初始化，后面进行比较
			if res == -1 {
				res = record[idx - coin] + 1
			} else {
				if res > record[idx - coin] + 1 {
					res = record[idx - coin] + 1
				}
			}
		}
		record[idx] = res
		idx++
	}
	return record[len(record)-1]
}
