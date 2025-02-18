package main

import (
    "fmt"
    "sort"
)

var number = 32.12
var pointer2 *float64 = &number

func main() {
    pointer := &number
    fmt.Printf("%v\n", *pointer)
    println(*pointer2)

    names := [3]string{"Alice", "Charlie", "Bob"}
    fmt.Println(&names)
    sort.Strings(names[:])
    fmt.Println(&names)
}
