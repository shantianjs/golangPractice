package practice

func permute(nums []int) [][]int {
	res := [][]int{}
	var search func(searchNums, cur []int)
	search = func(searchNums,  cur []int) {
		if len(searchNums) == 0 {
			res = append(res, cur)
			return
		}
		for idx, num := range searchNums {
			//直接修改searchNum，会导致当前层idx后面数字，被修改
			//cur只增加后面的不修改前面的，返回临时列表
			tem := make([]int, len(searchNums)-1, len(searchNums)-1)
			copy(tem, searchNums[:idx])
			copy(tem[idx:], searchNums[idx+1:])
			search(tem, append(cur, num))
		}
	}
	search(nums[:],[]int{})
	return res
}
