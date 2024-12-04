package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(T *tree.Tree, ch chan int) {
	WalkImpl(T, ch)
	close(ch)
}

func WalkImpl(T *tree.Tree, ch chan int) {
	if T == nil {
		return
	}
	WalkImpl(T.Left, ch)
	ch <- T.Value
	WalkImpl(T.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := make(chan int), make(chan int)

	go Walk(t1, w1)
	go Walk(t2, w2)

	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2

		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Println("tree.New(1) == tree.New(1): ")
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}

	fmt.Println("tree.New(1) == tree.New(2): ")
	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}
