package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := 0; i < 10; i++ {
		if <-c1 != <-c2 {
			return false
		}
	}

	return true
}

func Walk_print() {
	t := tree.New(1)
	ch := make(chan int)
	go Walk(t, ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	fmt.Println("Exercise Binary-Tree:")
	fmt.Println("Same(tree.New(1), tree.New(1)):", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2)):", Same(tree.New(1), tree.New(2)))
}
