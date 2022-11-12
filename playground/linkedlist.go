package main

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	head   *Node
	length int64
}

func NewLinkedList() LinkedList {
	var h Node = NewNode(0)
	var ll LinkedList = LinkedList{head: &h, length: 0}
	return ll
}

func (l *LinkedList) Length() int64 {
	return l.length
}

func (l *LinkedList) Append(value int) {
	newNode := NewNode(value)
	l.Tail().SetNext(&newNode)
	l.length++
}

func (l *LinkedList) Remove(index int) {
	if int64(index) >= l.Length() {
		/* Exception */
		return
	}
	var cur *Node = l.Head()
	for i := 0; i < index; i++ {
		cur = cur.Next()
	}

	newNextNode := cur.Next().Next()
	cur.SetNext(newNextNode)

	l.length--
}

func (l *LinkedList) Tail() *Node {
	cur := l.Head()
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) Traverse() []int {
	cur := l.Head()
	var result []int = []int{}
	for cur.Next() != nil {
		cur = cur.Next()
		result = append(result, cur.Value())
	}
	return result
}

func (l *LinkedList) Format() string {
	var elements []int = l.Traverse()
	var s []string
	for i := range elements {
		s = append(s, fmt.Sprint(elements[i]))
	}
	return "[" + strings.Join(s, ",") + "]"
}

func (l *LinkedList) reduce(method func(int, int) int) int {
	ag := 0
	cur := l.Head()
	for cur.Next() != nil {
		cur = cur.Next()
		ag = method(ag, cur.Value())
	}
	return ag
}
