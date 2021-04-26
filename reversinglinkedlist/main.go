package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
}

func (L *List) insertAtBegining(s int) {
	n := &Node{value: s}

	if L.head != nil {
		n.next = L.head
		L.head = n
	} else {
		L.head = n
	}
}

func (L *List) show() {
	currentNode := L.head
	if currentNode == nil {
		fmt.Println("List is empty")
	} else {
		for currentNode != nil {
			fmt.Printf("value: %v, next: %v \n", currentNode.value, currentNode.next)
			currentNode = currentNode.next
		}
	}
}

func (l *List) reverse() {
	var previous, next *Node
	current := l.head

	for current != nil {
		next = current.next
		current.next = previous

		previous = current
		current = next
	}
	l.head = previous
}

func main() {
	l := List{}
	l.insertAtBegining(10)
	l.insertAtBegining(2)
	l.insertAtBegining(6)
	l.insertAtBegining(1)
	l.show()
	l.reverse()
	l.show()

}
