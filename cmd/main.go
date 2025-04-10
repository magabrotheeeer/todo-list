package main

import (
	"todo-list/internal/linkedlist"
	"fmt"
)

func main() {
	x := linkedlist.NewLinkedList("asd")
	x.AddNewNode()
	fmt.Println(x)	
}