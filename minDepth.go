package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

type BinaryTree struct {
	Root *TreeNode
}

func (t *BinaryTree) Insert(data int) *BinaryTree {
	if t.Root == nil {
		t.Root = &TreeNode{Value: data}
	} else {
		t.Root.insert(data)
	}
	return t
}

func (n *TreeNode) insert(data int) {
	if n == nil {
		return
	}
	if data <= n.Value {
		if n.Left == nil {
			n.Left = &TreeNode{Value: data}
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Value: data}
		} else {
			n.Right.insert(data)
		}
	}
}

func minDepth() {
	tree := &BinaryTree{}
	tree.Insert(100).
		Insert(-20).
		Insert(-50).
		Insert(-15).
		Insert(-60)

	printTree(tree.Root, 0)
}

func printTree(root *TreeNode, space int) {
	if root == nil {
		return
	}

	space += 5

	printTree(root.Right, space)

	fmt.Printf("\n")
	for i := 5; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(root.Value)

	printTree(root.Left, space)
}
