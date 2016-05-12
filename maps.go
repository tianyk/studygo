// map 的文法跟结构体文法相似，不过必须有键名。
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex // map[keyType]valType

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
