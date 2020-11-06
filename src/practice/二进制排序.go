package practice

import "sort"

func sortByBits(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	res := make([]int, len(arr),len(arr))
	bitMap := make(map[int][]int, len(arr))
	calBitNum := func(num int) int {
		b := 0
		for num > 0 {
			num = num & (num - 1)
			b++
		}
		return b
	}
	//获取bit对应的数组
	for _, num := range arr {
		bits := calBitNum(num)
		bitMap[bits] = append(bitMap[bits], num)
	}
	//获取bit列表，排序
	keys := make([]int, 0,len(bitMap))
	for key := range bitMap {
		sort.Ints(bitMap[key])
		keys = append(keys, key)
	}
	sort.Ints(keys)
	//放入统一数组
	j := 0
	for _, key := range keys{
		copy(res[j:], bitMap[key])
		j += len(bitMap[key])
	}
	return res
}
