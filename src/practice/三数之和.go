package practice

import "sort"

func twoSum(nums []int, start, target int) [][]int {
	if len(nums) < 2 {
		return nil
	}
	i, j := start, len(nums) - 1
	res := [][]int{}
	for i < j {
		tem := nums[i] + nums[j]
		left := nums[i]
		right := nums[j]
		if tem == target {
			res = append(res, []int{left, right})
			//跳过重复元素
			for i < j && nums[i] == left {
				i++
			}
			for i < j && nums[j] == right {
				j--
			}
		} else if tem < target{
			i++
		} else {
			j--
		}

	}
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		for _, pair := range twoSum(nums, i+1, -nums[i]) {
			res = append(res, []int{nums[i], pair[0], pair[1]})
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
