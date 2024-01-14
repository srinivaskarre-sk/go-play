package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	res := strings.Split(s, " ")

	m := make(map[string]int)

	for _, v := range res {
		ele, ok := m[v]

		if ok {
			m[v] = ele + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
	fmt.Println(wordCount("hi hello how hi abc hello"))
}
