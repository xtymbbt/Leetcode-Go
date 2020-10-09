package notMain

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func notMain() {
	head2 := &ListNode{
		Val:  2,
		Next: nil,
	}

	head := &ListNode{
		Val:  1,
		Next: head2,
	}
	res := swapPairs(head)
	fmt.Println(res, res.Next, res.Next.Next)
}
/**
  1.val        2.val       3.val        4.val
  1.next : 2   2.next : 3  3.next : 4   4.next : nil
   实际的进行节点交换，意思是取地址交换？
swap:
  2.val        1.val
  2.next : 1   1.next : 4
 */
func swapPairs(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	//var mid = head
	//var back = head.Next.Next
	//head = head.Next
	//head.Next = mid
	//head.Next.Next = back
	//fmt.Println(head, head.Next, head.Next.Next)


	return 	swap(head)
}

func swap(node *ListNode) *ListNode{
	if node == nil {
		return nil
	}

	if node.Next == nil {
		return nil
	}

	var mid = node
	var back = node.Next.Next
	node = node.Next
	node.Next = mid
	node.Next.Next = back
	fmt.Println(node, node.Next, node.Next.Next)
	if node.Next.Next == nil {
		return node
	}

	if node.Next.Next.Next == nil {
		return node
	}

	node.Next.Next = node.Next.Next.Next
	swap(back)
	return node
}