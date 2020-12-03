package practice

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{val,nil, nil}
		} else {
			insertIntoBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{val,nil, nil}
		} else {
			insertIntoBST(root.Right,val)
		}
	}
	return root
}
