package greetings

import (
	"fmt"
	"errors"
)

// ***** In Go, a function whose name starts with a capital letter can be called by a function not in the same package!!!
// ***** Any Go function can return multiple values.
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// this declaration is a shortcut of:
	// var message string
    // message = fmt.Sprintf("Hi, %v. Welcome!", name)
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    
	return message, nil
}