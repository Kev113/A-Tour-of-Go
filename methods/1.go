// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct {
// 	X, Y float64
// }

// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(v.Abs())
// }

package main  

import (  
	"fmt"  
)  

// List represents a singly-linked list that holds values of any type.  
type List[T any] struct {  
	next *List[T]  
	val  T  
}  

// AddToEnd appends a new element with the given value at the end of the list.  
func (l *List[T]) AddToEnd(val T) {  
	current := l  
	for current.next != nil {  
		current = current.next  
	}  
	current.next = &List[T]{val: val}  
}  

// AddToStart inserts a new element with the given value at the start of the list.  
func (l *List[T]) AddToStart(val T) *List[T] {  
	newNode := &List[T]{val: val, next: l}  
	return newNode  
}  

// Remove removes the first occurrence of the specified value from the list.  
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

// Find searches for the first occurrence of a value and returns true if found.  
func (l *List[T]) Find(val T) bool {  
	current := l  
	for current != nil {  
		if current.val == val {  
			return true  
		}  
		current = current.next  
	}  
	return false  
}  

// Print displays all elements in the list.  
func (l *List[T]) Print() {  
	current := l  
	for current != nil {  
		fmt.Printf("%v -> ", current.val)  
		current = current.next  
	}  
	fmt.Println("nil")  
}  

func main() {  
	var list *List[int]  

	// Create an initial list with a single element  
	list = &List[int]{val: 10}  
	list.Print()  

	// Add elements to the end  
	list.AddToEnd(20)  
	list.AddToEnd(30)  
	list.Print()  

	// Add element to the start  
	list = list.AddToStart(5)  
	list.Print()  

	// Remove an element  
	list = list.Remove(20)  
	list.Print()  

	// Find an element  
	found := list.Find(30)  
	fmt.Printf("Found 30: %v\n", found)  

	found = list.Find(100)  
	fmt.Printf("Found 100: %v\n", found)  
}