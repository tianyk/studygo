// 函数也是值。
package main

import (
	"fmt"
	"math"
)

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Printf("Function type: %T\n", hypot)
	fmt.Println(hypot(3, 4))
}
