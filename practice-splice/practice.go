package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8
	for i := 0; i < dy; i++ {
		var temp []uint8
		for j := 0; j < dx; j++ {
			temp = append(temp, uint8(i%(j+1)))
		}
		result = append(result, temp)
	}
	return result
}

func main() {
	pic.Show(Pic)
}
