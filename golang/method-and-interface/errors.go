package main

import (
    "fmt"
    "time"
)

// Build a MyError struct
type MyError struct {
    When time.Time
    What string
}

// error is an interface have Error() function, return string
func (e *MyError) Error() string {
    return fmt.Sprintf("At %v, %v\n", e.When, e.What)
}

func run() error {
    return &MyError{
        time.Now(),
        "It crashed!",
    }
}

func main() {
    fmt.Println("Generaly a function return an error value,\n",
                "if not `nil`, then need to handle the error\n")
    if err := run(); err != nil {
        fmt.Println("Catch the error:")
        fmt.Println(err)
    }
}
