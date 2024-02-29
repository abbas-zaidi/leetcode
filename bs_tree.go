package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func bs_tree() {

	t := TreeNode{
		Value: 33,
		Left:  nil,
		Right: nil,
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

func (t *TreeNode) Insert(val int) error {

	if t == nil {
		return errors.New("Tree is nil")
	}

	if t.Value == val {
		return errors.New("This node already exists")
	}

	if t.Value > val {
		if t.Left == nil {
			t.Left = &TreeNode{Value: val}
		} else {
			t.Left.Insert(val)
		}
	}

	if t.Value < val {
		if t.Right == nil {
			t.Right = &TreeNode{Value: val}
		} else {
			t.Right.Insert(val)
		}
	}

	return nil
}

func (t *TreeNode) FindMin() int {
	if t.Left == nil {
		return t.Value
	} else {
		return t.Left.FindMin()
	}
}
