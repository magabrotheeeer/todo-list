package linkedlist

import (
	"time"
)

type Node struct {
	Content   string
	Next      *Node
	StartTime time.Time
	EndTime   time.Time
}

type LinkedList struct {
	Head *Node
	Size int
	Tail *Node
}

func (list *LinkedList) AddNewNode(NewContent string) {
	newNode := &Node{Content: NewContent, Next: nil}
	if list.Head == nil {
		list.Head = newNode
	} else {
		list.Tail = newNode
	}
	list.Size++
}
