package practice

import "math"

//递归加备忘录加二分法搜索
func superEggDrop(K int, N int) int {
	record := map[[2]int]int{}
	var dp func(k, n int) int
	dp = func(k, n int) int {
		if k == 1 || n == 1 {
			return n
		}
		if k == 0 || n == 0 {
			return 0
		}
		if num, ok := record[[2]int{k,n}];ok{
			return num
		}
		//比每层扔还大
		res := math.MaxInt32
		lo, hi := 0, n
		for lo <= hi {
			mid := (lo + hi) /2
			notBroken := dp(k, n - mid)
			broken := dp(k-1, mid - 1)
			//楼层过小
			if broken < notBroken {
				lo = mid + 1
				if res > notBroken + 1 {
					res = notBroken + 1
				}
			} else {
				hi = mid - 1
				if res > broken + 1 {
					res = broken + 1
				}
			}
		}
		record[[2]int{k,n}] = res
		return res
	}
	return dp(K, N)
}
