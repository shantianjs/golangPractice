package practice

var sum = 0

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	travel(root)
	return root
}

func travel(root *TreeNode) {
	if root == nil {
		return
	}
	travel(root.Right)
	root.Val += sum
	sum = root.Val
	travel(root.Left)
}
