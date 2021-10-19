package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	splits := strings.Fields(s)
	m := make(map[string]int)
	for _, s := range splits {
		_, ok := m[s]
		if ok {
			m[s]++
		} else {
			m[s] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
