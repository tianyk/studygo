package main

// import "code.google.com/p/go-tour/pic"
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy) // 初始化长度为dy
	for i := 0; i < dy; i++ {
		image[i] = make([]uint8, dx) // 初始化长度为dx
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			image[i][j] = (uint8)(i+j) % 255
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
