package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	head4 := &ListNode{
		Val:  4,
		Next: head5,
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
	reverseList(head)
}

var res *ListNode

func reverseList(head *ListNode) *ListNode {
	res = nil
	find(head, nil)
	return res
}

/**
  1->2->3->4->5-nil
              假设此时找到了5的下一个节点为空
              此时，返回的是former，即4
           在4这一轮递归中，接收到的former就是4
           这个4就是本节点node，还是有必要上一轮返回Former
                但不应该返回Former，应该返回res.next
           因为要更新res的这个链表。
           那么此时，former是4，former.next == nil
           我们要更新的就是将former.next赋值为3，即此一轮递归中的former
           然后，再把3返回。
 */
func find(node *ListNode, former *ListNode) *ListNode {
	if node == nil {
		return former // 这句代码其实没啥作用，就只有判断整个链表的头节点为不为空，而此时并不需要返回任何
		              // 数据，所以随便写都行。因为，如果不是头节点，那么他为不为空已经在node.Next == nil
		              // 那里面判断完了。
	}

	var next *ListNode
	if node.Next == nil {
		// 此node为最终节点。
		res = node
		res.Next = former
		return res.Next // 要把former返回，因为former还需要更新
	} else {
		// 前两种情况都不是，就要往下递归
		next = find(node.Next, node)
		// 每一轮递归完回来，要更新下一个节点
		next.Next = former
		return next.Next
	}
}