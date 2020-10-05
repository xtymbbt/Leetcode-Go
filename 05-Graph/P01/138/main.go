package main

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func main() {
	
}
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

var visited map[*Node]*Node

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	visited = make(map[*Node]*Node)
	return clone(head)
}

func clone(node *Node) *Node {
	if node == nil {
		return node
	}
	if v, ok := visited[node]; ok {
		return v
	}
	newNode := &Node{
		Val:    node.Val,
		Next:   nil,
		Random: nil,
	}
	visited[node] = newNode
	newNode.Next = clone(node.Next)
	newNode.Random = clone(node.Random)
	return newNode
}