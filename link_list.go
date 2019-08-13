package main

import "fmt"

type Node struct {
	Data int64
	Next *Node
}

type Linklisst struct {
	Head *Node
	Size int16
}

func (f *Linklisst) addNodeEnd(element int64) {

	newnode := Node{element, nil}

	if f.Head == nil {
		f.Head = &newnode
		return
	}

	var prev, current *Node
	prev = nil
	current = f.Head
	for current != nil {
		prev = current
		current = current.Next
	}
	prev.Next = &newnode
	return
}

func (f Linklisst) printList() {
	current := f.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func main() {
	var ll Linklisst
	ll.addNodeEnd(10)
	ll.addNodeEnd(20)
	ll.addNodeEnd(30)
	ll.printList()
}
