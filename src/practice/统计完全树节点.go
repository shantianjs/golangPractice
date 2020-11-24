package practice

func countNodes(root *TreeNode) (count int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		count++
		dfs(node.Right)
		dfs(node.Left)
	}
	dfs(root)
	return count
}
