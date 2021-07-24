package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Node struct
type Node struct {
	data int
	next *Node
}

// LinkedList struct
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func makeList() LinkedList {
	return LinkedList{head: nil, tail: nil, size: 0}
}

func makeNode(data int) Node {
	return Node{data: data, next: nil}
}

func (l *LinkedList) append(data int) *Node {
	newNode := makeNode(data)
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		l.size++
		return &newNode
	}
	l.tail.next = &newNode
	l.tail = &newNode
	l.size++
	return &newNode
}

func (l *LinkedList) prepend(data int) *Node {
	newNode := makeNode(data)
	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		l.size++
		return &newNode
	}

	newNode.next = l.head
	l.head = &newNode
	l.size++
	return &newNode
}

func (l *LinkedList) getSize() int {
	return l.size
}

func (l *LinkedList) print() string {
	xs := []string{}
	current := l.head

	for current != nil {
		xs = append(xs, strconv.Itoa(current.data))
		current = current.next
	}

	return strings.Join(xs, " -> ")

}

func (l *LinkedList) get(index int) *Node {
	if index == 0 {
		return l.head
	}
	if index > l.size {
		return nil
	}
	if index == l.size-1 {
		return l.tail
	}

	counter := 0
	current := l.head
	for current != nil {
		if counter == index {
			return current
		}
		counter++
		current = current.next
	}
	return nil
}

func main() {
	ll := makeList()
	ll.append(2)
	ll.append(12)
	ll.append(33)
	ll.append(45)
	ll.prepend(100)

	fmt.Println(ll.print())
}
