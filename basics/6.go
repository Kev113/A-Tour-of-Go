package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	x, y:= b, a
	fmt.Println(a, b)
	fmt.Println(x, y)
}