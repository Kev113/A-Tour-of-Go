package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for i := range pic {
		pic[i] = make([]uint8, dx)
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			pic[y][x] = uint8(x * y)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
