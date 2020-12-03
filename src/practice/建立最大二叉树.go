package practice

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := findMaxIndex(nums)
	var right *TreeNode
	var left *TreeNode
	if idx < len(nums) -1 {
		right = constructMaximumBinaryTree(nums[idx+1:])
	}
	if idx > 0 {
		left = constructMaximumBinaryTree(nums[:idx])
	}
	return &TreeNode{nums[idx], left,right}
}

func findMaxIndex(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[res] {
			res = i
		}
	}
	return res
}
