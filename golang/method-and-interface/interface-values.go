package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	s string
}

func (t *T) M() {
	fmt.Println(t.s)
}

type MyFloat float64

func (f MyFloat) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Println("Print the I interface (vaule, type)")
	fmt.Printf("(%v, %T)\n", i, i)
}

func simple_interface_value_demo() {
	var i I
	i = &T{"Hello"}
	describe(i)

	i = MyFloat(math.Pi)
	describe(i)
}

func process_nil() {
	var i I
	var t *T
	i = t
	//i.M()
	fmt.Println(i)
}

func main() {
	fmt.Println("Interface values:")

	simple_interface_value_demo()

	process_nil()
}
