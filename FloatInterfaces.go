package main

import (
	"fmt"
	"math"
)

type AbsoluteInterface interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (m MyFloat) Abs() float64 {
	return float64(m)
}

func (v Vertex) Abs() float64 {
	return v.X * v.Y
}

func main() {
	var a AbsoluteInterface
	var myf = MyFloat(-math.Sqrt2)
	var vert = Vertex{3.0, 5.0}

	a = myf
	fmt.Println(a.Abs())

	a = &vert
	fmt.Println(a.Abs())
}
