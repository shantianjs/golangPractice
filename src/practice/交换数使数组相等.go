package practice

func arrayPairSum(nums []int) int {
	nlen := len(nums)/2
	return nlen
}

func findSwapValues(array1 []int, array2 []int) []int {
	sumA, sumB := 0, 0
	record := make(map[int]bool,0)
	for _, num := range array1 {
		sumA += num
		record[num] = true
	}
	for _, num := range array2 {
		sumB += num
	}
	temp := (sumA - sumB) / 2
	if (sumA - sumB) % 2 != 0 {
		return []int{}
	}
	for _, num := range array2 {
		if _, ok := record[temp + num]; ok {
			return []int{temp + num, num}
		}
	}
	return []int{}
}