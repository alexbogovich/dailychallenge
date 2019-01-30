package main

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func NewNode(value string, left, right *Node) *Node {
	return &Node{Value: value, Left: left, Right: right}
}

func IsUnival(node *Node) (bool, int) {
	if node == nil {
		return true, 0
	}

	if node.Left == nil && node.Right == nil {
		return true, 1
	}

	leftIsUnival, leftCount := IsUnival(node.Left)
	rightIsUnival, rightCount := IsUnival(node.Right)

	childSum := leftCount + rightCount

	if node.Left == nil || node.Right == nil {
		return false, childSum
	}

	isUnival := leftIsUnival && rightIsUnival && node.Left.Value == node.Value && node.Right.Value == node.Value
	if isUnival {
		return true, childSum + 1
	}

	return false, childSum
}
