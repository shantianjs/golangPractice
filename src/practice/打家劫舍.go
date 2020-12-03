package practice

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	per := nums[0]
	cur := max(nums[0],nums[1])
	for idx := 2; idx < len(nums);idx++ {
		per, cur = cur, max(nums[idx] + per, cur)
	}
	return cur
}

func robII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(searchMax(nums[1:]), searchMax(nums[:len(nums)-1]))
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func searchMax(nums []int) int {
	per := nums[0]
	cur := max(nums[0],nums[1])
	for idx := 2; idx < len(nums);idx++ {
		per, cur = cur, max(nums[idx] + per, cur)
	}
	return cur
}

func robIII(root *TreeNode) int {
	record := map[*TreeNode][2]int{}
	var sum func(node *TreeNode) (int,int)
	sum = func(node *TreeNode) (int,int) {
		if node == nil {
			return 0,0
		}
		if d, ok := record[node]; ok {
			return d[0], d[1]
		}
		ls, ln := sum(node.Left)
		rs, rn := sum(node.Right)
		d := [2]int{node.Val + ln + rn, max(ls, ln) + max(rs, rn)}
		record[node] = d
		return d[0], d[1]
	}
	return max(sum(root))
}

