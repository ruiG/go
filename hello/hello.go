package main

import (
	"fmt"

	"github.com/ruiG/go/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Rui")
	fmt.Println(message)
}
