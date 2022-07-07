package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to")
var hello = flag.Bool("hello", true, "Saying hello")
var goodbye = flag.Bool("goodbye", false, "Saying goodbye")

func init() {
	flag.BoolVar(&hello, "h", true, "Do we say hello")
}

func main() {
	flag.Parse()
	if *hello && *goodbye {
		fmt.Println("Cannot print both hello and goodbye")
		return
	}
	if *hello {
		fmt.Println("Hello, ", *name)
	}
	if *goodbye {
		fmt.Println("Hello, ", *name)
	}
}
