package main

import "fmt"

func main() {
	a := make([]int, 0, 3)

	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(a), len(a), a)
	}
}
