package main

import (
	"fmt"
	"math"
)

// Simple demo for method
// the reciver is a value
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func method_simple() {
	fmt.Println("Simeple method demo:")
	f := MyFloat(-math.Sqrt2)
	fmt.Println("f is: ", f)
	fmt.Println("f.Abs() is: ", f.Abs())
}

// End of simple demo for method

// Method reciver type as pointer
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	fmt.Printf("I recive a value, the addr is: %p\n", &v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// if the reciver do not use *(pointer), will not change v
// func (v Vertex) Scale(f float64) {}
// the 'v' in Scale is a clone of v
func (v *Vertex) Scale(f float64) {
	fmt.Printf("I recive a pointer, the addr is: %p\n", v)
	v.X = v.X * f
	v.Y = v.Y * f
}

func method_with_pointer() {
	v := Vertex{3, 4}
	fmt.Printf("The original v is: %v, addr is: %p\n", v, &v)
	v.Scale(10)
	v.Abs()
	fmt.Println("Only method with pointer as reciver can modify the value!")
}

// End of method reciver type as pointer

func main() {
	fmt.Println("Hello Method")
	//method_simple()
	method_with_pointer()
}
