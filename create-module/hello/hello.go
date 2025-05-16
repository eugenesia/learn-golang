package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set propertiews of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		// log.Fatal(err) is equivalent to log.Print(err) followed by os.Exit(1).
		// It prints the error message and exits the program with a non-zero status.
		log.Fatal(err)
	}

	// If no error was returned, print the returned message.
	fmt.Println(message)
}
