package main

import (
    "fmt"
    "io"
    "os"
    "strings"
    "golang.org/x/tour/reader"
)

type MyReader struct {}

// MyReader.read(b []byte) always return a 'A' stream, infinite...
// Thinking: cannot use cap(b) as parameter
func (r MyReader) Read(b []byte) (int, error) {
    //fmt.Printf("len: %v, cap: %v\n", len(b), cap(b))
    length := len(b)
    for i := 0; i < length; i++ {
        b[i] = 'A'
    }

    return length, nil
}
func reader_exercise() {
    // simeple reader demo exercise
    fmt.Println("MyReader.read(b []byte) always return a 'A' stream, infinite...")
    reader.Validate(MyReader{})
}

// rot13Reader exercise
type rot13Reader struct {
    r io.Reader
}

func rot13(b byte) byte {
    if ((65 <= b && b <= 77) || (97 <= b && b <= 109)) {
        return b + 13
    } else if ((78 <= b && b <= 96) || (110 <= b && b <= 122)) {
        return b - 13
    }
    return b
}

func (r rot13Reader) Read(b []byte) (int, error) {
    n, err := r.r.Read(b)
    for i := 0; i < n; i++ {
        b[i] = rot13(b[i])
    }
    return n, err
}

func rot13Reader_test() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

func main() {
    fmt.Println()
    reader_exercise()

    rot13Reader_test()
}
