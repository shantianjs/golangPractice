package practice

func smallerNumbersThanCurrent(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	recode := make(map[int]int, len(nums))
	r := make([]int, len(nums))
	for i, num := range nums{
		if res, ok := recode[num]; !ok{
			for j,num1 := range nums {
				if i != j && num1 < num {
					r[i]++
				}
			}
			recode[num] = r[i]
		} else {
			r[i] = res
		}
	}
	return r
}
