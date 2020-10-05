package main

type Node struct {
	Val int
	Neighbors []*Node
}

func main() {

}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node,  visited map[*Node]*Node) *Node{
	if node == nil{
		return node
	}

	if v, ok := visited[node]; ok{
		return v
	}

	new_node := &Node{
		Val: node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	visited[node] = new_node

	for i := 0; i < len(node.Neighbors); i++{
		new_node.Neighbors[i] = clone(node.Neighbors[i], visited)
	}

	return new_node
}