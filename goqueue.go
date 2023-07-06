// Implementation queue with linkedlist
package goqueue

import "fmt"

// Interface Queue
type Queue interface {
	Enqueue(data interface{})
	Dequeue() interface{}
	String() string
	Peek() interface{}
}

// Create new node
func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

// Initialize new queue
func NewQueue() Queue {
	l := List{
		Head: nil,
		Tail: nil,
	}

	return &l
}

// Linkedlist node (Prev, Next, Data)
type Node struct {
	Prev *Node
	Next *Node
	Data interface{}
}

// Linkedlist (Head, Tail)
type List struct {
	Head *Node
	Tail *Node
}

// Inserts the specified element at the end of this queue
func (l *List) Enqueue(data interface{}) {
	newNode := NewNode(data)
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Prev = l.Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

// Remove and retrieves the first element of this deque
func (l *List) Dequeue() interface{} {
	if l.Head == nil {
		return nil
	} else if l.Head.Next == nil && l.Head == l.Tail { //if only 1 Data from List
		data := l.Head.Data
		l.Head.Next = nil
		l.Head = nil
		l.Tail = nil
		return data
	} else {
		data := l.Head.Data
		l.Head = l.Head.Next
		l.Head.Prev.Next = nil
		l.Head.Prev = nil
		return data
	}
}

// Get all element in Queue
func (l *List) String() string {
	// Print all Data in queue
	result := "["
	current := l.Head
	for current != nil {
		result += fmt.Sprint(current.Data)

		if current.Next != nil {
			result += ", "
		}

		current = current.Next
	}
	result += "]"
	return result
}

// Retrieves, but does not remove, the first element of this deque
func (l *List) Peek() interface{} {
	// Look first Data
	return l.Head.Data
}
