package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	data interface{}
}

type list struct {
	head *Node
	tail *Node
}

func (l *list) Enqueue(data interface{}) {
	// Add last O(1)
	newNode := NewNode(data)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *list) Dequeue() interface{} {
	// Remove first O(1)
	if l.head == nil {
		return nil
	} else if l.head.next == nil && l.head == l.tail { //if only 1 data from list
		data := l.head.data
		l.head.next = nil
		l.head = nil
		l.tail = nil
		return data
	} else {
		data := l.head.data
		l.head = l.head.next
		l.head.prev.next = nil
		l.head.prev = nil
		return data
	}
}

func (l *list) String() string {
	// Print all data in queue
	result := "["
	current := l.head
	for current != nil {
		result += fmt.Sprint(current.data)

		if current.next != nil {
			result += ", "
		}

		current = current.next
	}
	result += "]"
	return result
}

func (l *list) Peek() interface{} {
	// Look first data
	return l.head.data
}

func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

type Queue interface {
	Enqueue(data interface{}) //add last
	Dequeue() interface{}     //remove first
	String() string
	Peek() interface{} //lihat data paling awal
}

func NewQueue() Queue {
	l := list{
		head: nil,
		tail: nil,
	}

	return &l
}
