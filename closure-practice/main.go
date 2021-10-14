package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func(int) int {
	a := 0
	b := 1
	return func(x int) int {
		if x >= 2 {
			temp := a
			a = b
			b = temp + b
			return b
		}
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
