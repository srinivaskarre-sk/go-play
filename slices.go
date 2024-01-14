package main

import "fmt"

func main() {
	names := []string{
		"Kevin",
		"Pravin",
		"mevin",
		"vivin",
		"juvin",
	}

	first := names[0:2]
	second := names[2:4]

	fmt.Println(first, second, "original", names)

	first[0] = "updated"

	second[0] = "updatedFromSecond"

	fmt.Println(first, second, "original", names)

	//defining slices

	slice1 := []int{1, 2, 3, 4, 5, 6}

	slice2 := []bool{true, false, false, true}

	slice3 := []struct {
		x int
		y bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}
	fmt.Println(slice1, slice2, slice3)

	//slicing and expanding
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s = s[:4]
	printSlice(s)

	s = s[1:2]
	printSlice(s)

	var s1 []int
	fmt.Println(s1, len(s1), cap(s1))
	if s1 == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int) {
	fmt.Printf("len: %d, cap: %d, %v \n", len(s), cap(s), s)
}
