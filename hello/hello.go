package main

import (
    "fmt"
    "rsc.io/quote"
    "github.com/errdemk/go-trials/greetings"
) 

func main() {
    fmt.Println(quote.Hello())

    message := greetings.Hello("Captain Price")
    fmt.Println(message)
}

// https://go.dev/doc/