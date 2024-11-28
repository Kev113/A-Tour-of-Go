package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, _ := range s {
		if x == s[i] {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

}