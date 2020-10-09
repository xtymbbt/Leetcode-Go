package notMain

func notMain() {
	
}
func insertionSortList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	// 未排序数组
	cur := head.Next
	head.Next = nil
	// 排序数组尾
	end := head
	for cur != nil {
		pre := dummy
		next := cur.Next
		// 优化部分： 如果比排序链表的最后一个大，直接插到末尾
		if end.Val < cur.Val {
			cur.Next = nil
			end.Next = cur
			end = cur
			cur = next
			continue
		}
		// 寻找插入点
		for pre.Next != nil && pre.Next.Val <= cur.Val {
			pre = pre.Next
		}
		// 插入
		cur.Next = pre.Next
		pre.Next = cur
		cur = next
	}
	return dummy.Next
}