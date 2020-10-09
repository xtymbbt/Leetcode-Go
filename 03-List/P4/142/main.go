package notMain

type ListNode struct {
	Val int
	Next *ListNode
}

func notMain() {
	
}

var hashMap map[*ListNode]int

func detectCycle(head *ListNode) *ListNode {
	hashMap = make(map[*ListNode]int)

	return find(head)
}

func find(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	hashMap[node] = 1
	_, ok := hashMap[node.Next]
	if ok {
		return node.Next
	}
	if node.Next == nil {
		return nil
	}
	return find(node.Next)
}