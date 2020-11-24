package practice

func rightSideViewDFS(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		//每层节点个数
		l := len(stack)
		for l > 0 {
			l--
			//这一层最后一个
			if l == 0 {
				res = append(res, stack[0].Val)
			}
			if stack[0].Left != nil {
				stack = append(stack, stack[0].Left)
			}
			if stack[0].Right != nil {
				stack = append(stack, stack[0].Right)
			}
			//删除首节点
			stack = stack[1:]
		}
	}
	return res
}


func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	recorder := map[int]int{}

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if _, ok := recorder[level];!ok {
			recorder[level] = node.Val
		}
		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}
	dfs(root,0)
	res := make([]int, len(recorder), len(recorder))
	for i := 0; i < len(recorder); i++ {
		res[i] = recorder[i]
	}
	return res
}
