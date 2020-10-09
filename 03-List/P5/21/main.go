package notMain

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func notMain() {
	l23 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l22 := &ListNode{
		Val:  3,
		Next: l23,
	}
	l2 := &ListNode{
		Val:  1,
		Next: l22,
	}


	l13 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l12 := &ListNode{
		Val:  2,
		Next: l13,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l12,
	}
	fmt.Println("final result: ", mergeTwoLists(l1, l2))
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var res *ListNode

	if l1 == nil && l2 == nil {
		return res
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			res = l2
			res.Next = merge(l1, l2.Next, res.Next)
		} else {
			res = l1
			res.Next = merge(l1.Next, l2, res.Next)
		}
		return res
	} else {
		fmt.Println("出现了其他情况")
		return nil
	}
}

func merge(l1 *ListNode, l2 *ListNode, res *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return res
	} else if l1 == nil && l2 != nil {
		res = l2
		return res
	} else if l1 != nil && l2 == nil {
		res = l1
		return res
	} else if l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			res = l2
			res.Next = merge(l1, l2.Next, res.Next)
		} else {
			res = l1
			res.Next = merge(l1.Next, l2, res.Next)
		}
		return res
	} else {
		fmt.Println("出现了其他情况")
		return res
	}
}