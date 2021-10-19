package main

import (
	"fmt"
	"runtime"
)

func main() {
	//fmt.Println(sqrt(2))
	fmt.Println(runtime.GOOS)
}

func sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}
