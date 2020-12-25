package _35

import "fmt"

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val   int
*     Left  *TreeNode
*     Right *TreeNode
* }
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP, pathQ := findPath(root, p), findPath(root, q)
	return comparePath(pathP, pathQ)
}

func findPath(root, node *TreeNode) []*TreeNode {
	res := []*TreeNode{root}
	if root == node {
		return res
	}
	resCp := []*TreeNode{root}
	res = dfs(root.Left, node, res)
	if res == nil {
		res = dfs(root.Right, node, resCp)
	}
	return res
}

func dfs(root, node *TreeNode, res []*TreeNode) []*TreeNode {
	if root == nil {
		res = nil
		return res
	}

	this := make([]*TreeNode, len(res))
	copy(this, res)
	this = append(this, root)
	if root == node {
		return this
	}
	thisCp := make([]*TreeNode, len(this))
	copy(thisCp, this)
	this = dfs(root.Left, node, this)
	if this == nil {
		this = dfs(root.Right, node, thisCp)
	}
	return this
}

func comparePath(pathP []*TreeNode, pathQ []*TreeNode) *TreeNode {
	i := 0
	p, q := len(pathP), len(pathQ)
	fmt.Println(pathQ[i])
	fmt.Println(pathP[i])
	for i < p && i < q && pathP[i] == pathQ[i] {
		i++
	}
	return pathP[i-1]
}