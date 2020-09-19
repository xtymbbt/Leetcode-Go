package main

func main() {

}
func sortList2(head *ListNode) *ListNode {

	// 0 or 1 element.
	if head == nil || head.Next == nil {
		return head
	}
	// 2 pointers, if the fast point comes to the end, the slow one must be in the middle.
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// split into 2 parts.
	n := slow.Next
	slow.Next = nil
	// Sort recursivley
	return merge(sortList2(head), sortList2(n))

}

func merge(node1 *ListNode, node2 *ListNode) *ListNode {
	// Create a new empty list.
	node := &ListNode{Val : 0 }
	current := node
	// Compare one by one, put the smaller value into the new list.
	for node1 != nil && node2 != nil {
		if node1.Val <= node2.Val {
			current.Next = node1
			node1 = node1.Next
		} else {
			current.Next = node2
			node2 = node2.Next
		}
		current = current.Next
	}
	if node1 != nil {
		current.Next = node1
		node1 = node1.Next
	}
	if node2 != nil {
		current.Next = node2
		node2 = node2.Next
	}
	return node.Next
}
