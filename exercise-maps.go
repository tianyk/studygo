package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	fields := strings.Fields(s)
	for _, field := range fields {
		count, ok := words[field]
		if ok {
			words[field] = count + 1
		} else {
			words[field] = 1
		}
	}
	return words
}

func main() {
	fmt.Println(WordCount("Hello  世界 你好 世界"))
	// wc.Test(WordCount)
}
