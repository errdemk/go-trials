package greetings

import "fmt"

// In Go, a function whose name starts with a capital letter can be called by a function not in the same package!!!
// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
	
	// this declaration is a shortcut of:
	// var message string
    // message = fmt.Sprintf("Hi, %v. Welcome!", name)
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    
	return message
}