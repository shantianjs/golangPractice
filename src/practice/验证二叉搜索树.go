package practice

import "math"

func isValidBST(root *TreeNode) bool {

	return validBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func validBSTHelper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	//左右子树更新不同值
	return validBSTHelper(node.Left, min, node.Val) && validBSTHelper(node.Right, node.Val, max)
}
