package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("i, value is: %v, type is: %T\n", i, i)
}

func main() {
	var i interface{}
	fmt.Println("Init var i as a empty interface")
	fmt.Println("So i can be assigned to any type....")
	fmt.Println()
	i = 2
	describe(i)

	i = "two"
	describe(i)
}
