package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
	"icovn.dev/network"
)

func test() {
	// Get a greeting message and print it.
	message := network.Hello("Gladys")
	fmt.Println(reverse.String(message))
}
