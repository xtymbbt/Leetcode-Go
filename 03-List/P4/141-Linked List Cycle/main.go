package notMain

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func notMain() {
	head4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	head3 := &ListNode{
		Val:  3,
		Next: head4,
	}
	head2 := &ListNode{
		Val:  2,
		Next: head3,
	}
	head := &ListNode{
		Val:  1,
		Next: head2,
	}
	head4.Next = head2
	fmt.Println("final result: ", hasCycle(head))
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var hashMap map[*ListNode]int
var res bool

func hasCycle(head *ListNode) bool {
	hashMap = make(map[*ListNode]int)
	res = false
	find(head)
	return res
}

func find(node *ListNode) {
	if node == nil {
		return
	}
	hashMap[node] = 1
	_, ok := hashMap[node.Next]
	if ok {
		res = true
		return
	}
	if node.Next == nil {
		return
	}
	find(node.Next)
}