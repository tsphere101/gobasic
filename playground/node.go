package main

type Node struct {
	next  *Node
	value int
}

func NewNode(value int) Node {
	return Node{next: nil, value: value}
}

func (node *Node) SetNext(next *Node) {
	node.next = next
}

func (node *Node) SetValue(value int) {
	node.value = value
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) Value() int {
	return node.value
}
