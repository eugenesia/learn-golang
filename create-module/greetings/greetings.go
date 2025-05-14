package greetings

import "fmt"

// Hello returns a greeting for the named person.
// In Go, a function whose name starts with a capital letter is exported and
// can be called from other packages.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	// := declares and initialises a variable in one step.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
