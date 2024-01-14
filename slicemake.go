package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice1(a)

	a = make([]int, 0, 6)
	printSlice1(a)

}

func printSlice1(s []int) {
	fmt.Printf("len: %d cap: %d, %v \n", len(s), cap(s), s)
}
