//节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	low, fast := head, head.Next
	vHead := ListNode{0, head}

	for fast != nil {
		//摘出
		low.Next = fast.Next
		//插入
		tem := &vHead
		for tem != low {
			if tem.Next.Val > fast.Val {
				fast.Next = tem.Next
				tem.Next = fast
				break
			}
			tem = tem.Next
		}
		if tem == low {
			//最后向前移动
			low.Next = fast
			low = fast
			fast = fast.Next
		} else {
			fast = low.Next
		}
	}
	return vHead.Next
}

// var head *ListNode
// nums := []int{4, 2, 1, 3}
// idx := len(nums) - 1
// for idx >= 0 {
// 	head = &ListNode{nums[idx], head}
// 	idx--
// }
// insertionSortList(head)