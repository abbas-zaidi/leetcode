package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {

	t := TreeNode{
		val:   33,
		left:  nil,
		right: nil,
	}
	j := 0
	for i := 0; i < 50; i++ {

		e := t.Insert(rand.Intn(100))

		if e != nil {
			fmt.Println("Error", e)
		} else {
			j++
		}
	}
	fmt.Println("Node inserted=", j)

	fmt.Println("Min == ", t.FindMin())
}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(val int) error {

	if t == nil {
		return errors.New("Tree is nil")
	}

	if t.val == val {
		return errors.New("This node already exists")
	}

	if t.val > val {
		if t.left == nil {
			t.left = &TreeNode{val: val}
		} else {
			t.left.Insert(val)
		}
	}

	if t.val < val {
		if t.right == nil {
			t.right = &TreeNode{val: val}
		} else {
			t.right.Insert(val)
		}
	}

	return nil
}

func (t *TreeNode) FindMin() int {
	if t.left == nil {
		return t.val
	} else {
		return t.left.FindMin()
	}
}
