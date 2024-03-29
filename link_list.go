package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data int64
	Next *Node
}

type Linklist struct {
	Head *Node
	Size int16
}

func (f *Linklist) addNodeEnd(element int64) {

	newnode := Node{element, nil}

	if f.Head == nil {
		f.Size++
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
	f.Size++
	prev.Next = &newnode
	return
}

func (f *Linklist) deleteElementAt(position int16) (bool, error) {

	if position < 0 {
		return false, errors.New(" Invalid Position, position must be a positive value")
	}
	if position >= f.Size {
		return false, errors.New("Out of range")
	}

	if f.Head == nil {
		return false, nil
	}
	if position == 0 {
		f.Size--
		f.Head = f.Head.Next
		return true, nil
	}
	var prev, current *Node = nil, f.Head
	i := int16(0)
	for i < position {
		prev = current
		current = current.Next
		i++
	}
	f.Size--
	prev.Next = current.Next
	return true, nil
}

/**

	In this code we make a approch that is number of element are even then throw error
	you can decide by our way what to return

**/

func (f Linklist) middle() (*Node, error) {

	if f.Head == nil {
		return nil, errors.New("Empty List")
	}

	var fast, slow *Node = f.Head, f.Head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast.Next != nil {
		return nil, errors.New(" This is a even element list there is no middle element")
	}
	return slow, nil
}

func (f Linklist) printList() {
	current := f.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func main() {
	var ll Linklist
	ll.addNodeEnd(10)
	ll.addNodeEnd(20)
	ll.addNodeEnd(30)
	ll.addNodeEnd(30)
	ll.printList()

	/** For Delete Element
	_, err := ll.deleteElementAt(0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ll.Size)
	ll.printList()
	**/

	midNode, err := ll.middle()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(midNode.Data)
	}

}
