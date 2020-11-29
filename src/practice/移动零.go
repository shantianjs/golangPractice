func moveZeroes(nums []int)  {
	low, fast := 0,0
	for fast < len(nums) {
		if nums[low] == 0 {
			fast++
			if fast > low && num[fast] != 0{
				nums[low], nums[fast] = nums[fast], nums[low]
			}
		} else {
			low++
		}
	}
}
