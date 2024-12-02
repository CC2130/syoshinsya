package main

import (
	"fmt"
	//"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		//time.Sleep(300 * time.Millisecond)
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 3, 7, 8, -1, 5}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
	fmt.Println("Try to uncomment the time.Sleep above, and run sereral times")
}
