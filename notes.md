# Leetcode Notes
### 445. 两数相加 II

给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

#### 进阶：
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

#### 示例：
输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

#### 解决思路1：反转链表
```go
package main
type ListNode struct {
	Val int
	Next *ListNode
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
```
#### 解决思路2：利用栈来“反转”链表
```c++
// Author: Huahua
// Running time: 40 ms
class Solution {
public:
  ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    stack<int> s1;
    stack<int> s2;
    while (l1) {
      s1.push(l1->val);
      l1 = l1->next;
    }    
    while (l2) {
      s2.push(l2->val);
      l2 = l2->next;
    }    
    ListNode* head = nullptr;
    int sum = 0;
    while (!s1.empty() || !s2.empty() || sum) {
      sum += s1.empty() ? 0 : s1.top();
      sum += s2.empty() ? 0 : s2.top();
      if (!s1.empty()) s1.pop();
      if (!s2.empty()) s2.pop();            
      ListNode* n = new ListNode(sum % 10);
      sum /= 10;
      n->next = head;
      head = n;      
    }    
    return head;      
  }
};
```

### 34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
#### 进阶：
你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
#### 示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
#### 示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
#### 示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
#### 解决思路：双指针+二分搜索
```go
package _4

func searchRange(nums []int, target int) []int {
	return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := l + (r - l) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}

	if l == len(nums) || nums[l] != target {
		return -1
	}
	return l
}

func findRight(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := l + (r - l) / 2
		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	// l points to the first element this is greater than target.
	l--
	if l < 0 || nums[l] != target {
		return -1
	}
	return l
}

```