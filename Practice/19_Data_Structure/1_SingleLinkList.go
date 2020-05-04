package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkList interface {
	addNodeLast()
	addNodeFirst()
	addNodeMiddle()
	deleteNode()
	printList()
}

type nodeList struct {
	head *node
	tail *node
}

func (li *nodeList) addNodeLast(data int) {
	newNode := new(node)
	newNode.data = data
	newNode.next = nil
	if li.head == nil {
		li.head = newNode
		li.tail = newNode
	} else {
		li.tail.next = newNode
		li.tail = newNode
	}
}

func (li *nodeList) printList() {
	for i := li.head; i != nil; {
		fmt.Println("Node ", i.data)
		i = i.next
	}
}

func myList() {
	li := &nodeList{head: nil, tail: nil}
	li.addNodeLast(1)
	li.addNodeLast(2)
	li.addNodeLast(3)
	li.addNodeLast(4)
	li.addNodeLast(5)
	li.printList()
}

func main() {
	myList()
}
