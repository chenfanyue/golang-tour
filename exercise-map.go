package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	slice := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range slice {
		m[v] += 1
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}
