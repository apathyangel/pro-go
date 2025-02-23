package main

import "fmt"

func main() {
    arr := []string{"123"}

    for index, value := range arr {
        switch value {
        case "123":
            fmt.Println("index:", index, "value:", value)
            fallthrough
        case "421":
            fmt.Println("index:", index, "value:", value)
        }
    }

    counter := 0
target:
    fmt.Println("Counter", counter)
    counter++
    if counter < 5 {
        goto target
    }
}
