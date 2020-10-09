package notMain

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
type ListNode struct {
	Val int
	Next *ListNode
}
func notMain() {

}

func insertionSortList(head *ListNode) *ListNode {
	// 就从前往后找呗
	if head == nil {
		return head
	}
	head = insert(head, head.Next)
	return head
}

func insert(head *ListNode, nodeToInsert *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if nodeToInsert == nil {
		return head
	}
	nextNodeToInsert := nodeToInsert.Next
	formerHead := head
	var lastHead *ListNode
	for head != nodeToInsert {
		if head.Val <= nodeToInsert.Val {
			lastHead = head
			head = head.Next
		} else {
			if lastHead != nil {
				lastHead.Next = nodeToInsert
			}
			nodeToInsertNextBackup := nodeToInsert.Next
			nodeToInsert.Next = head
			for head != nodeToInsert {
				if head.Next == nodeToInsert {
					head.Next = nodeToInsertNextBackup
					break
				}
				head = head.Next
			}
			break
		}
	}
	if lastHead == nil {
		formerHead = nodeToInsert
	}
	head = insert(formerHead, nextNodeToInsert)
	return head
}