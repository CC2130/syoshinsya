package main

import "fmt"

type Introduce interface {
    introduce()
}

type MyInt int

func (i MyInt) introduce() {
    fmt.Printf("Hi, I'm a MyInt '%v'\n", i)
}

type MyFloat float64

func (i MyFloat) introduce() {
    fmt.Printf("Hi, I'm MyFloat '%v'\n", i)
}

type MyString string

func (i MyString) introduce() {
    fmt.Printf("Hi, I'm MyString '%v'\n", i)
}

func assert_interface_type() {
    fmt.Println("type assertions:")
    var i Introduce
    var value interface{}
    fmt.Println("value is empty interface.")
    i = MyInt(3.0)
    i.introduce()

    // assert 1
    value = i.(MyInt)
    fmt.Println("value = i.(MyInt) // can assert if i is MyInt type, and get the value")
    fmt.Printf("value is: %v\n", value)
    fmt.Println()

    // assert 2
    // i = MyFloat(22)
    // fmt.Println("re assigned i as MyFloat(22)")
    value, res := i.(MyFloat)
    fmt.Println("value, res := i.(MyFloat) // if MyFloat, res == true, or value init as MyFloat ZERO")
    if res {
        fmt.Printf("value(MyFloat) is: %v\n", value)
    } else {
        fmt.Printf("i is not MyFloat type, set value as %v\n", value)
    }
    fmt.Println()
}

func type_switches(i interface{}) {
    switch v := i.(type) {
    case MyInt:
        fmt.Printf("I'm %T %v.\n", v, v)
    case float64:
        fmt.Printf("I'm %T %v.\n", v, v)
    case string:
        fmt.Printf("I'm %T %v.\n", v, v)
    default:
        fmt.Printf("Unknown type: %T, value is: %v\n", v, v)
    }
}

func handle_interface_type_switche() {
    fmt.Println("Tyep-Switche:")
    fmt.Println("I can recongnize servral types...")
    type_switches(MyInt(3))
    type_switches(2)
    type_switches("Hello")
    type_switches(40.0)
    type_switches(true)
}

func main() {
    assert_interface_type()
    handle_interface_type_switche()

}
