package main

import "fmt"

type BSTreeNode struct {
	Value int
	Left  *BSTreeNode
	Right *BSTreeNode
}

func TreeToLinkedList(root *BSTreeNode) *BSTreeNode {
	var head, tail *BSTreeNode
	helper(&head, &tail, root)
	return head
}

// 注意对指针的操作（取地址符的使用）
func helper(head, tail **BSTreeNode, node *BSTreeNode) {
	var lt, rh *BSTreeNode
	if node == nil {
		*head, *tail = nil, nil
		return
	}
	helper(head, &lt, node.Left)
	helper(&rh, tail, node.Right)
	if lt != nil {
		lt.Right = node
		node.Left = lt
	} else {
		*head = node
	}

	if rh != nil {
		rh.Left = node
		node.Right = rh
	} else {
		*tail = node
	}
}

func RunTest() {
	// create a BSTree
	left := &BSTreeNode{1, nil, nil}
	right := &BSTreeNode{3, nil, nil}
	root := &BSTreeNode{2, left, right}

	head := TreeToLinkedList(root)

	for p := head; p != nil; p = p.Right {
		fmt.Printf("%d ", p.Value)
	}
}

func main() {
	RunTest()
}
