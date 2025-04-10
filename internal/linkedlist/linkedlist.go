package linkedlist

type Node struct {
	content_ string
	next_    *Node
	isEnd_   bool
}

type LinkedList struct {
	head_ *Node
	size_ int
	tail_ *Node
}

func NewLinkedList(content string) *LinkedList {
	initNode := &Node{
		content_: content,
		next_:    nil,
		isEnd_:   false,
	}
	resLinkedList := &LinkedList{
		head_: initNode,
		size_: 1,
		tail_: initNode,
	}
	return resLinkedList
}
func NewNode() *Node {
		
}

func (list *LinkedList) AddNewNode(node *Node) {
	list.tail_ = node
	list.size_++
}
