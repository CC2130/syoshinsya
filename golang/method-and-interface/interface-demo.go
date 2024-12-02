package main

import (
	"fmt"
	"math"
)

// Simeple dmeo for interface
// A value of interface type can hold any value that implement those methods
// in this demo any value have Abs() method can be use as Abser type
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func interface_simple() {
	var a Abser
	fmt.Println("Init var a Abser, Abser is an interface implement Abs() method.")
	f := MyFloat(2)
	fmt.Println("  f is MyFloat type value.")
	v := Vertex{3, 4}
	fmt.Println("  v is Vertex type value.")
	var vp *Vertex
	vp = &v
	fmt.Println("  vp is Vertex pointer type")

	// MyFloat have Abs() method
	fmt.Println("MyFloat have Abs() method, a = f is good")
	a = f

	// Vertex do not have Abs() method
	// a = v
	fmt.Println("Vertex do not have Abs() method, a = v cause error!")

	// *Vertex have Abs() method
	fmt.Println("*Vertex have Abs() method, a = vp be fine.")
	a = vp

	a.Abs()
	fmt.Println("")
}

// End of simeple dmeo for interface

// Interface are implemented implicitly
func interface_implicitly() {
	fmt.Println("Interface are implemented implicitly!")
	fmt.Println("var a Abser = MyFloat(10)")
	var a Abser = MyFloat(10)
	fmt.Println("a.Abs() = ", a.Abs())
}

// End of interface are implemented implicitly

func main() {
	fmt.Println("Study Interface!\n")
	interface_simple()
	interface_implicitly()
}
