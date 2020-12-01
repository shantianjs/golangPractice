package practice

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	workQue := map[*TreeNode]bool{root:true}
	step := 0
	for len(workQue) != 0 {
		step++
		tem := map[*TreeNode]bool{}
		//遍历当前层节点，是否返回
		for node := range workQue {
			if node.Left == nil && node.Right == nil {
				return step
			}
			if node.Left != nil {
				tem[node.Left] = true
			}
			if node.Right != nil {
				tem[node.Right] = true
			}
		}
		workQue = tem
	}
	return step
}
