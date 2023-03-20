package main

import "fmt"

func main() {
    msg := hello("Vadym")
    fmt.Println(msg)
}

func hello(name string) string {
    return fmt.Sprintf("Hi %s", name)
}