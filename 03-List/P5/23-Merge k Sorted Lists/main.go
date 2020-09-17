package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var res *ListNode
	res = merge(lists, res)
	return res
}

func merge(lists []*ListNode, res *ListNode) *ListNode {
	var values = 0
	var first = true // 此bool值用于确认是否到了第一个不为空的list
	// 若遍历完了之后，first仍然为true，则表明，整个
	// lists中的所有list都被遍历完了，可以返回了。
	var index = -1
	// 首先找到最小值元素所在list以及其索引位置
	for i, list := range lists {
		if list != nil && first {
			values = list.Val
			index = i
			first = false
		}

		if list != nil {
			if list.Val < values {
				values = list.Val
				index = i
			}
		}
	}

	// 若first为true，直接返回，不再进行递归。
	if first {
		return res
	}
	var list = lists[index]
	// 先将list写入res，再更新lists。
	res = list

	// 找到之后，需要更新整个lists，
	// 更新的部分只有：最小值所在的list
	fmt.Printf("看一看list位置对不对，list: %v \n", list)
	list = list.Next
	fmt.Printf("看一看list位置对不对，list: %v \n", list)
	lists[index] = list
	fmt.Printf("看一看list位置对不对，lists[index]: %v \n", list)

	// 更新完之后，再进行递归
	res.Next = merge(lists, res.Next)
	return res
}