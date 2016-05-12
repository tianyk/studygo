package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5) // len=5 cap=5 [0 0 0 0 0]
	printSlice("b", b)
	c := b[:2] // len=2 cap=5 [0 0]
	printSlice("c", c)
	d := c[2:5] // len=3 cap=3 [0 0 0]
	printSlice("d", d)
	e := d[:2] // len=2 cap=3 [0 0] 仅仅提供结束索引时cap不会变
	printSlice("e", e)
}
