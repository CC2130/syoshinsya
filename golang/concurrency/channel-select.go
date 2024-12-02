package main

import (
	"fmt"
	"time"
)

// review
// build and print a fibonacci list
func fibonacci_demo() {
	x, y := 0, 1
	l := make([]int, 10)

	for i := 0; i < 10; i++ {
		l[i] = x
		x, y = y, x+y
	}

	fmt.Println(l)
}

// exsrcise for select, case/default
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit!")
			return
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Printf(".")
		}
	}
}

func main() {
	//fibonacci_demo()
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("Every 200 ms take a number from channel...")
		for i := 0; i < 10; i++ {
			fmt.Printf("\n%v", <-c)
			time.Sleep(time.Millisecond)
		}
		fmt.Println("\nExit in 1 second")
		time.Sleep(1000 * time.Millisecond)
		quit <- 0
	}()
	fibonacci(c, quit)
}
