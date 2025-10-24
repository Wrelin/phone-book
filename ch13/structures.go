package main

import (
	"fmt"
)

type node[T comparable] struct {
	Data T
	next *node[T]
}

type list[T comparable] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{
		Data: data,
		next: nil,
	}

	if l.start == nil {
		l.start = &n
		return
	}

	if l.start.next == nil {
		l.start.next = &n
		return
	}

	temp := l.start
	l.start = l.start.next
	l.add(data)
	l.start = temp
}

func (l *list[T]) search(data T) *node[T] {
	curNode := l.start
	for curNode != nil {
		if data == curNode.Data {
			return curNode
		}

		curNode = curNode.next
	}

	return nil
}

func (l *list[T]) delete(data T) {
	rmNode := l.search(data)
	if rmNode == nil {
		return
	}

	prevNode := l.start
	if prevNode == rmNode {
		prevNode = nil
	}

	for prevNode != nil {
		if prevNode.next == rmNode {
			break
		}
		prevNode = prevNode.next
	}

	if prevNode == nil {
		l.start = rmNode.next
		return
	}

	prevNode.next = rmNode.next
}

func (l *list[T]) print() {
	curNode := l.start
	for curNode != nil {
		fmt.Println("*", curNode)
		curNode = curNode.next
	}
}

func main() {
	var myList list[int]
	fmt.Println(myList)
	myList.add(12)
	myList.add(9)
	myList.add(3)
	myList.add(9)

	// Print all elements
	myList.print()

	fmt.Println(myList.search(9))
	fmt.Println(myList.search(20))

	myList.delete(3)
	myList.delete(12)
	myList.print()
}
