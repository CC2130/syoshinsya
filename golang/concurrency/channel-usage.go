package main

import (
	"fmt"
	"time"
)

func reciver(c chan int, id int) {
	for i := range c {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("I am reciver %v, recived %v\n", id, i)
	}
}

func main() {
	fmt.Println("Create a buffered channel c, put 1..10 in it.\n")
	c := make(chan int, 100)
	for i := 1; i < 11; i++ {
		c <- i
	}

	v, ok := <-c
	fmt.Printf("v, ok = <-c\n")
	fmt.Printf("v = %v, ok = %v\n\n", v, ok)

	fmt.Println("Use 'for v := range c' to recived from channel:")
	go reciver(c, 1)
	go reciver(c, 2)

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("\nclose the channel\n")
	close(c)

	v, ok = <-c
	fmt.Printf("v, ok = <-c\n")
	fmt.Printf("v = %v, ok = %v\n", v, ok)
}
