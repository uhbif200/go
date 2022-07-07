package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var Options struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to"`
	Goodbye bool   `short:"g" long:"goodbye" description:"Say goodbye instead hello"`
}

// do not use default in bool

func main() {
	flags.Parse(&Options)
	if Options.Goodbye {
		fmt.Println("Goodbye, ", Options.Name)
	} else {
		fmt.Println("Hello, ", Options.Name)
	}
}
