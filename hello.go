package main

import (
    "fmt"
    "hrutvik.gg/go_greet"
)

import "rsc.io/quote/v4"


func main() {
    msg := go_greet.Greet("Hrutvik")
    fmt.Println(msg, "\n", quote.Go())
}
