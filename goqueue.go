// Implementation queue with linkedlist
package goqueue

import "fmt"

// Linkedlist node (prev, next, data)
type Node struct {
	prev *Node
	next *Node
	data interface{}
}

// Linkedlist (head, tail)
type list struct {
	head *Node
	tail *Node
}

// Inserts the specified element at the end of this queue
func (l *list) Enqueue(data interface{}) {
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

// Remove and retrieves the first element of this deque
func (l *list) Dequeue() interface{} {
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

// Get all element in Queue
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

// Retrieves, but does not remove, the first element of this deque
func (l *list) Peek() interface{} {
	// Look first data
	return l.head.data
}

// Create new node
func NewNode(data interface{}) *Node {
	return &Node{data: data}
}

// Interface Queue
type Queue interface {
	Enqueue(data interface{}) //add last
	Dequeue() interface{}     //remove first
	String() string
	Peek() interface{} //lihat data paling awal
}

// Initialize new queue
func NewQueue() Queue {
	l := list{
		head: nil,
		tail: nil,
	}

	return &l
}
