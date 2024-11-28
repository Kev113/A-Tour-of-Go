package main

import (
	"fmt"
)

type List[T comparable] struct {
	next *List[T]
	val T
}

func (l *List[T]) AddToEnd(val T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: val}
}

func (l *List[T]) AddToStart(val T) *List[T] {
	newNode := &List[T]{val: val, next: l}
	return newNode
}

func (l *List[T]) Remove(val T) *List[T] {
	if l == nil {
		return nil
	}
	if l.val == val {
		return l.next
	}
	current := l
	for current.next != nil && current.next.val != val {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
	return l
}

func (l *List[T]) Find(val T) bool {
	current := l
	for current.next != nil {
		if current.val == val {
			return true
		}
		current = current.next
	}
	return false
}

func (l *List[T]) Print() {
	current := l
	for current.next != nil {
		fmt.Printf("%v -> ", current.val)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	var list *List[int]

	list = &List[int]{val: 10}
	list.Print()

	list.AddToEnd(20)
	list.AddToEnd(30)
	list.AddToEnd(40)
	list.Print()

	list = list.Remove(20)
	list.Print()

	found := list.Find(30)
	fmt.Printf("Found: %v\n", found)

	found = list.Find(100)
	fmt.Printf("Found: %v\n", found)
}