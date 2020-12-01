package practice

func reversePairs(nums []int) int {
	res := 0
	doulNums := make([]int, len(nums),len(nums))
	for idx, num := range nums {
		doulNums[idx] = 2 * num
	}
	for idx, num := range nums {
		for j := idx+1; j < len(nums); j++ {
			if num > doulNums[j] {
				res++
			}
		}
	}
	return res
}

