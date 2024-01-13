package main

import "fmt"

type IPAddr [4]byte

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func (i IPAddr) String() string {
	s := ""
	for a := range i {
		s += string(a)
		fmt.Println(a)
	}
	return s
}
