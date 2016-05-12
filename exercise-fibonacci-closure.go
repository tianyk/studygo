package main

import "fmt"

func fibonacci() func() int {
	var prev, current, next int
	current = 1
	return func() int {
		next = current + prev
		prev = current
		current = next
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f())
	}
}
