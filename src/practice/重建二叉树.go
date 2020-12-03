package practice

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	cur := preorder[0]
	idx := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == cur {
			idx = i
			break
		}
	}
	//按idx个数，来区分左右树的先序遍历个数
	return &TreeNode{cur, buildTree(preorder[1:idx+1], inorder[:idx]), buildTree(preorder[1 + idx:], inorder[idx+1:])}
}
