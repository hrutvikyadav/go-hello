package main

import (
    "fmt"
    "hrutvik.gg/go_greet"
    "log"
)

//import "rsc.io/quote/v4"


func main() {
    log.SetPrefix("[6GREETER9]: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := go_greet.Greets(names)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(messages)
}
