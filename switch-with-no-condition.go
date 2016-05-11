// 没有条件的 switch 同 `switch true` 一样。

// 这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch { // switch true
	case t.Hour() < 12: // t.Hour() < 12 is true?
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoo.")
	default:
		fmt.Println("Good evening.")
	}
}
