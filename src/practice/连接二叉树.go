package practice

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	workSlice := []*Node{root}
	for len(workSlice) > 0 {
		nextWork := []*Node{}
		for i := 0;i<len(workSlice);i++ {
			if i < len(workSlice) - 1 {
				workSlice[i].Next = workSlice[i+1]
			}
			if workSlice[i].Left != nil{
				nextWork = append(nextWork, workSlice[i].Left)
				nextWork = append(nextWork, workSlice[i].Right)
			}
		}
		workSlice = nextWork
	}
	return root
}
