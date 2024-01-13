package main

import "fmt"

func main() {
	var i interface{} = "string"

	doSomething(i)

	var j interface{} = 100
	doSomething(j)

	var k interface{} = 100.9099
	doSomething(k)
}

func doSomething(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println("it is int")

	case float64:
		fmt.Println("float")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("don't know %T", t)
	}
}
