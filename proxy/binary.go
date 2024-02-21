package main

import (
	"fmt"
)

type NodeBinary struct {
	Key    int
	Height int
	Left   *NodeBinary
	Right  *NodeBinary
}

type AVLTree struct {
	Root *NodeBinary
}

func NewNode(key int) *NodeBinary {
	return &NodeBinary{Key: key, Height: 1}
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func (t *AVLTree) ToMermaid() string {
	return traverseGraph(t.Root)
}

func height(node *NodeBinary) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func updateHeight(node *NodeBinary) {
	node.Height = 1 + max(height(node.Left), height(node.Right))
}

func getBalance(node *NodeBinary) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func leftRotate(x *NodeBinary) *NodeBinary {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	updateHeight(x)
	updateHeight(y)

	return y
}

func rightRotate(y *NodeBinary) *NodeBinary {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	updateHeight(y)
	updateHeight(x)

	return x
}

func insert(node *NodeBinary, key int) *NodeBinary {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	} else {
		return node // Duplicate keys are not allowed
	}

	updateHeight(node)

	balance := getBalance(node)

	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}

	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}

	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

func traverseGraph(node *NodeBinary) string {
	if node == nil {
		return ""
	}

	result := fmt.Sprintf("%d\n", node.Key)

	if node.Left != nil {
		result += fmt.Sprintf("%d --> %d\n", node.Key, node.Left.Key)
		result += traverseGraph(node.Left)
	}

	if node.Right != nil {
		result += fmt.Sprintf("%d --> %d\n", node.Key, node.Right.Key)
		result += traverseGraph(node.Right)
	}

	return result
}

func GenerateTree(count int) *AVLTree {
	tree := &AVLTree{}
	for i := 0; i < count; i++ {
		tree.Insert(i)
	}
	return tree
}
