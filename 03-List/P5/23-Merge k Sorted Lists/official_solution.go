package notnotMain

func notMain() {
	
}

func mergeKListsOFFICIAL(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	// 思路：将之分为左右两边，形同分治法。
	// 左边和右边都是已经处理成最终list结果的链表。
	// 然后使用21题中的合并两个链表的简单算法即可。
	return mergeLists(mergeKListsOFFICIAL(lists[:len(lists)/2]), mergeKListsOFFICIAL(lists[len(lists)/2:]))
}

func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	newHead := &ListNode{}
	p := newHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
			p = p.Next
		} else {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
		}
	}
	for l1 != nil {
		p.Next = l1
		l1 = l1.Next
		p = p.Next
	}
	for l2 != nil {
		p.Next = l2
		l2 = l2.Next
		p = p.Next
	}
	return newHead.Next
}