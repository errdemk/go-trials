package main

import (
    "fmt"
    "log"
    "rsc.io/quote"
    "github.com/errdemk/go-trials/greetings"
) 

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    fmt.Println(quote.Hello())

    message1, err1 := greetings.Hello("Captain Price")
    if err1 != nil {
        log.Fatal(err1)
    }
    fmt.Println(message1)

    message2, err2 := greetings.Hello("")
    if err2 != nil {
        log.Fatal(err2)
    }
    // If no error was returned, print the returned message to the console.
    fmt.Println(message2)
}

// https://go.dev/doc/