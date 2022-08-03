package main

import "fmt"

type Node struct {
	Value               int
	left, right, parent *Node
}

func NewNode(value int, left *Node, right *Node) *Node {
	node := &Node{Value: value, left: left, right: right}
	left.parent = node
	right.parent = node
	return node
}

func NewTerminalNode(value int) *Node {
	return &Node{Value: value}
}

type InOrderIterator struct {
	Current, root *Node
	returnedStart bool
}

func (i *InOrderIterator) Reset() {
	i.Current = i.root
	i.returnedStart = false
}

func (i *InOrderIterator) MoveNext() bool {
	if i.Current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true
	}

	if i.Current.right != nil {
		i.Current = i.Current.right
		for i.Current.left != nil {
			i.Current = i.Current.left
		}
		return true
	} else {
		par := i.Current.parent
		for par != nil && i.Current == par.right {
			i.Current = par
			par = par.parent
		}
		i.Current = par
		return i.Current != nil
	}
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	iterator := &InOrderIterator{
		Current:       root,
		root:          root,
		returnedStart: false,
	}
	for iterator.Current.left != nil {
		iterator.Current = iterator.Current.left
	}
	return iterator
}

type BinaryTree struct {
	root *Node
}

func (b *BinaryTree) InOrderIterator() *InOrderIterator {
	return NewInOrderIterator(b.root)
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root: root}
}

func main() {
	//   1
	//  / \
	// 2   3

	// in-order: 213
	// preorder: 123
	// postorder: 231

	root := NewNode(1, NewTerminalNode(2), NewTerminalNode(3))
	tree := NewBinaryTree(root)
	iter := tree.InOrderIterator()
	for iter.MoveNext() {
		fmt.Print(iter.Current.Value)
	}
}
