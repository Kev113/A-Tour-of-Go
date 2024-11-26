package main

import "fmt"

func fibonacci() func(int) int {
	previous := 1
	previousv_2 := 0
	var result int
	return func(x int) int {
		result = previous + previousv_2
		previousv_2 = previous
		previous = result
		return previous
	}
}

func main() {
	f := fibonacci()
	for i := 0; i< 10; i++ {
		fmt.Println(f(i))
	}
}