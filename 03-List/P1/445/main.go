package notMain

type ListNode struct {
	Val int
	Next *ListNode
}

func notMain() {
	
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 先反转
	l1R := &ListNode{
		Val:  l1.Val,
		Next: nil,
	}
	l2R := &ListNode{
		Val:  l2.Val,
		Next: nil,
	}
	rootR1 := reverse(l1.Next, l1R)
	rootR2 := reverse(l2.Next, l2R)
	// 反转完了用之前的算法计算出值
	res := addTwoNumbers2(rootR1, rootR2)
	// 再反转
	resR := &ListNode{
		Val:  res.Val,
		Next: nil,
	}
	return reverse(res, resR)
}

func reverse(l *ListNode, lR *ListNode) *ListNode{
	if l != nil {
		var lR2 = &ListNode{
			Val:  l.Val,
			Next: lR,
		}
		return reverse(l.Next, lR2)
	} else {
		return lR
	}
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{
		Val:  0,
		Next: nil,
	}
	sum(l1, l2, false, res)
	return res
}

func sum(l1 *ListNode, l2 *ListNode, tensUP bool, res *ListNode) {
	if l1 == nil {
		l1 = &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	if l2 == nil {
		l2 = &ListNode{
			Val:  0,
			Next: nil,
		}
	}
	if tensUP {
		res.Val = l1.Val + l2.Val + 1
	} else {
		res.Val = l1.Val + l2.Val
	}
	if res.Val >= 10 {
		tensUP = true
		res.Val -= 10
		res.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	} else {
		tensUP = false
	}
	if l1.Next != nil || l2.Next != nil{
		res.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		sum(l1.Next, l2.Next, tensUP, res.Next)
	}
}