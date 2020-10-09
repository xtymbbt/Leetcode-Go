package notMain

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func notMain() {
	h4 := &ListNode{
		Val:  3,
		Next: nil,
	}
	h3 := &ListNode{
		Val:  1,
		Next: h4,
	}
	h2 := &ListNode{
		Val:  2,
		Next: h3,
	}
	h1 := &ListNode{
		Val:  4,
		Next: h2,
	}
	fmt.Println(sortList(h1))
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 取开头的元素作为基准元素。
	baseEle := head
	head = head.Next
	var left *ListNode
	var right *ListNode
	for head != nil {
		next := head.Next
		if baseEle.Val >= head.Val {
			left = addList(left, head)
		} else {
			right = addList(right, head)
		}
		head = next
	}
	left = sortList(left)
	right = sortList(right)
	baseEle.Next = right
	if left == nil {
		return baseEle
	}
	backup := left
	for left != nil {
		if left.Next == nil {
			left.Next = baseEle
			break
		}
		left = left.Next
	}
	return backup
}

func addList(list *ListNode, nodeToAdd *ListNode) *ListNode {
	nodeToAdd.Next = list
	return nodeToAdd
}