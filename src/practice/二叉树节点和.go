package practice

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return calSum(root,0)
}

func calSum(node *TreeNode, sum int) int {
	newSum := 10 * sum + node.Val
	if node.Left == nil && node.Right == nil{
		return newSum
	} else if node.Left == nil {
		return calSum(node.Right, newSum)
	} else if node.Right == nil {
		return calSum(node.Left, newSum)
	}
	return calSum(node.Left,newSum) + calSum(node.Right, newSum)
}
