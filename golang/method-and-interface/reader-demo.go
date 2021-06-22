package main

import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello Reader!")

    b := make([]byte, 8)
    for {
        n, err := r.Read(b)
        // won't print useless line such as "n = 0, err...."
        if n != 0 {
            fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
        } else if err == io.EOF {
            break
        }
    }
}
