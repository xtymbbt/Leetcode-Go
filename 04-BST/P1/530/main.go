package _30

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Value interface {}

var min int
func getMinimumDifference(root *TreeNode) int {
	min = math.MaxInt64
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}

	search(root.Left, nil, root.Val)
	search(root.Right, root.Val, nil)

	return min
}

func search(root *TreeNode, lower Value, upper Value) {
	if root == nil {
		return
	}
	if lower != nil {
		x := root.Val-lower.(int)
		if x < 0 {x = -x}
		if x < min {min = x}
	}
	if upper != nil {
		x := root.Val-upper.(int)
		if x < 0 {x = -x}
		if x < min {min = x}
	}
	search(root.Left, lower, root.Val)
	search(root.Right, root.Val, upper)
}
