package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
	"icovn.dev/network"
)

func main() {
	// Get a greeting message and print it.
	message := network.Hello("Gladys")
	fmt.Println(reverse.String(message), reverse.Int(24601))
}
