package practice


//简单二分法，从小到大，无重复元素
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = left + 1
		}
	}
	return -1
}


//左右边界查找
func searchRange(nums []int, target int) []int {
	res := make([]int, 2, 2)
	res[0] = leftBound(nums, target)
	res[1] = rightBound(nums, target)
	return res
}

func leftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle
		}
	}
	//越界检查
	if len(nums) <= left || left < 0|| nums[left] != target{
		return -1
	}
	return left
}

func rightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			left = middle + 1
		} else if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle
		}
	}
	//单个时,left可能为0，越界
	if len(nums) < left || left - 1 < 0 || nums[left-1] != target{
		return -1
	}
	return left - 1
}