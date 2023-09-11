package main

import (
    "fmt"
    "hrutvik.gg/go_greet"
    "log"
)

import "rsc.io/quote/v4"


func main() {
    log.SetPrefix("[6GREETER9]: ")
    log.SetFlags(0)

    msg, err := go_greet.Greet("")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(msg, "\n", quote.Go())
}
