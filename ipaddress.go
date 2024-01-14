package main

import "fmt"

type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr{
		"loopback": {1, 2, 3, 4},
		"gdns":     {3, 4, 5, 6},
	}

	for st, ip := range hosts {
		fmt.Println(st, ip)
	}

	defer fmt.Println("world")

	fmt.Println("hello")
}

func (i IPAddr) String() string {
	s := ""
	for a := range i {
		s += string(a)
		fmt.Println(a)
	}
	return s
}
