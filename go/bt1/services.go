package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, value int) *Node {

	if root == nil {
		return &Node{value: value}
	}

	if value < root.value {
		root.left = insert(root.left, value)
	} else {
		root.right = insert(root.right, value)
	}

	return root
}
func printDESC(root *Node) {

	if root == nil {
		return
	}

	printDESC(root.right)
	fmt.Print(root.value, " ")
	printDESC(root.left)
}

func buildTree(values []int) *Node {

	var root *Node

	for _, v := range values {
		root = insert(root, v)
	}

	return root
}