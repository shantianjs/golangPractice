package practice

import "math"

func numSquares(n int) int {
	//记录平方数
	length := int(math.Sqrt(float64(n))) + 1
	mapNums := make(map[int]bool, length)
	seqNums := make([]int, length)
	i := 1
	for i <= length{
		mapNums[i*i] = true
		seqNums[i - 1] = i * i
		i++
	}
	//从结果为数开始试探，最多到n
	for i := 1; i <= n; i++ {
		if isDivided(n, i, mapNums, seqNums) {
			return i
		}
	}
	return -1
}

//最多四次
func isDivided(targe, count int, s map[int]bool, seq []int) bool {
	if count == 1{
		_, ok := s[targe]
		if ok {
			return true
		}
		return false
	}
	for _, num := range seq {
		if num > targe {
			break
		}
		if isDivided(targe - num, count - 1, s, seq) {
			return true
		}
	}
	return false
}
