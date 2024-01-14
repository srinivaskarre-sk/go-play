package main

import "fmt"

func main() {
	a := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	for i, v := range a {
		fmt.Println(i, v)
	}

	b := map[string]int{
		"abc": 1,
		"bcd": 2,
	}

	for key, value := range b {
		fmt.Println(key, value)
	}
}
