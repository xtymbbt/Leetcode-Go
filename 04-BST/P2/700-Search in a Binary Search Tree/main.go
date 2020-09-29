package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res *TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	res = nil
	if root == nil {
		return nil
	}
	if val < root.Val {
		// find left
		search(root.Left, val)
		return res
	}
	if val == root.Val {
		return root
	}
	// find right
	search(root.Right, val)
	return res
}

func search(root *TreeNode, val int) {
	if root == nil {
		return
	}
	if val < root.Val {
		// find left
		search(root.Left, val)
	}
	if val == root.Val {
		res = root
		return
	}
	// find right
	search(root.Right, val)
}